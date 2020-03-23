package bitbot

import (
	"fmt"
	"strings"
	"time"

	"github.com/markcheno/go-quote"
	"github.com/whyrusleeping/hellabot"
	log "gopkg.in/inconshreveable/log15.v2"
)

var StockTrigger = NamedTrigger{
	ID:   "stockQuotes",
	Help: "$SYMBOL to get current price for a stock symbol",
	Condition: func(irc *hbot.Bot, m *hbot.Message) bool {
		return m.Command == "PRIVMSG" && strings.HasPrefix(m.Content, "$")
	},
	Action: func(irc *hbot.Bot, m *hbot.Message) bool {
		var q quote.Quote
		var err error
		symbol := strings.TrimPrefix(m.Content, "$")
		if len(symbol) > 1 {
			q, err = getStockSymbol(symbol)
		} else {
			log.Error("Error: Empty symbol provided to symbol lookup")
			return true
		}

		if err != nil {
			log.Error("Got an error during quote lookup!")
			irc.Reply(m, fmt.Sprintf("Got an error during quote lookup: %s", err.Error()))
			return true
		}

		replyWithQuote(irc, m, q)
		return true
	},
}

func getStockSymbol(symbol string) (quote.Quote, error) {
	var (
		q   quote.Quote
		err error
	)

	today := time.Now().Format("2006-01-02")
	fmt.Println(today)
	// has to be a 24 hour span over a trading day. error handling is overrated
	q, err = quote.NewQuoteFromYahoo(symbol, "2019-10-30", "2019-10-31", quote.Daily, true)
	if err != nil {
		return q, err
	}
	log.Info(fmt.Sprintf("%+v", q))
	return q, err
}

func replyWithQuote(irc *hbot.Bot, m *hbot.Message, q quote.Quote) {
	irc.Reply(m, fmt.Sprintf("%s open: \t%2f", q.Symbol, q.Open[0]))
	irc.Reply(m, fmt.Sprintf("%s high: \t%2f", q.Symbol, q.High[0]))
	irc.Reply(m, fmt.Sprintf("%s low: \t%2f", q.Symbol, q.Low[0]))
	irc.Reply(m, fmt.Sprintf("%s close: \t%2f", q.Symbol, q.Close[0]))
}
