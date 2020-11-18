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
	ID   int    `gorm:"unique;AUTO_INCREMENT;PRIMARY_KEY"` // Primary Key
	Nick string `gorm:"unique"`                            // Nickname to use
	From string // Submitter of the nickname
}

var NickShuffleTrigger = NamedTrigger{ //nolint:gochecknoglobals,golint
	ID:   "nickShuffle",
	Help: "Add a nick to my shuffle. Usage: !nick add <nick>",
	Init: func() error {
		return b.DB.AutoMigrate(&NickList{}).Error
	},
	Condition: func(irc *hbot.Bot, m *hbot.Message) bool {
		return m.Command == "PRIVMSG" && strings.HasPrefix(m.Trailing, "!nick add <nickname>")
	},
	Action: func(irc *hbot.Bot, m *hbot.Message) bool {
		ok, err := addNickToDB(m)
		if !ok {
			irc.Reply("Looks like something went wrong.")
			irc.Reply(err.Error())
		}
		return true
	},
}

func addNickToDB(m *hbot.Message) (bool, error) {
	// split message, error out if too short
	splitMsg := strings.Split(m.Content, " ")
	if len(splitMsg) < 3 {
		return false, fmt.Errorf("Not enough arguments. See !help nickShuffle")
	}

	// grab nick, error out if invalid
	newNick := NickList{
		Nick: splitMsg[2],
		From: m.From,
	}

	// insert into database
	b.DB.NewRecord(newNick)
	if err := b.DB.Create(&newNick); err != nil {
		return false, err
	}

	return true, nil
}
