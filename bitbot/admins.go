package bitbot

import (
	"fmt"

	"github.com/whyrusleeping/hellabot"
)

func (b Bot) isAdmin(m *hbot.Message) bool {
	fullname := fmt.Sprintf("%s@%s", m.Name, m.Host)
	b.Config.Logger.Info("Someone requested admin", "fullname", fullname)
	for _, u := range b.Config.Admins.Permitted {
		if u == fullname {
			b.Config.Logger.Info("Admin request granted", "username", u)
			return true
		}
	}

	b.Config.Logger.Warn("Admin request declined", "username", fullname)
	return false
}
