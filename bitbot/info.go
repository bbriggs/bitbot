package bitbot

import (
	"github.com/whyrusleeping/hellabot"
)

var InfoTrigger = hbot.Trigger{
	func(bot *hbot.Bot, m *hbot.Message) bool {
		return m.Command == "PRIVMSG" && m.Content == "!info"
	},
	func(irc *hbot.Bot, m *hbot.Message) bool {
		irc.Reply(m, "https://github.com/bbriggs/bitbot")
		return false
	},
}
