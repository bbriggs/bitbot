package bitbot

import (
	"github.com/whyrusleeping/hellabot"
)

var IgnoreTrigger = NamedTrigger{ //nolint:gochecknoglobals,golint
	ID:   "ignore",
	Help: "Ignore messages from other bots.",
	Condition: func(irc *hbot.Bot, m *hbot.Message) bool {
		return stringInSlice(m.From, b.Config.Ignored)
	},
	Action: func(irc *hbot.Bot, m *hbot.Message) bool {
		b.Config.Logger.Info("Ignored message", "from", m.From)
		return true // consume the message, so that other triggers don't get to process it
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
