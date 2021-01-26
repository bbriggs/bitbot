package bitbot

import (
	"fmt"
	"strings"
	"time"

	"github.com/whyrusleeping/hellabot"
)

// WorldClockTrigger sends back the time in a specified timezone.
var WorldClockTrigger = NamedTrigger{ //nolint:gochecknoglobals,golint
	ID:   "worldClock",
	Help: "Returns the local time in a given time zone from the IANA Time Zone database. Usage: !time [TZ]. Returns time in UTC when used with no args.",
	Condition: func(irc *hbot.Bot, m *hbot.Message) bool {
		return m.Command == "PRIVMSG" && strings.HasPrefix(m.Content, "!time")
	},
	Action: func(irc *hbot.Bot, m *hbot.Message) bool {
		tz := "UTC"
		if len(strings.Split(m.Content, " ")) > 1 {
			tz = strings.Split(m.Content, " ")[1]
		}

		t, err := getLocalTime(tz)
		if err != nil {
			irc.Reply(m,
			"Unknown TZ, assuming UTC. Please use a time zone from the IANA Time Zone Database: https://gist.github.com/aviflax/a4093965be1cd008f172")
		}

		irc.Reply(m, fmt.Sprintf("Time: %s", t.Format("02 Jan 06 15:04 MST")))
		return true
	},
}

func getLocalTime(name string) (time.Time, error) {
	now := time.Now()
	loc, err := time.LoadLocation(name)
	if err != nil {
		err = fmt.Errorf("unknown time zone: %w", err)
		loc, _ = time.LoadLocation("UTC")
	}
	return now.In(loc), err
}
