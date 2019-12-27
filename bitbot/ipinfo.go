package bitbot

import (
	"encoding/json"
	"fmt"
	"github.com/whyrusleeping/hellabot"
	"io/ioutil"
	"log"
	"net/http"
)

type Geo_data struct {
	IP       string
	Hostname string
	City     string
	Region   string
	Country  string
	Loc      string
	Org      string
	Postal   string
	Timezone string
	Readme   string
}

var IPinfoTrigger = NamedTrigger{
	ID:   "ipinfo",
	Help: "!ipinfo <valid IP>",
	Condition: func(irc *hbot.Bot, m *hbot.Message) bool {
		return m.Command == "PRIVMSG" && strings.HasPrefix(m.Content, "!ipinfo")
	},
	Action: func(irc *hbot.Bot, m *hbot.Message) bool {
		var resp string
		cmd := strings.Split(m.Content, " ")
		if len(cmd) > 1 {
			resp = query(cmd[1])
		} else {
			resp = "please provide an ip...ya twatsicle"
		}
		irc.Reply(m, resp)
		return true
	},
}

func decode_json(b []byte) string {
	var ipinfo Geo_data
	var reply string
	err := json.Unmarshal(b, &ipinfo)
	if err != nil {
		log.Println(err)
	}
	reply = fmt.Sprintf("ip: %s\nhostname: %s\ncity: %s\nregion: %s\ncountry: %s\ncoords: %s\norg: %s\npostal: %s\ntimezone: %s", ipinfo.IP, ipinfo.Hostname, ipinfo.City, ipinfo.Region, ipinfo.Country, ipinfo.Loc, ipinfo.Org, ipinfo.Postal, ipinfo.Timezone)
	return reply

}

func query(ip string) string {
	url := "http://ipinfo.io/" + ip
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	jsonData, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	return decode_json(jsonData)
}
