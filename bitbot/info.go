package bitbot

import (
	"github.com/whyrusleeping/hellabot"
	"strings"
	"fmt"
)

var InfoTrigger = hbot.Trigger{
	func(irc *hbot.Bot, m *hbot.Message) bool {
		return m.Command == "PRIVMSG" && strings.TrimSpace(m.Content) == "!info"
	},
	func(irc *hbot.Bot, m *hbot.Message) bool {
		resp := fmt.Sprintf("Bitbot version %s (%s/%s) | %s", GitVersion, GitBranch, GitCommit, SourceRepo)
		irc.Reply(m, resp)
		return true
	},
}
