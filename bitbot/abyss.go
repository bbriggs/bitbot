package bitbot

import (
	"github.com/whyrusleeping/hellabot"
	"math/rand"
)

var AbyssTrigger = NamedTrigger{ //nolint:gochecknoglobals,golint
	ID:   "abyss",
	Help: "State of the art advanced Abyss simulator. Non-interactive.",
	Condition: func(irc *hbot.Bot, m *hbot.Message) bool {
		return m.Command == "PRIVMSG" && rand.Intn(1000) < 2
	},
	Action: func(irc *hbot.Bot, m *hbot.Message) bool {
		irc.Reply(m, "0.0")
		return true
	},
}
