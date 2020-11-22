package bitbot

import (
	"strings"

	"github.com/whyrusleeping/hellabot"
)

var NickTakenTrigger = NamedTrigger{ //nolint:gochecknoglobals,golint
	ID:   "nick",
	Help: "Avoids nick collisions by renaming the bot if the nick is already taken. Not recommended to use with nickRandomizer.",
	Condition: func(irc *hbot.Bot, m *hbot.Message) bool {
		/* get the host's name by cutting the port number, and making sure that the message comes from host */
		var comesFromHost = (m.From == strings.Split(irc.Host, ":")[0])

		var nickTaken = strings.Contains(m.Content, "Nickname is already in use")

		return comesFromHost && nickTaken
	},
	Action: func(irc *hbot.Bot, m *hbot.Message) bool {
		irc.SetNick(irc.Nick + "_")
		return false
	},
}

var NickRecoverTrigger = NamedTrigger{ //nolint:gochecknoglobals,golint
	ID:   "nick-recover",
	Help: "Watch for QUIT messages, and recover nick at first occasion",
	Condition: func(irc *hbot.Bot, m *hbot.Message) bool {
		return m.Command == "QUIT" && m.From == b.Config.Nick
	},
	Action: func(irc *hbot.Bot, m *hbot.Message) bool {
		tryNickRecovery(irc)
		return false
	},
}

var ManualNickRecoverTrigger = NamedTrigger{ //nolint:gochecknoglobals,golint
	ID:   "nick-manual",
	Help: "Attempt nick recovery at !nick-reco",
	Condition: func(irc *hbot.Bot, m *hbot.Message) bool {
		return m.Command == "PRIVMSG" &&
			strings.ToLower(strings.TrimSpace(m.Content)) == "!nick-reco"
	},
	Action: func(irc *hbot.Bot, m *hbot.Message) bool {
		tryNickRecovery(irc)
		return false
	},
}

var NickRandomizerTrigger = NamedTrigger{ //nolint:gochecknoglobals,golint
	ID:   "nickRandomizer",
	Help: "Avoids nick collisions by renaming the bot to a random nick from the database if the nick is already taken. Not recommended to use with NickTakenTrigger.",
	Condition: func(irc *hbot.Bot, m *hbot.Message) bool {
		/* get the host's name by cutting the port number, and making sure that the message comes from host */
		var comesFromHost = (m.From == strings.Split(irc.Host, ":")[0])

		var nickTaken = strings.Contains(m.Content, "Nickname is already in use")

		return comesFromHost && nickTaken
	},
	Action: func(irc *hbot.Bot, m *hbot.Message) bool {
		nick, err := getRandomNick()
		if err != nil {
			b.Config.Logger.Error("nickRandomizer: Unable to get a new nickname")
			b.Config.Logger.Error("nickRandomizer: Falling back to underscore behavior")
			irc.SetNick(irc.Nick + "_")
		}
		irc.SetNick(nick)
		return false
	},
}

func tryNickRecovery(irc *hbot.Bot) {
	b.Config.Logger.Info("Attempting Nick recovery")
	irc.SetNick(b.Config.Nick)
}
