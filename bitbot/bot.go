package bitbot

import (
	"os"

	bolt "go.etcd.io/bbolt"
	"github.com/whyrusleeping/hellabot"
	log "gopkg.in/inconshreveable/log15.v2"
)

const Version string = "0.1.1"
var GitCommit string

type Bot struct {
	Bot *hbot.Bot
	DB  *bolt.DB
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
	// Triggers to run
	b.Bot.AddTrigger(InfoTrigger)
	b.Bot.AddTrigger(TrackIdleUsers)
	b.Bot.AddTrigger(ReportIdleUsers)
	b.Bot.AddTrigger(URLReaderTrigger)
	b.Bot.Logger.SetHandler(log.StreamHandler(os.Stdout, log.JsonFormat()))

	// GOOOOOOO
	defer b.DB.Close()
	b.Bot.Run()
}

