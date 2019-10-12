package bitbot

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/bbriggs/quotes/model"
	"github.com/whyrusleeping/hellabot"
)

var raiderQuote = NamedTrigger{
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
		fmt.Println(err.Error())
		return resp, false
	}
	defer r.Body.Close()
	err = json.NewDecoder(r.Body).Decode(&resp)
	if err != nil {
		fmt.Println(err)
	}
	return resp, false
}
