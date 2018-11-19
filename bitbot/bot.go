package bitbot

import (
	"math/rand"
	"os"
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
}

type Config struct {
	NickservPass string   // Nickserv password
	OperUser     string   // Username for server oper
	OperPass     string   // Password for server oper
	Channels     []string // slice of channels to connect to (must include #)
	Nick         string   // nick to use
	Server       string   // server:port for connections
	SSL          bool     // Enable SSL for the connection
	Admins       []string // slice of masks representing administrators
}

var b Bot = Bot{}

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

	// Triggers to run
	// Passive triggers. Unskippable.
	b.Bot.AddTrigger(TrackIdleUsers)
	b.Bot.AddTrigger(OperLogin)
	// Begin with skip prefix (!skip)
	b.Bot.AddTrigger(SkipTrigger)
	b.Bot.AddTrigger(InfoTrigger)
	b.Bot.AddTrigger(ShrugTrigger)
	//b.Bot.AddTrigger(ReportIdleUsers)
	b.Bot.AddTrigger(URLReaderTrigger)
	b.Bot.AddTrigger(AbyssTrigger)

	// GOOOOOOO
	defer b.DB.Close()
	b.Bot.Run()
}
