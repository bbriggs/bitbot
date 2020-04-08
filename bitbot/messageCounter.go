package bitbot

import (
	"github.com/whyrusleeping/hellabot"
)

var MessageCounterTrigger = NamedTrigger{ //nolint:gochecknoglobals
	ID:   "messageCounter",
	Help: "Increments a counter for every message it sees in chat.",
	Condition: func(irc *hbot.Bot, m *hbot.Message) bool {
		return m.Command == "PRIVMSG"

	},
	Action: func(irc *hbot.Bot, m *hbot.Message) bool {
		b.counters["messageCounter"].WithLabelValues(m.To, m.Name).Inc()
		return false
	},
}
