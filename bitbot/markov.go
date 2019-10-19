package bitbot

import (
	"math/rand"
	"strings"

	"github.com/mb-14/gomarkov"
	"github.com/whyrusleeping/hellabot"
)

var MarkovTrainerTrigger = NamedTrigger{
	ID:   "markovTrainer",
	Help: "Incrementally trains bitbot's markov model on every new privmsg",
	Condition: func(irc *hbot.Bot, m *hbot.Message) bool {
		return m.Command == "PRIVMSG"
	},
	Action: func(irc *hbot.Bot, m *hbot.Message) bool {
		if b.mChain == nil {
			// Initialize markov chain
			b.mChain = gomarkov.NewChain(1)
		}
		b.markovMutex.Lock()
		b.mChain.Add(strings.Split(m.Content, " "))
		b.markovMutex.Unlock()
		return false
	},
}

var MarkovResponseTrigger = NamedTrigger{
	ID:   "markovResponse",
	Help: "Returns a randomly generated markov string. Usage: !babble <seed words>",
	Condition: func(irc *hbot.Bot, m *hbot.Message) bool {
		return m.Command == "PRIVMSG" && (m.Content == "!babble" || rand.Intn(100) == 0) // 1% chance
	},
	Action: func(irc *hbot.Bot, m *hbot.Message) bool {
		irc.Reply(m, generateBabble(b.mChain))
		return false
	},
}

func generateBabble(chain *gomarkov.Chain) string {
	tokens := []string{gomarkov.StartToken}
	for tokens[len(tokens)-1] != gomarkov.EndToken {
		next, _ := chain.Generate(tokens[(len(tokens) - 1):])
		tokens = append(tokens, next)
	}
	return strings.Join(tokens[1:len(tokens)-1], " ")
}
