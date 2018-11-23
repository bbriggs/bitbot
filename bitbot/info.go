package bitbot

import (
	"fmt"
	"github.com/whyrusleeping/hellabot"
	"strings"
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
