package bitbot

import (
	"strings"

	"github.com/whyrusleeping/hellabot"
)

var NickTrigger = NamedTrigger{
	ID: "nick",
	Condition: func(irc *hbot.Bot, m *hbot.Message) bool {
		var comesFromHost = (m.From == strings.Split(irc.Host, ":")[0])

		var nickTaken = strings.Contains(m.Content,
			"Nickname is already in use")

		return comesFromHost && nickTaken
	},
	Action: func(irc *hbot.Bot, m *hbot.Message) bool {
		irc.SetNick(irc.Nick + "_")
		return false
	},
}
