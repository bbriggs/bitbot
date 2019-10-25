package bitbot

import (
	"fmt"
	"github.com/whyrusleeping/hellabot"
	"gopkg.in/sorcix/irc.v1"
	"time"
)

var OperLogin = hbot.Trigger{
	func(bot *hbot.Bot, m *hbot.Message) bool {
		return m.Command == irc.RPL_MYINFO // message type 004
	},

	func(bot *hbot.Bot, m *hbot.Message) bool {
		ns, ok := b.NickservLogin()
		if ok {
			bot.Msg("NickServ", ns)
		}
		time.Sleep(5 * time.Second)

		op, ok := b.OperLogin()
		if ok {
			bot.Send(op)
			time.Sleep(5 * time.Second)
			b.GetOper()
		}

		return true
	},
}

func (b Bot) OperLogin() (string, bool) {
	//login := fmt.Sprintf("%+v", b.Config)
	if b.Config.OperUser == "" || b.Config.OperPass == "" {
		return "", false
	}
	login := fmt.Sprintf("OPER %s %s", b.Config.OperUser, b.Config.OperPass)
	return login, true
}

func (b Bot) NickservLogin() (string, bool) {
	if b.Config.NickservPass == "" {
		return "", false
	}
	login := fmt.Sprintf("IDENTIFY %s", b.Bot.Nick, b.Config.NickservPass)
	return login, true
}

func (b Bot) GetOper() {
	for _, channel := range b.Config.Channels {
		b.Bot.ChMode(b.Bot.Nick, channel, "+o")
	}
}
