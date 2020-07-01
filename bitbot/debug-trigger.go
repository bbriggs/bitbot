package bitbot

import (
	"fmt"
	"github.com/whyrusleeping/hellabot"
	"strings"
)

type DebugRecord struct {
	ID  int `gorm:"unique;AUTO_INCREMENT;PRIMARY_KEY"`
	Str string
}

var ReminderDebugTrigger = NamedTrigger{ //nolint:gochecknoglobals,golint
	ID:   "reminder-debug",
	Help: "You shouldn't see this.",
	Condition: func(irc *hbot.Bot, m *hbot.Message) bool {
		return m.Command == "PRIVMSG" && strings.HasPrefix(m.Trailing, "!db-debug")
	},
	Action: func(irc *hbot.Bot, m *hbot.Message) bool {
		irc.Reply(m, "writing on db")
		d := DebugRecord{
			Str: "aaa",
		}
		if b.DB.NewRecord(d) {
			fmt.Println("PRIMARY KEY IS BLANK")
		}
		b.DB.Debug().Create(&d)
		irc.Reply(m, "written")

		irc.Reply(m, "reading")

		var dr DebugRecord
		b.DB.Debug().Model(&DebugRecord{}).First(&dr)
		fmt.Println(dr.Str)
		irc.Reply(m, fmt.Sprintf("got back : %s", dr.Str))

		return true
	},
	Init: func() error {
		b.DB.Debug().AutoMigrate(&DebugRecord{})
		return nil
	},
}
