package bitbot

import (
	"os"
	"fmt"
	"github.com/whyrusleeping/hellabot"
	log "gopkg.in/inconshreveable/log15.v2"
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

	irc.Logger.SetHandler(log.StreamHandler(os.Stdout, log.JsonFormat()))
	irc.Run()
	fmt.Println("Bot shutting down.")
}
