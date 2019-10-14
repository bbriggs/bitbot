package bitbot

import (
	"github.com/whyrusleeping/hellabot"
	"gopkg.in/sorcix/irc.v1"
	//"log"
	"strings"
)

var ChannelPopGaugeTrigger = NamedTrigger{
	ID:   "channelPopGauge",
	Help: "Updates a Prometheus gauge with the value of a channel's population",
	Condition: func(bot *hbot.Bot, m *hbot.Message) bool {
		return m.Command == "JOIN" || m.Command == "PART" || m.Command == "QUIT" || m.Command == "SQUIT"

	},
	Action: func(bot *hbot.Bot, m *hbot.Message) bool {
		if m.Command == "JOIN" {
			b.gauges["channel_pop"].WithLabelValues(m.To).Inc()
			return true
		}
		b.gauges["channel_pop"].WithLabelValues(m.To).Dec()
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
		b.gauges["channel_pop"].WithLabelValues(channel).Set(float64(len(names))) // Exclude ourselves. ChanPopGaugeTrigger catches our own join.
		return true
	},
}
