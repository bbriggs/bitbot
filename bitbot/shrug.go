package bitbot

import (
	"github.com/whyrusleeping/hellabot"
	"fmt"
)

var ShrugTrigger = hbot.Trigger{
	func(irc *hbot.Bot, m *hbot.Message) bool {
		return m.Command == "PRIVMSG" && m.Content == "!shrug"
	},
	func(irc *hbot.Bot, m *hbot.Message) bool {
		resp := fmt.Sprintf("¯\_(ツ)_/¯") irc.Reply(m, resp) return true
	},
}
