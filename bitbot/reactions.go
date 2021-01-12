package bitbot

import (
	"github.com/whyrusleeping/hellabot"
	"strings"
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
func containsOwOLike(message string) bool {
	ws := strings.Split(message, "")
	for x := 0; x < len(ws)-2; x++ {
		if ws[x] == ws[x+2] && (ws[x+1] == "W" || ws[x+1] == "w") {
			return true
		}
	}
	return false
}
