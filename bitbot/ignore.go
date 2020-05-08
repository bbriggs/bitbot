package bitbot

import (
	"github.com/whyrusleeping/hellabot"
)

var IgnoreTrigger = NamedTrigger{ //noliny:gochecknoglobals,golint
	ID:   "ignore",
	Help: "Ignore messages from other bots.",
	Condition: func(irc *hbot.Bot, m *hbot.Message) bool {
		return stringInSlice(m.From, b.Config.Ignored)
	},
	Action: func(irc *hbot.Bot, m *hbot.Message) bool {
		return true // consume the message
	},
}

func stringInSlice(element string, whole []string) bool {
	for _, e := range whole {
		if element == e {
			return true
		}
	}
	return false
}
