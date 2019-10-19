package bitbot

import (
	"math/rand"
	"strings"
	"net/http"

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
		markovAdd(b.mChain)
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

var MarkovInitTrigger = NamedTrigger{
	ID: "markovInit",
	Help: "Resets markov chain to a fresh chain, or bootstraps it with sample texts. Usage: !markov reset, !markov init",
	Condition: func(irc *hbot.Bot, m *hbot.Message) bool {
		return m.Command == "PRIVMSG" && strings.HasPrefix(m.Content == "!markov ")
	},
	Action: func(irc *hbot.bot, m *hbot.Message) bool {
		cmd := strings.Split(m.Content, " ")
		if len(cmd) < 2 {
			return false
		}
		switch cmd[1]{
		case "reset":
			// do stuff
			b.mChain = gomarkov.NewChain(1)
			return true
		case "init":
			// do other stuff
			b.mChain = gomarkov.NewChain(1)
			markovInit(b.mChain)
			return true
		default:
			// didn't recognize the subcommands
			return false
		}
	}
	
}

func generateBabble(chain *gomarkov.Chain) string {
	tokens := []string{gomarkov.StartToken}
	for tokens[len(tokens)-1] != gomarkov.EndToken {
		next, _ := chain.Generate(tokens[(len(tokens) - 1):])
		tokens = append(tokens, next)
	}
	return strings.Join(tokens[1:len(tokens)-1], " ")
}

func markovInit(chain *gomarkov.Chain) string {
	//todo
}

// wrapper for b.mChain.Add that includes file lock/unlock
func markovAdd(text string, chain *gomarkov.Chain) {
	b.markovMutex.Lock()
	b.mChain.Add(strings.Split(text, " "))
	b.markovMutex.Unlock()
}
