package bitbot

import (
	"io/ioutil"
	"net/http"

	"github.com/whyrusleeping/hellabot"
)

var DadJokeTrigger = NamedTrigger{ //nolint:gochecknoglobals,golint
	ID:   "dadjoke",
	Help: "get a icanhazdadjoke.com joke.",
	Condition: func(itc *hbot.Bot, m *hbot.Message) bool {
		return m.Command == "PRIVMSG" && m.Trailing() == "!dadjoke"
	},
	Action: func(irc *hbot.Bot, m *hbot.Message) bool {
		req, _ := http.NewRequest("GET", "https://icanhazdadjoke.com/", nil)
		req.Header.Set("Accept", "text/plain")

		resp, err := b.HTTPClient.Do(req)
		if err != nil {
			b.Config.Logger.Warn("Couldn't get dad joke from API", "err", err)
			irc.Reply(m, "Why didn't you get a dad joke? Because the API was unavailable")
			return true
		}

		answer, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			b.Config.Logger.Warn("Couldn't get dad joke from API answer", "err", err)
			irc.Reply(m, "Why didn't you get a dad joke? Because we are getting rate limited")
			return true
		}

		irc.Reply(m, string(answer))

		return true
	},
}
