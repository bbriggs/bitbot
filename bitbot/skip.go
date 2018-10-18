package bitbot

import (
	"github.com/whyrusleeping/hellabot"
	"strings"
)

// SkipTrigger sets a message prefix that instructs bitbot not to process the message
// Should be set before any "skippable" triggers and after any triggers that run on all messages (unskippable)
var SkipTrigger = hbot.Trigger{
	func(irc *hbot.Bot, m *hbot.Message) bool {
		return m.Command == "PRIVMSG" && strings.HasPrefix(m.Content,"!skip")
	},
	func(irc *hbot.Bot, m *hbot.Message) bool {
		return true  // Do nothing and stop processing any more triggers
	},
}
