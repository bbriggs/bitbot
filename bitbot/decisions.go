package bitbot

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/whyrusleeping/hellabot"
)

var DecisionsTrigger = NamedTrigger{
	ID:   "decisions",
	Help: "Let the bot decide something for you. Usage: ${bot_name} choose option option [option...]",
	Condition: func(irc *hbot.Bot, m *hbot.Message) bool {
		prefix := fmt.Sprintf("%s choose", irc.Nick)
		return m.Command == "PRIVMSG" && strings.HasPrefix(m.Content, prefix)
	},
	Action: func(irc *hbot.Bot, m *hbot.Message) bool {
		prefix := fmt.Sprintf("%s choose", irc.Nick)
		msg := strings.TrimPrefix(m.Content, prefix)
		r := choose(msg)
		if r == "" {
			r = "Choose what?"
		}
		irc.Reply(m, r)
		return false
	},
}

func choose(m string) string {
	s := strings.Split(m, " or ")
	if len(s) < 2 {
		return ""
	}
	return strings.TrimSpace(s[rand.Intn(len(s))])
}
