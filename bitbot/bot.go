package bitbot

import (
	"os"

	bolt "go.etcd.io/bbolt"
	"github.com/whyrusleeping/hellabot"
	log "gopkg.in/inconshreveable/log15.v2"
)

type Bot struct {
	Bot *hbot.Bot
	DB  *bolt.DB
}

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

	bot := Bot{
		Bot: irc,
		DB:  db,
	}

	// Triggers to run
	bot.Bot.AddTrigger(InfoTrigger)
	bot.Bot.Logger.SetHandler(log.StreamHandler(os.Stdout, log.JsonFormat()))

	// GOOOOOOO
	defer bot.DB.Close()
	irc.Run()
}

