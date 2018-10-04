package bitbot

import (
	"github.com/whyrusleeping/hellabot"
	"github.com/justinian/dice"
	"strings"
	
	log "gopkg.in/inconshreveable/log15.v2"
)
var RollTrigger = hbot.Trigger{
	func(irc *hbot.Bot, m *hbot.Message) bool {
		return m.Command == "PRIVMSG" && strings.TrimSpace(m.Content) == "!roll"
	},
	func(irc *hbot.Bot, m *hbot.Message) bool {
		roll := strings.SplitAfter(m.Content, " ")
		rollResult, blah, err := dice.Roll(roll[1])
		if err != nil {
			log.Error(err.Error())
			return false  // keep processing triggers
		}
		irc.Reply(m, rollResult)
		irc.Reply(m, blah)
		return true
	},
}