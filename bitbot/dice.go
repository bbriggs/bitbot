package bitbot

import (
	"github.com/whyrusleeping/hellabot"
	"github.com/justinian/dice"
	"strings"
	"fmt"
)
var RollTrigger = hbot.Trigger{
	func(irc *hbot.Bot, m *hbot.Message) bool {
		return m.Command == "PRIVMSG" && strings.TrimSpace(m.Content) == "!roll"
	},
	func(irc *hbot.Bot, m *hbot.Message) bool {
		
		return true
	},
}