package bitbot

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/bbriggs/quotes/model"
	"github.com/whyrusleeping/hellabot"
)

var RaiderQuoteTrigger = NamedTrigger{ //nolint:gochecknoglobals,golint
	ID: "raider",
	Condition: func(irc *hbot.Bot, m *hbot.Message) bool {
		return m.Command == "PRIVMSG" && b.Random.Intn(1000) == 1
	},
	Action: func(irc *hbot.Bot, m *hbot.Message) bool {
		r, ok := getQuote("/fallout/raider")
		if ok {
			irc.Reply(m, r.Quote)
		}
		return false
	},
}

func getQuote(endpoint string) (model.Response, bool) {
	var resp model.Response
	r, err := http.Get(fmt.Sprintf("https://quotes.fraq.io%s", endpoint))
	if err != nil {
		b.Config.Logger.Warn("Quote trigger, couldn't get page", "error", err.Error())
		return resp, false
	}
	defer r.Body.Close() //nolint:errcheck,gosec
	err = json.NewDecoder(r.Body).Decode(&resp)
	if err != nil {
		b.Config.Logger.Warn("Quote trigger, couldn't decode page", "error", err.Error())
	}
	return resp, false
}
