package bitbot

import (
	"fmt"
	"strings"

	"github.com/whyrusleeping/hellabot"
)

// NickList : The Gorm struct that represents a nickname row in the database
type NickList struct {
	ID   int    `gorm:"unique;AUTO_INCREMENT;PRIMARY_KEY"` // Primary Key
	Nick string `gorm:"unique"`                            // Nickname to use
	From string // Submitter of the nickname
}

var NickShuffleTrigger = NamedTrigger{ //nolint:gochecknoglobals,golint
	ID:   "nickShuffle",
	Help: "Add a nick to my shuffle. Usage: !nick add|drop|shuffle [nick]",
	Init: func() error {
		return b.DB.AutoMigrate(&NickList{}).Error
	},
	Condition: func(irc *hbot.Bot, m *hbot.Message) bool {
		return m.Command == "PRIVMSG" && strings.HasPrefix(m.Trailing, "!nick")
	},
	Action: func(irc *hbot.Bot, m *hbot.Message) bool {
		msg, err := nickShuffleDispatcher(irc, m)
		if err != nil {
			irc.Reply(m, err.Error())
			return true
		}
		irc.Reply(m, msg)
		return true
	},
}

func nickShuffleDispatcher(irc *hbot.Bot, m *hbot.Message) (string, error) {
	// split message, error out if too short
	splitMsg := strings.Split(m.Content, " ")
	if len(splitMsg) < 2 {
		return "", fmt.Errorf("Not enough arguments. See !help nickShuffle")
	}

	switch splitMsg[1] {
	case "add":
		return addNickToDB(m)
	case "shuffle":
		return shuffleNickFromDB(irc)
	case "drop":
		return dropNickFromDB(m)
	default:
		return "", fmt.Errorf("Invalid argument. See !help nickShuffle")
	}

	return "", fmt.Errorf("Switch statement failed somehow")
}

func addNickToDB(m *hbot.Message) (string, error) {
	// split message, error out if too short
	splitMsg := strings.Split(m.Content, " ")

	if len(splitMsg) < 3 {
		return "", fmt.Errorf("Not enough arguments. See !help nickShuffle")
	}

	// grab nick, error out if invalid
	newNick := NickList{
		Nick: splitMsg[2],
		From: m.From,
	}

	// insert into database
	b.DB.NewRecord(newNick)
	if res := b.DB.Create(&newNick); res.Error != nil {
		return "", res.Error
	}

	return fmt.Sprintf("%s added %s to my rotation", newNick.From, newNick.Nick), nil
}

func shuffleNickFromDB(irc *hbot.Bot) (string, error) {

	nick, err := getRandomNick()
	if err != nil {
		return "", nil
	}

	irc.SetNick(nick)
	return "honk", nil
}

func getRandomNick() (string, error) {
	var nicks []NickList

	res := b.DB.Find(&nicks)

	if res.Error != nil {
		return "", res.Error
	}

	// embrace the nest
	return nicks[b.Random.Intn(len(nicks))].Nick, nil
}

func dropNickFromDB(m *hbot.Message) (string, error) {
	return "can't delete yet, boss", nil
}
