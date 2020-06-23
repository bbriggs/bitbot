package bitbot

import (
	"math/rand"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/mb-14/gomarkov"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/whyrusleeping/hellabot"
	log "gopkg.in/inconshreveable/log15.v2"
)

type Bot struct {
	Bot    *hbot.Bot
	DB     *gorm.DB
	Random *rand.Rand // Initialized PRNG
	Config Config

	triggers     map[string]NamedTrigger // For "registered" triggers
	triggerMutex *sync.RWMutex
	counters     map[string]*prometheus.CounterVec
	gauges       map[string]*prometheus.GaugeVec
	mChain       *gomarkov.Chain // Initialized Markov chain. Accessed and updated by markov triggers.
	markovMutex  *sync.RWMutex
}

type Config struct {
	NickservPass string         // Nickserv password
	OperUser     string         // Username for server oper
	OperPass     string         // Password for server oper
	Channels     []string       // slice of channels to connect to (must include #)
	Nick         string         // nick to use
	Server       string         // server:port for connections
	SSL          bool           // Enable SSL for the connection
	Admins       ACL            // slice of masks representing administrators
	Ignored      []string       // slice of strings containing ignored nicks (bots)
	Plugins      []NamedTrigger // Plugins to start with
	Prometheus   bool           // Enable Prometheus
	PromAddr     string         // Listen address for prometheus endpoint
	DBConfig     DBConfig       // Configuration settings for Database connection
	Logger       log.Logger     // The logger used by all of the bot
}

// Configuration struct for Postgresql backend
type DBConfig struct {
	User    string
	Pass    string
	Host    string
	Port    string
	Name    string
	SSLMode string
}

var b Bot = Bot{}

func (b *Bot) RegisterTrigger(t NamedTrigger) {
	var err error

	b.triggerMutex.Lock()
	b.triggers[t.Name()] = t
	b.triggerMutex.Unlock()
	b.Bot.AddTrigger(t)

	if t.Init != nil {
		err = t.Init()
	}

	if err != nil {
		b.Config.Logger.Error("Trigger " + t.Name() + " failed to initialize: " + err.Error())
	}
}

func (b *Bot) FetchTrigger(name string) (NamedTrigger, bool) {
	b.triggerMutex.RLock()
	defer b.triggerMutex.RUnlock()
	t, ok := b.triggers[name]
	return t, ok
}

func (b *Bot) DropTrigger(t NamedTrigger) bool {
	b.triggerMutex.Lock()
	delete(b.triggers, t.Name())
	b.triggerMutex.Unlock()
	return true
}

func Run(config Config) {
	config.Logger.Info("Initializing bitbot...")
	config.Logger.Info("Setting up IRC connection...")
	// Initialize connection
	chans := func(bot *hbot.Bot) {
		bot.Channels = b.Config.Channels
	}
	sslOptions := func(bot *hbot.Bot) {
		bot.SSL = b.Config.SSL
	}
	b.Config = config
	irc, err := hbot.NewBot(b.Config.Server, b.Config.Nick, chans, sslOptions)
	if err != nil {
		config.Logger.Error(err.Error())
		os.Exit(1)
	}
	b.Bot = irc

	//TODO get hellabot logs in a subfield to indicate clearly that they are not direct
	// bitbot logs
	b.Bot.Logger.SetHandler(log.StreamHandler(os.Stdout, log.JsonFormat()))

	config.Logger.Info("Connecting to postgres...")

	db, err := newDB(config.DBConfig)
	if err != nil {
		config.Logger.Error("Database connection unsuccessful: " + err.Error())
	} else {
		config.Logger.Info("Database connection successful!")
	}
	b.DB = db

	b.Random = rand.New(rand.NewSource(time.Now().UnixNano()))
	b.triggerMutex = &sync.RWMutex{}
	b.markovMutex = &sync.RWMutex{}
	b.triggers = make(map[string]NamedTrigger)
	b.counters = make(map[string]*prometheus.CounterVec)
	b.gauges = make(map[string]*prometheus.GaugeVec)

	config.Logger.Info("Loading triggers...")
	// These are non-optional and added to every bot instance
	b.Bot.AddTrigger(IgnoreTrigger)
	b.Bot.AddTrigger(OperLogin)
	b.Bot.AddTrigger(loadTrigger)
	b.Bot.AddTrigger(unloadTrigger)
	b.Bot.AddTrigger(NickTakenTrigger)
	for _, trigger := range config.Plugins {
		config.Logger.Info(trigger.Name() + " loaded")
		b.RegisterTrigger(trigger)
	}

	config.Logger.Info("Starting prometheus on http:// " + b.Config.PromAddr + "/metrics")
	// Prometheus stuff
	if b.Config.Prometheus {
		b.createCounters()
		b.Bot.AddTrigger(MessageCounterTrigger)
		b.Bot.AddTrigger(ChannelPopGaugeTrigger)
		b.Bot.AddTrigger(SetChanPopGaugeTrigger)
		b.Bot.AddTrigger(HandleListReplyTrigger)
		http.Handle("/metrics", promhttp.Handler())
		go http.ListenAndServe(b.Config.PromAddr, nil)
	}

	// GOOOOOOO
	defer b.DB.Close()
	config.Logger.Info("Starting bitbot...")
	b.Bot.Run()

}
