package bitbot

import (
	"github.com/whyrusleeping/hellabot"
	"strings"
	"time"
)

var location *time.Location

var ReminderTrigger = NamedTrigger { //nolint:gochecknoglobals,golint
	ID:   "reminder",
	Help: "Set up events and remind them to concerned people. Usage: !remind list|time|add|remove|delete|join|part",
	Condition: func(irc *hbot.Bot, m *hbot.Message) bool {
		return m.Command == "PRIVMSG" && strings.HasPrefix(m.Trailing, "!remind")
	},
	Action: func(irc *hbot.Bot, m *hbot.Message) bool {
		location, _ = time.LoadLocation("UTC")

		splitMSG := strings.Split(m.Content, " ")
		if len(splitMSG) < 2 {
			irc.Reply(m, "Not enough arguments provided")
			return true
		}
		switch splitMSG[1] {
		case "time":
			irc.Reply(m, getTime())
		default:
			irc.Reply(m, "Wrong argument")
		}
		return true
	},
}

func getTime() string {
	now := time.Now().In(location)
	return now.Format("2006-01-02 15:04")
}
