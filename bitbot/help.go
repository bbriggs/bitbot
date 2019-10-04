package bitbot

import (
	"github.com/whyrusleeping/hellabot"
	"strings"
)

var HelpTrigger = NamedTrigger{
	ID: "help",
	Condition: func(irc *hbot.Bot, m *hbot.Message) bool {
		return m.Command == "PRIVMSG" && strings.TrimSpace(m.Content) == "!help"

	},
	Action: func(irc *hbot.Bot, m *hbot.Message) bool {
		triggers := b.ListTriggers()
		irc.Reply(m, "Currently loaded plugins: "+strings.Join(triggers, ", "))
		return true
	},
}
