package bitbot

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

func (b *Bot) createCounters() {
	totalMessages := promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "bitbot_messages_total",
		Help: "The total number of processed messages",
	},
		[]string{"channel", "user"})
	b.counters["messageCounter"] = totalMessages

	channelPop := promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "bitbot_channel_pop",
		Help: "Number of users in a given channel",
	},
		[]string{"channel"})

	b.gauges["channel_pop"] = channelPop
}
