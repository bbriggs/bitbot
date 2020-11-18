package bitbot

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/whyrusleeping/hellabot"
)

var (
	location   *time.Location
	timeFormat string
)

// ReminderEvent : The Gorm struct that represents an event in the DB.
type NickList struct {
	ID   int    `gorm:"unique;AUTO_INCREMENT;PRIMARY_KEY"`
	Nick string `gorm:"unique"`
}

var NickShuffleTrigger = NamedTrigger{ //nolint:gochecknoglobals,golint
	ID:   "nickShuffle",
	Help: "Add a nick to my shuffle. Usage: !nick add|drop <nick>",
	Init: func() error {
		return b.DB.AutoMigrate(&NickList{}).Error
	},
	Condition: func(irc *hbot.Bot, m *hbot.Message) bool {
		return m.Command == "PRIVMSG" && strings.HasPrefix(m.Trailing, "!remind")
	},
	Action: func(irc *hbot.Bot, m *hbot.Message) bool {
		timeFormat = "2006-01-02 15:04"

		splitMSG := strings.Split(m.Content, " ")
		if len(splitMSG) < 2 {
			irc.Reply(m, "Not enough arguments provided")
			return true
		}

		switch splitMSG[1] {
		case "time":
			irc.Reply(m, getTime())
		case "add":
			irc.Reply(m, addEvent(m, irc))
		case "remove":
			irc.Reply(m, removeEvent(m))
		case "list":
			irc.Reply(m, listEvents(m, irc))
		case "join":
			irc.Reply(m, joinEvent(m))
		case "part":
			irc.Reply(m, partEvent(m))
		default:
			irc.Reply(m, "Wrong argument")
		}
		return true
	},
}
