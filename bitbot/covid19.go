package bitbot

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/whyrusleeping/hellabot"
)

type Covid19Data struct {
	Confirmed Confirmed `json:"confirmed"`
	Deaths    Deaths    `json:"deaths"`
	Latest    Latest    `json:"latest"`
	Recovered Recovered `json:"recovered"`
	UpdatedAt string    `json:"updatedAt"`
}

type Coordinates struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type Locations struct {
	Coordinates Coordinates `json:"coordinates"`
	Country     string      `json:"country"`
	CountryCode string      `json:"country_code"`
	Latest      int         `json:"latest"`
	Province    string      `json:"province"`
}

type Confirmed struct {
	Latest    int         `json:"latest"`
	Locations []Locations `json:"locations"`
}

type Deaths struct {
	Latest    int         `json:"latest"`
	Locations []Locations `json:"locations"`
}

type Latest struct {
	Confirmed int `json:"confirmed"`
	Deaths    int `json:"deaths"`
	Recovered int `json:"recovered"`
}

type Recovered struct {
	Latest    int         `json:"latest"`
	Locations []Locations `json:"locations"`
}

var Covid19Trigger = NamedTrigger{
	ID:   "covid19",
	Help: "Fetch stats on Coronavirus by region",
	Condition: func(irc *hbot.Bot, m *hbot.Message) bool {
		return m.Command == "PRIVMSG" && strings.HasPrefix(m.Trailing, "!covid")
	},
	Action: func(irc *hbot.Bot, m *hbot.Message) bool {
		data, ok := getCovid19Data()
		if !ok {
			irc.Reply(m, "Unable to get data at this time")
			return true
		}
		resp := parseCovid19Trigger(strings.Split(m.Trailing, " "), &data)
		irc.Reply(m, resp)
		return true
	},
}

func parseCovid19Trigger(args []string, data *Covid19Data) string {
	// WIP
	var resp string
	switch len(args) {
	case 0:
		resp = "Error: empty request"
	case 1:
		resp = fmt.Sprintf("Total confirmed: %d | Total deaths: %d", data.Confirmed.Latest, data.Deaths.Latest)
	default:
		country, confirmed, dead := covid19StatsByCountryCode(strings.ToUpper(args[1]), data)
		resp = fmt.Sprintf("Stats for %s || Confirmed: %d | Deaths %d", country, confirmed, dead)
	}
	return resp
}

func covid19StatsByCountryCode(cc string, data *Covid19Data) (string, int, int) {
	var (
		confirmed int
		deaths    int
		country   string
	)

	for _, v := range data.Confirmed.Locations {
		log.Println(v.CountryCode)
		if v.CountryCode == cc && v.Province == "nan" {
			country = v.Country
			confirmed = v.Latest
			break
		}
	}

	for _, v := range data.Deaths.Locations {
		if v.CountryCode == cc && v.Province == "nan" {
			deaths = v.Latest
			break
		}
	}
	return country, confirmed, deaths
}

func getCovid19Data() (Covid19Data, bool) {
	var resp Covid19Data

	r, err := http.Get("https://covid19api.herokuapp.com/")
	if err != nil {
		log.Println(err.Error())
		return resp, false
	}
	defer r.Body.Close()
	err = json.NewDecoder(r.Body).Decode(&resp)
	if err != nil {
		log.Println(err)
		return resp, false
	}
	return resp, true
}
