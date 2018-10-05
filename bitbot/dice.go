package bitbot

import (
	"github.com/whyrusleeping/hellabot"
	"github.com/justinian/dice"
	"strings"
	"fmt"
	
	log "gopkg.in/inconshreveable/log15.v2"
)
var RollTrigger = hbot.Trigger{
	func(irc *hbot.Bot, m *hbot.Message) bool {
		return m.Command == "PRIVMSG"
	},
	func(irc *hbot.Bot, m *hbot.Message) bool {
		roll := strings.SplitAfter(m.Content, " ")
		rollResult, _, err := dice.Roll(roll[1])
		if err != nil {
			log.Error(err.Error())
			return false  // keep processing triggers
		}
		resp := fmt.Sprintf("Roll: %s Result: %v", rollResult.Description(), rollResult.Int())
		irc.Reply(m, resp)
		return true
	},
}