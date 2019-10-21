package bitbot

import (
	"github.com/whyrusleeping/hellabot"
	"regexp"
)

var BeefyTrigger = NamedTrigger{
	ID:   "beefy",
	Help: "It's big. It's beefy. It triggers any time someone says \"beefy\" in a message.",
	Condition: func(irc *hbot.Bot, m *hbot.Message) bool {
		match, _ := regexp.MatchString(`(?i)beefy`, m.Trailing)
		return m.Command == "PRIVMSG" && match
	},
	Action: func(irc *hbot.Bot, m *hbot.Message) bool {
		responses := []string{
			"BEEFY",
			"it's what's for dinner",
			"https://i.imgur.com/VbC5GLl.jpg",
			"mmmmmm",
		}
		irc.Reply(m, responses[b.Random.Intn(len(responses))])
		return true
	},
}
