package bitbot

import (
	"github.com/whyrusleeping/hellabot"
	"strings"
)

var ShrugTrigger = NamedTrigger{ //nolint:gochecknoglobals,golint
	ID:   "shrug",
	Help: "Usage: !shrug",
	Condition: func(irc *hbot.Bot, m *hbot.Message) bool {
		return m.Command == "PRIVMSG" && strings.TrimSpace(m.Content) == "!shrug"
	},
	Action: func(irc *hbot.Bot, m *hbot.Message) bool {
		irc.Reply(m, `¯\_(ツ)_/¯`)
		return true
	},
}

var LennyTrigger = NamedTrigger{ //nolint:gochecknoglobals,golint
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

var WeebTrigger = NamedTrigger{ //nolint:gochecknoglobals,golint
	ID:   "DamnWeebs",
	Help: "Usage: !weeb",
	Condition: func(irc *hbot.Bot, m *hbot.Message) bool {
		return m.Command == "PRIVMSG" && strings.TrimSpace(strings.ToLower(m.Content)) == "!weeb"
	},
	Action: func(irc *hbot.Bot, m *hbot.Message) bool {
		irc.Reply(m, `gahd damn weebs`) //add list to randomly choose from
		return true
	},
}
