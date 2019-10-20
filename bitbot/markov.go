package bitbot

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
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
		markovAdd(m.Content, b.mChain)
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
	ID:   "markovInit",
	Help: "Resets markov chain to a fresh chain, or bootstraps it with sample texts. Usage: !markov reset, !markov init",
	Condition: func(irc *hbot.Bot, m *hbot.Message) bool {
		return m.Command == "PRIVMSG" && strings.HasPrefix(m.Content, "!markov ")
	},
	Action: func(irc *hbot.Bot, m *hbot.Message) bool {
		cmd := strings.Split(m.Content, " ")
		if len(cmd) < 2 {
			return false
		}
		switch cmd[1] {
		case "reset":
			// do stuff
			b.mChain = gomarkov.NewChain(1)
			return true
		case "init":
			// do other stuff
			b.mChain = gomarkov.NewChain(1)
			return markovInit(b.mChain)
		default:
			// didn't recognize the subcommands
			return false
		}
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

func markovInit(chain *gomarkov.Chain) bool {
	var sources = [3]string{"https://gist.githubusercontent.com/bbriggs/60e907f3571a1ca7c41cd99f78052d78/raw/fe6d0bd96ee97c9b5df2794ae683d24a404b4433/bible.txt",
		"https://gist.githubusercontent.com/bbriggs/f63340a3ed1a1439b6f3f8d619eacac1/raw/1f363d500226c55bab735fe59074f06721348546/world_factbook.txt",
		"https://gist.githubusercontent.com/parsec/2f4d4edf55336c0a2994cfcf951a8ea7/raw/4b66c99f1879b927ebc2b2ffb8fdd39dc9a4f7d2/SnwCrsh"}

	i := 0
	for i <= len(sources) {
		resp, err := http.Get(sources[i])
		if err != nil {
			fmt.Println(err)
			return false
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		bodyString := string(body)
		markovAdd(bodyString, chain)
	}
	return true
}

// wrapper for b.mChain.Add that includes file lock/unlock
func markovAdd(text string, chain *gomarkov.Chain) {
	b.markovMutex.Lock()
	b.mChain.Add(strings.Split(text, " "))
	b.markovMutex.Unlock()

}
