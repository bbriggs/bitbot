package bitbot

import (
	"github.com/whyrusleeping/hellabot"
	"math/rand"
	"strings"
)

var EpeenTrigger = NamedTrigger{
	ID:   "epeen",
	Help: "epeen returns the length of the requesters epeen. Usage: !epeen",
	Condition: func(irc *hbot.Bot, m *hbot.Message) bool {
		return m.Command == "PRIVMSG" && strings.TrimSpace(m.Content) == "!epeen"
	},
	Action: func(irc *hbot.Bot, m *hbot.Message) bool {
		var nick = m.From
		var peepee = ""
		if nick == "suser" || nick == "not_suser" {
			peepee = "8=D"
		} else {
			peepee = "8" + strings.Repeat("=", rand.Intn(20)) + "D"
		}
		var reply = nick + "'s peepee: " + peepee

		irc.Reply(m, reply)
		return true
	},
}
