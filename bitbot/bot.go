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
	Plugins      []NamedTrigger // Plugins to start with
	Prometheus   bool           // Enable Prometheus
	PromAddr     string         // Listen address for prometheus endpoint
	DBConfig     DBConfig       // Configuration settings for Database connection
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
	b.triggerMutex.Lock()
	b.triggers[t.Name()] = t
	b.triggerMutex.Unlock()
	b.Bot.AddTrigger(t)
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
	db, err := newDB(config.DBConfig)
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}

	b.DB = db
	b.Random = rand.New(rand.NewSource(time.Now().UnixNano()))
	b.Config = config
	b.triggerMutex = &sync.RWMutex{}
	b.markovMutex = &sync.RWMutex{}
	b.triggers = make(map[string]NamedTrigger)
	b.counters = make(map[string]*prometheus.CounterVec)
	b.gauges = make(map[string]*prometheus.GaugeVec)

	chans := func(bot *hbot.Bot) {
		bot.Channels = b.Config.Channels
	}
	sslOptions := func(bot *hbot.Bot) {
		bot.SSL = b.Config.SSL
	}

	irc, err := hbot.NewBot(b.Config.Server, b.Config.Nick, chans, sslOptions)
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}

	b.Bot = irc
	b.Bot.Logger.SetHandler(log.StreamHandler(os.Stdout, log.JsonFormat()))
	// These are non-optional and added to every bot instance
	b.Bot.AddTrigger(OperLogin)
	b.Bot.AddTrigger(loadTrigger)
	b.Bot.AddTrigger(unloadTrigger)
	b.Bot.AddTrigger(NickTakenTrigger)
	for _, trigger := range config.Plugins {
		log.Info(trigger.Name() + " loaded")
		b.RegisterTrigger(trigger)
	}

	b.Bot.Logger.SetHandler(log.StreamHandler(os.Stdout, log.JsonFormat()))

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
	b.Bot.Run()

}
