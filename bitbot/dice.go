package bitbot

import (
	"fmt"
	"strings"

	"github.com/justinian/dice"

	log "gopkg.in/inconshreveable/log15.v2"
)

var RollTrigger = hbot.Trigger{
	func(irc *hbot.Bot, m *hbot.Message) bool {
		return m.Command == "PRIVMSG" && strings.TrimSpace(m.Content) == "!roll"
	},
	func(irc *hbot.Bot, m *hbot.Message) bool {
		roll := strings.SplitAfter(m.Content, " ")
		rollResult, _, err := dice.Roll(roll[1])
		if err != nil {
			log.Error(err.Error())
			irc.Reply(m, "Usage: [num dice]d[sides](+/-num) (opt: if fudging)")
			return false // keep processing triggers
		}
		if roll[0] != "!roll" {
			log.Error(err.Error())
			irc.Reply(m, "You didn't call the command correctly.")
			return false
		}

		resp := fmt.Sprintf("%v", rollResult.Int())
		irc.Reply(m, resp)
		return false
	},
}
