package bitbot

import (
	"log"
	"strconv"
	"strings"

	"github.com/whyrusleeping/hellabot"
	"gopkg.in/sorcix/irc.v1"
)

var ChannelPopGaugeTrigger = NamedTrigger{
	ID:   "channelPopGauge",
	Help: "Updates a Prometheus gauge with the value of a channel's population",
	Condition: func(bot *hbot.Bot, m *hbot.Message) bool {
		return m.Command == "JOIN" || m.Command == "PART" || m.Command == "QUIT"

	},
	Action: func(bot *hbot.Bot, m *hbot.Message) bool {
		switch m.Command {
		case "JOIN":
			b.gauges["channel_pop"].WithLabelValues(m.To).Inc()
		case "QUIT":
			for _, c := range b.Bot.Channels {
				bot.Send("LIST " + c)
			}
		default:
			b.gauges["channel_pop"].WithLabelValues(m.To).Dec()
		}

		return true
	},
}

var HandleListReplyTrigger = NamedTrigger{
	ID:   "handleListReply",
	Help: "Sets the gauge for a channel pop when an RPL_LIST command is detected",
	Condition: func(bot *hbot.Bot, m *hbot.Message) bool {
		return m.Command == irc.RPL_LIST
	},
	Action: func(bot *hbot.Bot, m *hbot.Message) bool {
		log.Println("List Reply Received")
		channel := m.Params[1]
		pop, err := strconv.Atoi(m.Params[2])
		if err != nil {
			log.Println("RPL_LIST returned invalid reply")
		}
		b.gauges["channel_pop"].WithLabelValues(channel).Set(float64(pop))
		return true
	},
}

var SetChanPopGaugeTrigger = NamedTrigger{
	ID:   "setChannelPopGauge",
	Help: "Sets the gauge for a channel's population when the bot joins",
	Condition: func(bot *hbot.Bot, m *hbot.Message) bool {
		return m.Command == irc.RPL_NAMREPLY
	},
	Action: func(bot *hbot.Bot, m *hbot.Message) bool {
		names := strings.Split(strings.TrimSpace(m.Content), " ")
		channel := m.Params[len(m.Params)-1]
		// log.Printf("Setting channel_pop to %d\n\n", len(names))
		b.gauges["channel_pop"].WithLabelValues(channel).Set(float64(len(names)))
		return true
	},
}
