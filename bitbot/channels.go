package bitbot

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/whyrusleeping/hellabot"
)

var InviteTrigger = NamedTrigger{ //nolint:gochecknoglobals,golint
	ID:   "invite",
	Help: "Follow invites to other channels. Usage: /invite [bot nick]",
	Condition: func(irc *hbot.Bot, m *hbot.Message) bool {
		return m.Command == "INVITE"
	},
	Action: func(irc *hbot.Bot, m *hbot.Message) bool {
		// In the rfc1459 and 2812, irc invites have the channel in parameters
		// of the message, but putting it in the trailing part of the message
		// seems very popular, as at least UnrealIRCD and freenode's ircd-seven
		// do it this way.
		// However, oragono ircd follows rfc 1459 correctly, the only way we will implement.
		// Here, we have two parameters to an invite, the name of the bot in first, and
		// The name of the channel in second.
		channel := m.Params[1]
		b.Config.Logger.Info("Got an invite message", "channel", channel)
		irc.Join(channel)
		return true
	},
}

var PartTrigger = NamedTrigger{ //nolint:gochecknoglobals,golint
	ID:   "part",
	Help: "Command the bot to leave the channel. Usage: [bot nick] part [channel]",
	Condition: func(irc *hbot.Bot, m *hbot.Message) bool {
		isPartMessage, err := regexp.MatchString("^"+irc.Nick+".*part",
			m.Content)
		if err != nil {
			b.Config.Logger.Error(err.Error())
		}
		return m.Command == "PRIVMSG" && isPartMessage
	},
	Action: func(irc *hbot.Bot, m *hbot.Message) bool {
		splitMsg := strings.Split(m.Content, " ")
		fmt.Println(splitMsg)
		if len(splitMsg) == 2 {
			// part same channel message was received on
			irc.Part(m.Params[0], irc.Nick)
			return true
		} else if len(splitMsg) > 2 {
			// part channel provided in args
			irc.Part(splitMsg[2], irc.Nick)
			return true
		}
		// fallthrough; not enough args
		return false
	},
}
