package bitbot

import (
	"github.com/whyrusleeping/hellabot"
	"strings"
)

var HelpTrigger = NamedTrigger{
	ID:   "help",
	Help: "Usage: !help [trigger name]",
	Condition: func(irc *hbot.Bot, m *hbot.Message) bool {
		return m.Command == "PRIVMSG" && strings.HasPrefix(m.Content, "!help")

	},
	Action: func(irc *hbot.Bot, m *hbot.Message) bool {
		splitMsg := strings.Split(m.Trailing, " ")
		triggers := b.ListTriggers()
		if len(splitMsg) < 2 {
			irc.Reply(m, "Currently loaded plugins: "+strings.Join(triggers, ", "))
			return true
		}

		// Argument provided, search for trigger and return help text
		for _, t := range triggers {
			if splitMsg[1] == t {
				// fetch and return help text
				t, ok := b.FetchTrigger(splitMsg[1])
				if ok && t.Help != "" { // Most triggers probably won't have a help field right away
					irc.Reply(m, t.Help)
				} else {
					irc.Reply(m, "Trigger found but help unavailalbe")
				}
				return true
			}
		}

		// Fallthrough
		irc.Reply(m, "Help text unavailable")
		return true

	},
}
