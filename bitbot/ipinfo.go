package bitbot

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strings"

	"github.com/whyrusleeping/hellabot"
)

type geoData struct {
	IP       string
	City     string
	Region   string
	Country  string
	Loc      string
	Org      string
	Postal   string
	Timezone string
}

var IPinfoTrigger = NamedTrigger{ //nolint:gochecknoglobals,golint
	ID:   "ipinfo",
	Help: "!ipinfo <valid IP/domain name>",
	Condition: func(irc *hbot.Bot, m *hbot.Message) bool {
		return m.Command == "PRIVMSG" && strings.HasPrefix(m.Content, "!ipinfo")
	},
	Action: func(irc *hbot.Bot, m *hbot.Message) bool {
		var resp string
		cmd := strings.Split(m.Content, " ")
		if len(cmd) > 1 {
			resp = lookup(cmd[1])
		} else {
			resp = "please provide an argument...ya twatsicle"
		}
		irc.Reply(m, resp)
		return true
	},
}

func lookup(arg string) string {
	IP := net.ParseIP(arg)

	if IP == nil { // what is provided isn't an IP
		ips, err := net.LookupIP(arg)
		if err != nil {
			b.Config.Logger.Warn(fmt.Sprintf("Couldn't look up %s", arg))
			return "IP or domain not found"
		}

		IP = ips[0]
	}

	return ipLookup(IP.String())
}

func ipLookup(ip string) string {
	b.Config.Logger.Info(fmt.Sprintf("Looking up %s", ip))

	req, _ := http.NewRequest("GET", fmt.Sprintf("https://ipinfo.io/%s", ip), nil)
	res, err := b.HTTPClient.Do(req)

	if err != nil {
		b.Config.Logger.Warn("IPinfo trigger, couldn't query ipinfo.io", "error", err)
	}

	jsonData, err := ioutil.ReadAll(res.Body)

	if err != nil {
		b.Config.Logger.Warn("IPinfo trigger, couldn't read ipinfo.io answer", "error", err)
	}

	res.Body.Close() //nolint:errcheck,gosec
	return decodeJSON(jsonData)
}

func decodeJSON(encodedJSON []byte) string {
	var (
		ipinfo geoData
		reply  string
	)

	err := json.Unmarshal(encodedJSON, &ipinfo)
	if err != nil {
		b.Config.Logger.Warn("IPinfo trigger, couldn't decode JSON", "error", err)
	}

	if ipinfo.IP == "" {
		reply = "We are being rate limited, try again later or use ipinfo.io yourself."
	} else {
		reply = fmt.Sprintf("\u000312\u001f%s\u000f (%s): in %s, %s, %s (\u000312%s\u000f) postal code: %s, TZ: %s",
			ipinfo.IP,
			ipinfo.Org,
			ipinfo.City,
			ipinfo.Region,
			ipinfo.Country,
			ipinfo.Loc,
			ipinfo.Postal,
			ipinfo.Timezone)
	}
	return reply
}
