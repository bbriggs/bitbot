package bitbot

import (
	"fmt"

	"github.com/whyrusleeping/hellabot"
)

var InviteTrigger = NamedTrigger{
	ID:   "invite",
	Help: "Follow invites to other channels. Usage: /invite [bot nick]",
	Condition: func(irc *hbot.Bot, m *hbot.Message) bool {
		return m.Command == "INVITE"
	},
	Action: func(irc *hbot.Bot, m *hbot.Message) bool {
		fmt.Println(m)
		return true
	},
}
