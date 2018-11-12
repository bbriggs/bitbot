package bitbot

import (
	"github.com/whyrusleeping/hellabot"
)

var AbyssTrigger = hbot.Trigger{
	func(irc *hbot.Bot, m *hbot.Message) bool {
		return m.Command == "PRIVMSG" && b.Random.Intn(100) < 5
	},
	func(irc *hbot.Bot, m *hbot.Message) bool {
		irc.Reply(m, "0.0")
		return true
	},
}
