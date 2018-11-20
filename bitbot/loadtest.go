package bitbot

import (
	"fmt"
	"github.com/whyrusleeping/hellabot"
	"strings"
)

var loadTrigger = hbot.Trigger{
	func(irc *hbot.Bot, m *hbot.Message) bool {
		return m.Command == "PRIVMSG" && strings.TrimSpace(m.Content) == "!load"
	},
	func(irc *hbot.Bot, m *hbot.Message) bool {
		irc.Reply(m, "Loading testTrigger")
		b.Bot.AddTrigger(testTrigger)
		return true
	},
	"loadTrigger",
}

var unloadTrigger = hbot.Trigger{
	func(irc *hbot.Bot, m *hbot.Message) bool {
		return m.Command == "PRIVMSG" && strings.TrimSpace(m.Content) == "!unload"
	},
	func(irc *hbot.Bot, m *hbot.Message) bool {
		irc.Reply(m, "Unloading testTrigger")
		b.Bot.DropTrigger(testTrigger)
		return true
	},
	"unloadTrigger",
}

var testTrigger = hbot.Trigger{
	func(irc *hbot.Bot, m *hbot.Message) bool {
		return m.Command == "PRIVMSG" && strings.TrimSpace(m.Content) == "!test"
	},
	func(irc *hbot.Bot, m *hbot.Message) bool {
		irc.Reply(m, fmt.Sprintf("Hello, %s", m.From))
		return false
	},
	"testTrigger",
}
