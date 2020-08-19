package bitbot

import (
	"github.com/whyrusleeping/hellabot"
	"strings"
)

var LennyTrigger = NamedTrigger{
	ID:   "lennyface",
	Help: "Usage: !lenny",
	Condition: func(irc *hbot.Bot, m *hbot.Message) bool {
		return m.Command == "PRIVMSG" && strings.TrimSpace(m.Content) == "!lenny"
	},
	Action: func(irc *hbot.Bot, m *hbot.Message) bool {
		irc.Reply(m, `( ͡° ͜ʖ ͡°)`)
		return true
	},
}
