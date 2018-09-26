package main

import (
	"github.com/whyrusleeping/hellabot"
	"fmt"
	"os"
)

func main() {
	chans := func(bot *hbot.Bot) {
		bot.Channels = []string{"#social"}
	}
	sslOptions := func(bot *hbot.Bot) {
		bot.SSL = true
	}

	irc, err := hbot.NewBot("irc.sithmail.com:6697", "bitbot2", chans, sslOptions)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	irc.Run()
}
