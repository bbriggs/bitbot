package bitbot

import (
	"fmt"
	"github.com/whyrusleeping/hellabot"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

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

func query(ip string) string {
	url := "http://ipinfo.io/" + ip
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	json, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
