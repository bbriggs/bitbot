package bitbot

import (
	"github.com/whyrusleeping/hellabot"
	"math/rand"
	"strings"
)

var EpeenTrigger = NamedTrigger{ //nolint:gochecknoglobals,golint
	ID:   "epeen",
	Help: "epeen returns the length of the requesters epeen. Usage: !epeen",
	Condition: func(irc *hbot.Bot, m *hbot.Message) bool {
		return m.Command == "PRIVMSG" && strings.TrimSpace(m.Content) == "!epeen"
	},
	Action: func(irc *hbot.Bot, m *hbot.Message) bool {
		var epeen = makeEpeenAnswer(m.From)
		irc.Reply(m, epeen)
		return true
	},
}

func makeEpeenAnswer(nick string) string {
	var peepee = ""

	switch nick {
	case "daemon":
		peepee = strings.Repeat("\\", rand.Intn(200)) + "D" + "\n(_)_)"
	case "m242":
		peepee = "8D" // Yup, childish :D
	case "skidd0":
		peepee = "8=ancap=D"
	default:
		peepee = "8" + strings.Repeat("=", rand.Intn(20)) + "D"
	}
	return nick + "'s peepee: " + peepee

}
