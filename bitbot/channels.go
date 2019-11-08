package bitbot

import (
	"github.com/whyrusleeping/hellabot"
)

var InviteTrigger = NamedTrigger{
	ID:   "invite",
	Help: "Follow invites to other channels. Usage: /invite [bot nick]",
	Condition: func(irc *hbot.Bot, m *hbot.Message) bool {
		return m.Command == "INVITE"
	},
	Action: func(irc *hbot.Bot, m *hbot.Message) bool {
		irc.Join(m.Content)
		return true
	},
}
