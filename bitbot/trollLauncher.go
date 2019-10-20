package bitbot

import (
	"fmt"
	hbot "github.com/whyrusleeping/hellabot"
	"math/rand"
	"strings"
)

var TrollLauncherTrigger = NamedTrigger{
	ID:   "troll",
	Help: "Launches a random number of trolls for a random amount of damage. Usage: !troll <nick>",
	Condition: func(irc *hbot.Bot, m *hbot.Message) bool {
		return m.Command == "PRIVMSG" && strings.HasPrefix(m.Content, "!troll ")
	},
	Action: func(irc *hbot.Bot, m *hbot.Message) bool {
		var damage_type = [13]string{"bludgeoning", "piercing", "slashing", "cold", "fire", "acid", "poison",
			"psychic", "necrotic", "radiant", "lightning", "thunder", "force"}
		cmd := strings.Split(m.Content, " ")
		if len(cmd) < 2 {
			irc.Reply(m, "Usage: !troll <nick>")
			return true
		}
		var nick = cmd[1]
		var trolls = rand.Intn(10)
		if trolls == 0 {
			irc.Reply(m, "The troll launcher malfunctioned.")
		} else {
			reply := fmt.Sprintf("Firing %d trolls at %s! You take %d points of %s damage!", trolls, nick,
				rand.Intn(20), damage_type[rand.Intn(12)])
			irc.Reply(m, reply)
		}
		return true
	},
}
