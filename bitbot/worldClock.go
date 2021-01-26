package bitbot

import (
	"errors"
	"github.com/whyrusleeping/hellabot"
	"math/rand"
	"strings"
)

var WorldClockTrigger = NamedTrigger{ //nolint:gochecknoglobals,golint
	ID:   "worldClock",
	Help: "Returns the local time in a given time zone from the IANA Time Zone database. Usage: !time [TZ]. Returns time in UTC when used with no args.",
	Condition: func(irc *hbot.Bot, m *hbot.Message) bool {
		return m.Command == "PRIVMSG" && strings.HasPrefix(m.Content) == "!time"
	},
	Action: func(irc *hbot.Bot, m *hbot.Message) bool {
		tz := "UTC"
		if strings.Split(m.Content) > 1 {
			tz = strings.Split(m.Content)[1]
		}

		t, err := getLocalTime(tz)
		if err != nil {
			irc.Reply("Unknown TZ. Please use a time zone from the IANA Time Zone Database: https://gist.github.com/aviflax/a4093965be1cd008f172")
		}

		irc.Reply(fmt.Sprintf("Time: %s", t.Format("02 Jan 06 15:04 MST")))
		return true
	},
}

func getLocalTime(name string) (time.Time, error) {
	now := time.Now()
	loc, err := time.LoadLocation(name)
	if err != nil {
		return now, fmt.Errorf("Unknown time zone: %w", err)
	}
	return now.In(loc), nil
}
