package bitbot

import (
	"fmt"
	"github.com/whyrusleeping/hellabot"
	"strings"
)

var loadTrigger = hbot.Trigger{
	func(irc *hbot.Bot, m *hbot.Message) bool {
		return m.Command == "PRIVMSG" && strings.HasPrefix(m.Content, "!load") && b.isAdmin(m)
	},
	func(irc *hbot.Bot, m *hbot.Message) bool {
		split := strings.Split(m.Content, " ")
		if len(split) < 2 {
			irc.Reply(m, "Load what?")
		} else {
			t, ok := b.FetchTrigger(split[1])
			if !ok {
				irc.Reply(m, "Looks like this trigger was never registered. Can't load it.")
			} else {
				b.Bot.AddTrigger(t)
				irc.Reply(m, "Trigger loaded.")
			}
		}
		return true
	},
}

var unloadTrigger = hbot.Trigger{
	func(irc *hbot.Bot, m *hbot.Message) bool {
		return m.Command == "PRIVMSG" && strings.HasPrefix(m.Content, "!unload") && b.isAdmin(m)
	},
	func(irc *hbot.Bot, m *hbot.Message) bool {
		split := strings.Split(m.Content, " ")
		if len(split) < 2 {
			irc.Reply(m, "Unload what?")
		} else {
			t, ok := b.FetchTrigger(split[1])
			if ok {
				dropped := b.DropTrigger(t)
				if dropped {
					irc.Reply(m, fmt.Sprintf("Trigger %s dropped", t.Name()))
				} else {
					irc.Reply(m, fmt.Sprintf("Unable to drop %s for some reason...", t.Name()))
				}
			} else {
				irc.Reply(m, fmt.Sprintf("I don't think you registered the %s trigger...", split[1]))
			}
		}

		return true
	},
}

var listTriggers = NamedTrigger{ //nolint:gochecknoglobals,golint
	ID: "listTriggers",
	Condition: func(irc *hbot.Bot, m *hbot.Message) bool {
		return m.Command == "PRIVMSG" && strings.TrimSpace(m.Content) == "!triggers" && b.isAdmin(m)
	},
	Action: func(irc *hbot.Bot, m *hbot.Message) bool {
		triggers := b.ListTriggers()
		irc.Reply(m, strings.Join(triggers, ", "))
		return true
	},
}
