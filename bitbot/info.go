package bitbot

import (
	"fmt"
	"github.com/whyrusleeping/hellabot"
	"strings"
)

var InfoTrigger = NamedTrigger{
	ID:   "info",
	Help: "Get version and repo information about this bot. Usage: !info",
	Condition: func(irc *hbot.Bot, m *hbot.Message) bool {
		return m.Command == "PRIVMSG" && strings.TrimSpace(m.Content) == "!info"
	},
	Action: func(irc *hbot.Bot, m *hbot.Message) bool {
		resp := fmt.Sprintf("Bitbot version %s (%s/%s) | %s", GitVersion, GitBranch, GitCommit, SourceRepo)
		irc.Reply(m, resp)
		return true
	},
}
