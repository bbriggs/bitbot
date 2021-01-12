package bitbot

import (
	"github.com/whyrusleeping/hellabot"
	"regexp"
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
	Help: "Usage: mention uwu",
	Condition: func(irc *hbot.Bot, m *hbot.Message) bool {
		// thanks d4 for the regex help!
		match, _ := regexp.MatchString(`(?i)(.)w(\1)`, m.Content)
		return m.Command == "PRIVMSG" && match
	},
	Action: func(irc *hbot.Bot, m *hbot.Message) bool {
		adjectives := []string{ //add more as necessary
			"damn",
			"fucking",
			"degenerate",
			"incurable",
			"disgusting",
			"wonderful",
			"utter",
		}
		adj := adjectives[b.Random.Intn(len(adjectives))]
		reply := m.Name + `, you ` + adj + ` weeb!`
		irc.Reply(m, reply)
		return false
	},
}
