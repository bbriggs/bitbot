package bitbot

import (
	"github.com/whyrusleeping/hellabot"
)

var AbyssTrigger = NamedTrigger{
	ID: "abyss",
	Condition: func(irc *hbot.Bot, m *hbot.Message) bool {
		return m.Command == "PRIVMSG" && b.Random.Intn(1000) < 2
	},
	Action: func(irc *hbot.Bot, m *hbot.Message) bool {
		irc.Reply(m, "0.0")
		return true
	},
}
