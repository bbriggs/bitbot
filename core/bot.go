package core

import (
	"github.com/whyrusleeping/hellabot"
	"github.com/bbriggs/bitbot/triggers"
	log "gopkg.in/inconshreveable/log15.v2"
	"os"
)

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

	// Triggers to run
	//irc.AddTrigger(triggers.EchoTrigger)
	irc.AddTrigger(triggers.InfoTrigger)
	irc.Logger.SetHandler(log.StreamHandler(os.Stdout, log.JsonFormat()))
	irc.Run()
}
