package bitbot

import (
	"fmt"
	"github.com/whyrusleeping/hellabot"
	"log"
)

func (b Bot) isAdmin(m *hbot.Message) bool {
	fullname := fmt.Sprintf("%s@%s", m.Name, m.Host)
	log.Println(fullname)
	for _, u := range b.Config.Admins.Permitted {
		log.Println(u)
		if u == fullname {
			return true
		}
	}
	return false
}
