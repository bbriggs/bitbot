package bitbot

import (
	"math/rand"
	"os"
	"sync"
	"time"

	"github.com/whyrusleeping/hellabot"
	bolt "go.etcd.io/bbolt"
	log "gopkg.in/inconshreveable/log15.v2"
)

type Bot struct {
	Bot    *hbot.Bot
	DB     *bolt.DB
	Random *rand.Rand // Initialized PRNG
	Config Config

	triggers     map[string]NamedTrigger // For "registered" triggers
	triggerMutex *sync.RWMutex
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
	db, err := newDB()
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}

	b.DB = db
	b.Random = rand.New(rand.NewSource(time.Now().UnixNano()))
	b.Config = config
	b.triggerMutex = &sync.RWMutex{}
	b.triggers = make(map[string]NamedTrigger)

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
	//b.Bot.Log("Loading triggers")
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

	// GOOOOOOO
	defer b.DB.Close()
	b.Bot.Run()
}
