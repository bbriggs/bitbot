package bitbot

import (
	"github.com/whyrusleeping/hellabot"
	"math/rand"
    "hash/crc64"
	"strings"
    "time"
)

var EpeenTrigger = NamedTrigger{ //nolint:gochecknoglobals,golint
	ID:   "epeen",
	Help: "epeen returns the length of the requesters epeen. Usage: !epeen",
	Condition: func(irc *hbot.Bot, m *hbot.Message) bool {
		return m.Command == "PRIVMSG" && strings.ToLower(strings.TrimSpace(m.Content)) == "!epeen"
	},
	Action: func(irc *hbot.Bot, m *hbot.Message) bool {
		var epeen = makeEpeenAnswer(m.From)
		irc.Reply(m, epeen)
		return true
	},
}

func makeEpeenAnswer(nick string) string {
	peepeeSize := 20
	peepeeCrc := crc64.Checksum([]byte(nick+time.Now().Format("2006-01-02")),crc64.MakeTable(crc64.ECMA))
	peepeeRnd := rand.New(rand.NewSource(int64(peepeeCrc)))
	peepee := "8" + strings.Repeat("=", peepeeRnd.Intn(peepeeSize)) + "D"
	return nick + "'s peepee: " + peepee
}
