package bitbot

import (
	"github.com/whyrusleeping/hellabot"
	"math/rand"
	"strconv"
	"strings"
)

var TarotTrigger = NamedTrigger{ //nolint:gochecknoglobals
	ID:   "Tarot",
	Help: "Request tarot cards, default 1. Usage !tarot [num cards].",
	Condition: func(irc *hbot.Bot, m *hbot.Message) bool {
		return m.Command == "PRIVMSG" && strings.HasPrefix(m.Trailing, "!tarot")
	},
	Action: func(irc *hbot.Bot, m *hbot.Message) bool {
		if len(m.Content) < 7 {
			resp := tarotCards[rand.Intn(len(tarotCards)-1)]
			irc.Reply(m, resp)
		} else {
			// Random permutation of indexes for drawing tarotCards
			deck := rand.Perm(len(tarotCards))
			msg := strings.TrimPrefix(m.Content, "!tarot ")
			if num, err := strconv.Atoi(msg); err == nil {
				if num < 1 || num > len(tarotCards) {
					num = 1
				}
				if num > 5 {
					for i := 0; i < num; i++ {
						irc.Msg(m.From, tarotCards[deck[i]])
					}
					irc.Reply(m, "I have PMed you your reading.")
				} else {
					for i := 0; i < num; i++ {
						irc.Reply(m, tarotCards[deck[i]])
					}
				}
			} else {
				irc.Reply(m, "Try again..")
			}
		}
		return true
	},
}

var tarotCards = []string{
	"I of Swords",
	"II of Swords",
	"III of Swords",
	"IIII of Swords",
	"V of Swords",
	"VI of Swords",
	"VII of Swords",
	"VIII of Swords",
	"VIIII of Swords",
	"X of Swords",
	"Page of Swords",
	"Queen of Swords",
	"King of Swords",
	"Knight of Swords",
	"I of Cups",
	"II of Cups",
	"III of Cups",
	"IIII of Cups",
	"V of Cups",
	"VI of Cups",
	"VII of Cups",
	"VIII of Cups",
	"VIIII of Cups",
	"X of Cups",
	"Page of Cups",
	"Queen of Cups",
	"King of Cups",
	"Knight of Cups",
	"I of Wands",
	"II of Wands",
	"III of Wands",
	"IIII of Wands",
	"V of Wands",
	"VI of Wands",
	"VII of Wands",
	"VIII of Wands",
	"VIIII of Wands",
	"X of Wands",
	"Page of Wands",
	"Queen of Wands",
	"King of Wands",
	"Knight of Wands",
	"I of Pentacles",
	"II of Pentacles",
	"III of Pentacles",
	"IIII of Pentacles",
	"V of Pentacles",
	"VI of Pentacles",
	"VII of Pentacles",
	"VIII of Pentacles",
	"VIIII of Pentacles",
	"X of Pentacles",
	"Page of Pentacles",
	"Queen of Pentacles",
	"King of Pentacles",
	"Knight of Pentacles",
	"( ) The Fool",
	"(I) The Magician",
	"(II) The High Priestess",
	"(III) The Empress",
	"(IIII) The Emperor",
	"(V) The Pope",
	"(VI) The Lover",
	"(VII) The Chariot",
	"(VIII) Justice",
	"(VIIII) The Hermit",
	"(X) The Wheel of Fortune",
	"(XI) Strength",
	"(XII) The Hanged Man",
	"(XIII) The Nameless Arcanum",
	"(XIIII) Temperance",
	"(XV) The Devil",
	"(XVI) The Tower",
	"(XVII) The Star",
	"(XVIII) The Moon",
	"(XVIIII) The Sun",
	"(XX) Judgement",
	"(XXI) The World"}
