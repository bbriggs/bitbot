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
}

var b Bot = Bot{}

func Run(server string, nick string, channels []string, ssl bool) {
	chans := func(bot *hbot.Bot) {
		bot.Channels = channels
	}
	sslOptions := func(bot *hbot.Bot) {
		bot.SSL = ssl
	}

	irc, err := hbot.NewBot(server, nick, chans, sslOptions)
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}

	db, err := newDB()
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}

	b.Bot = irc
	b.DB = db
	b.Random = rand.New(rand.NewSource(time.Now().UnixNano()))

	// Triggers to run
	// Passive triggers. Unskippable.
	b.Bot.AddTrigger(TrackIdleUsers)

	// Begin with skip prefix (!skip)
	b.Bot.AddTrigger(SkipTrigger)
	b.Bot.AddTrigger(InfoTrigger)
	b.Bot.AddTrigger(ShrugTrigger)
	//b.Bot.AddTrigger(ReportIdleUsers)
	b.Bot.AddTrigger(URLReaderTrigger)
	b.Bot.AddTrigger(AbyssTrigger)
	b.Bot.Logger.SetHandler(log.StreamHandler(os.Stdout, log.JsonFormat()))

	// GOOOOOOO
	defer b.DB.Close()
	b.Bot.Run()
}
