package bitbot

import (
	"fmt"
	"github.com/whyrusleeping/hellabot"
	"strings"
)

var loadTrigger = hbot.Trigger{
	func(irc *hbot.Bot, m *hbot.Message) bool {
		return m.Command == "PRIVMSG" && strings.TrimSpace(m.Content) == "!load" && b.isAdmin(m)
	},
	func(irc *hbot.Bot, m *hbot.Message) bool {
		split := strings.Split(m.Content, " ")
		if len(split) < 2 {
			irc.Reply(m, "Load what?")
		} else {
			b.Bot.AddTrigger(split[1])
		}
		return true
	},
}

var unloadTrigger = hbot.Trigger{
	func(irc *hbot.Bot, m *hbot.Message) bool {
		return m.Command == "PRIVMSG" && strings.TrimSpace(m.Content) == "!unload" && b.isAdmin(m)
	},
	func(irc *hbot.Bot, m *hbot.Message) bool {
		split := strings.Split(m.Content, " ")
		if len(split) < 2 {
			irc.Reply(m, "Unload what?")
		} else {
			if b.Bot.DropTrigger(split[1]) {
				irc.Reply(m, fmt.Sprintf("%s unloaded", split[1]))
			} else {
				irc.Reply(m, fmt.Sprintf("I don't think %s is loaded...", split[1]))
			}
		}
		return true
	},
}

var testTrigger = NamedTrigger{
	ID: "testTrigger",
	Condition: func(irc *hbot.Bot, m *hbot.Message) bool {
		return m.Command == "PRIVMSG" && strings.TrimSpace(m.Content) == "!test"
	},
	Action: func(irc *hbot.Bot, m *hbot.Message) bool {
		fullname := fmt.Sprintf("%s!%s@%s", m.Name, m.User, m.Host)
		irc.Reply(m, fullname)
		return false
	},
}
