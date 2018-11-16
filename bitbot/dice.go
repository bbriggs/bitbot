package bitbot

import (
	"fmt"
	"strings"

	"github.com/justinian/dice"
	"github.com/whyrusleeping/hellabot"
	log "gopkg.in/inconshreveable/log15.v2"
)

const DICE_USAGE = "Usage: [num dice]d[sides](+/-num) (opt: if fudging)"

var RollTrigger = hbot.Trigger{
	func(irc *hbot.Bot, m *hbot.Message) bool {
		return m.Command == "PRIVMSG" && strings.HasPrefix(m.Content, "!roll")
	},
	func(irc *hbot.Bot, m *hbot.Message) bool {
		var resp string
		cmd := strings.Split(m.Content, " ")
		if len(cmd) > 1 {
			resp = roll(cmd[1])
		} else {
			resp = DICE_USAGE
		}
		irc.Reply(m, resp)
		return false
	},
}

// function roll always returns a string to send back to chat
// and logs an error if one appears
func roll(r string) string {
	res, _, err := dice.Roll(r)
	if err != nil {
		log.Error(err.Error())
		return DICE_USAGE
	}
	return fmt.Sprintf("%v", res.Int())
}
