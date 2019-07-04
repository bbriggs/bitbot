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

	// private map of triggers
	triggers sync.Map
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
	Plugins      []hbot.Handler // Plugins to start with
}

var b Bot = Bot{}

func (b *Bot) RegisterTrigger(t NamedTrigger) {
	b.triggers.Store(t.Name(), t)
}

func (b *Bot) FetchTrigger(name string) (NamedTrigger, bool) {
	res, ok := b.triggers.Load(name)
	if !ok {
		return NamedTrigger{}, false
	}
	return res.(NamedTrigger), true
}

func (b *Bot) DropTrigger(t NamedTrigger) bool {
	b.triggers.Delete(t.Name())
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
	b.Bot.AddTrigger(NickTrigger)
	for _, trigger := range config.Plugins {
		b.Bot.AddTrigger(trigger)
	}

	b.Bot.Logger.SetHandler(log.StreamHandler(os.Stdout, log.JsonFormat()))

	// GOOOOOOO
	defer b.DB.Close()
	b.Bot.Run()
}
