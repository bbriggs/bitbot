package bitbot

import (
	"github.com/whyrusleeping/hellabot"
)

var TableFlipTrigger = NamedTrigger{
	ID:   "tableflip",
	Help: "Flip a table. Usage: !tableflip",
	Condition: func(irc *hbot.Bot, m *hbot.Message) bool {
		match := (m.Content == "!tableflip")
		return m.Command == "PRIVMSG" && match
	},
	Action: func(irc *hbot.Bot, m *hbot.Message) bool {
		irc.Reply(m, "(╯°□°）╯︵ ┻━┻")
		return true
	},
}

var TableUnflipTrigger = NamedTrigger{
	ID: "unflip",
	Condition: func(irc *hbot.Bot, m *hbot.Message) bool {
		match := (m.Content == "!unflip")
		return m.Command == "PRIVMSG" && match
	},
	Action: func(irc *hbot.Bot, m *hbot.Message) bool {
		irc.Reply(m, "┬─┬ ノ( ゜-゜ノ)")
		return true
	},
}
