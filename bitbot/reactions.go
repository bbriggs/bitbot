package bitbot

import (
	"strings"

	"github.com/whyrusleeping/hellabot"
)

var ShrugTrigger = NamedTrigger{ //nolint:gochecknoglobals,golint
	ID:   "shrug",
	Help: "Usage: !shrug",
	Condition: func(irc *hbot.Bot, m *hbot.Message) bool {
		return m.Command == "PRIVMSG" && strings.TrimSpace(m.Content) == "!shrug"
	},
	Action: func(irc *hbot.Bot, m *hbot.Message) bool {
		irc.Reply(m, `¯\_(ツ)_/¯`)
		return true
	},
}

var LennyTrigger = NamedTrigger{ //nolint:gochecknoglobals,golint
	ID:   "lennyface",
	Help: "Usage: !lenny",
	Condition: func(irc *hbot.Bot, m *hbot.Message) bool {
		return m.Command == "PRIVMSG" && strings.TrimSpace(m.Content) == "!lenny"
	},
	Action: func(irc *hbot.Bot, m *hbot.Message) bool {
		irc.Reply(m, `( ͡° ͜ʖ ͡°)`)
		return true
	},
}

var WeebTrigger = NamedTrigger{ //nolint:gochecknoglobals,golint
	ID:   "DamnWeebs",
	Help: "Usage: mention uwu",
	Condition: func(irc *hbot.Bot, m *hbot.Message) bool {
		// Thanks d4 for the help!
		return m.Command == "PRIVMSG" && containsOwOLike(m.Content)
	},
	Action: func(irc *hbot.Bot, m *hbot.Message) bool {
		adjectives := []string{ //add more as necessary
			"damn",
			"fucking",
			"degenerate",
			"incurable",
			"disgusting",
			"wonderful",
			"utter",
		}
		adj := adjectives[b.Random.Intn(len(adjectives))]
		reply := m.Name + `, you ` + adj + ` weeb!`
		irc.Reply(m, reply)
		return false
	},
}

// containsOwOLike returns true if the message argument contains owo, uWu, 0w0,
// or any other similarly repeating a character around a w.
// We could have done that with a PCRE regex, but golang uses re2[0] and I'm not
// adding a dependency wrapping PCRE simply to shame weebs, furries, and similar
// creatures.
// [0] https://github.com/google/re2
// Go suxx, brainfuck rewrite when
func containsOwOLike(message string) bool {
	words := strings.Split(message, " ")
	for _, w := range words {
		switch len(w) {
		case 1, 2: //nolint:gomnd
		case 3: //nolint:gomnd
			if (w[0] == w[2]) && (w[1] == 87 || w[1] == 119) {
				return true
			}

			break
		default:
			if owoInWord(w) {
				return true
			}

			break
		}
	}
	return false
}

func owoInWord(word string) bool { // The word contains a [^A-Z]W[^A-Z] or a [A-Z]w[A-Z]
	ws := strings.Split(word, "")
	for x := 1; x < len(word)-1; x++ {
		if ws[x] == "w" && word[x-1] > 64 && word[x-1] < 88 && ws[x-1] == ws[x+1] {
			return true
		} else if ws[x] == "W" && (word[x-1] < 65 || word[x-1] > 87) {
			return true
		}
	}
	return false
}
