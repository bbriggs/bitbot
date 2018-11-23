package bitbot

import (
	"fmt"
	"github.com/whyrusleeping/hellabot"
	"strings"
)

type NamedTrigger struct {
	ID        string
	Condition func(*hbot.Bot, *hbot.Message) bool
	Action    func(*hbot.Bot, *hbot.Message) bool
}

func (t NamedTrigger) Name() string {
	return t.ID
}

// Handle executes the trigger action if the condition is satisfied
func (t NamedTrigger) Handle(b *hbot.Bot, m *hbot.Message) bool {
	if !t.Condition(b, m) {
		return false
	}
	return t.Action(b, m)
}

var loadTrigger = hbot.Trigger{
	func(irc *hbot.Bot, m *hbot.Message) bool {
		return m.Command == "PRIVMSG" && strings.TrimSpace(m.Content) == "!load"
	},
	func(irc *hbot.Bot, m *hbot.Message) bool {
		irc.Reply(m, "Loading testTrigger")
		b.Bot.AddTrigger(testTrigger)
		return true
	},
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
}

var testTrigger = NamedTrigger{
	ID: "testTrigger",
	Condition: func(irc *hbot.Bot, m *hbot.Message) bool {
		return m.Command == "PRIVMSG" && strings.TrimSpace(m.Content) == "!test"
	},
	Action: func(irc *hbot.Bot, m *hbot.Message) bool {
		irc.Reply(m, fmt.Sprintf("Hello, %s", m.From))
		return false
	},
}
