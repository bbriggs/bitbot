package bitbot

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/whyrusleeping/hellabot"
)

type covidData struct {
	Country  string
	Infected int
	Deceased int
}

// Covid19Trigger This trigger reads data from https://apify.com/covid-19 and formats it.
var Covid19Trigger = NamedTrigger{
	ID:   "covid19",
	Help: "!covid19 [<Country Name>]",
	Condition: func(irc *hbot.Bot, m *hbot.Message) bool {
		return m.Command == "PRIVMSG" && strings.HasPrefix(m.Trailing, "!covid19")
	},
	Action: func(irc *hbot.Bot, m *hbot.Message) bool {
		data, err := getCovidData()
		if err != nil {
			irc.Reply(m, "Couldn't get data")
		}

		cs := strings.Split(m.Content, " ")
		if len(cs) == 1 {
			irc.Reply(m, getCovidGlobalStats(data))
		}
		if len(cs) >= 2 {
			countryName := strings.Join(cs[1:], " ")
			irc.Reply(m, getCovidCountryStats(data, countryName))
		}
		return true
	},
}

func (c *covidData) Add(c2 covidData) {
	c.Infected += c2.Infected
	c.Deceased += c2.Deceased
}

func (c *covidData) String() string {
	percent := 0
	if c.Infected > 0 {
		percent = c.Deceased * 100 / c.Infected
	}
	return fmt.Sprintf(
		"\x02\x1f%s covid stats today:\x0f \x0304%d dead\x0f out of \x0303%d infected\x0f (\x1f%d%%\x0f mortality)",
		c.Country,
		c.Deceased,
		c.Infected,
		percent,
	)
}

func getCovidGlobalStats(stats []covidData) string {
	g := covidData{
		Country:  "World",
		Infected: 0,
		Deceased: 0,
	}

	for _, c := range stats {
		g.Add(c)
	}

	return g.String()
}

func getCovidCountryStats(stats []covidData, country string) string {
	for _, c := range stats {
		if c.Country == country {
			return c.String()
		}
	}

	return "No data on this country. Data source: https://apify.com/covid-19"
}

func getCovidData() ([]covidData, error) {
	// TODO take care of fields not always filled (recovered and tested)
	var d []covidData

	req, _ := http.NewRequest("GET", //nolint:noctx
		"https://api.apify.com/v2/key-value-stores/tVaYRsPHLjNdNBu7S/records/LATEST?disableRedirect=true",
		nil)
	r, err := b.HTTPClient.Do(req)
	if err != nil {
		b.Config.Logger.Info("covid19: Couldn't fetch data", "err", err)
		return d, err
	}
	defer r.Body.Close() //nolint:errcheck

	a, _ := ioutil.ReadAll(r.Body)
	err = json.Unmarshal(a, &d)
	if err != nil {
		b.Config.Logger.Info("covid19: Couldn't parse data", "err", err)
		return nil, err
	}

	return d, err
}
