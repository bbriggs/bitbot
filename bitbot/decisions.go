package bitbot

import (
	"fmt"
	"strings"

	"github.com/whyrusleeping/hellabot"
)

var DecisionsTrigger = NamedTrigger{
	ID: "decisions",
	Condition: func(irc *hbot.Bot, m *hbot.Message) bool {
		prefix := fmt.Sprintf("%s choose", irc.Nick)
		return m.Command == "PRIVMSG" && strings.HasPrefix(m.Content, prefix)
	},
	Action: func(irc *hbot.Bot, m *hbot.Message) bool {
		prefix := fmt.Sprintf("%s choose", irc.Nick)
		msg := strings.TrimPrefix(m.Content, prefix)
		splitMsg := strings.Split(msg, " or ")
		if len(splitMsg) < 2 {
			irc.Reply(m, "What am I supposed to be deciding here?")
			return false
		}
		r := strings.TrimSpace(splitMsg[b.Random.Intn(len(splitMsg))])
		irc.Reply(m, r)
		return false
	},
}
