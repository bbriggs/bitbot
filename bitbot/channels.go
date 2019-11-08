package bitbot

import (
	"fmt"
	"strings"

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

var PartTrigger = NamedTrigger{
	ID:   "part",
	Help: "Command the bot to leave the channel. Usage: [bot nick] part [channel]",
	Condition: func(irc *hbot.Bot, m *hbot.Message) bool {
		return m.Command == "PRIVMSG" && strings.HasPrefix(m.Content, irc.Nick)
	},
	Action: func(irc *hbot.Bot, m *hbot.Message) bool {
		splitMsg := strings.Split(m.Content, " ")
		fmt.Println(splitMsg)
		if len(splitMsg) == 2 {
			// part same channel message was received on
			partChannel(irc, m.Params[0], irc.Nick)
			return true
		} else if len(splitMsg) > 2 {
			// part channel provided in args
			partChannel(irc, splitMsg[2], irc.Nick)
			fmt.Printf("\n\n%+v\n\n", m)
			return true
		}
		// fallthrough; not enough args
		return false
	},
}
