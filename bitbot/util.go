package bitbot

import (
	"encoding/binary"
	"fmt"
	"github.com/whyrusleeping/hellabot"
	"gopkg.in/sorcix/irc.v1"
	"time"
)

// NamedTrigger is a local re-implementation of hbot.Trigger to support unique names
type NamedTrigger struct {
	ID        string // Name of trigger, to use used to registering, searching, and deregistering
	Help      string // Help text
	Condition func(*hbot.Bot, *hbot.Message) bool
	Action    func(*hbot.Bot, *hbot.Message) bool
}

// Name satisfies the hbot.Handler interface
func (t NamedTrigger) Name() string {
	return t.ID
}

// Handle executes the trigger action if the condition is satisfied
func (t NamedTrigger) Handle(b *hbot.Bot, m *hbot.Message) bool {
	if !t.Condition(b, m) {
		return false
	}
	return t.Action(b, m)
}

// ACL defines access lists the bot may use to check Authorization to use a trigger
type ACL struct {
	// Defines users explicitly allowed in this ACL
	Permitted []string
	// Defines users explicitly rejected by this ACL
	Rejected []string
}

// ListTriggers gets all trigger IDs currently registered to the bot
func (b *Bot) ListTriggers() []string {
	var triggers []string
	b.triggerMutex.RLock()
	defer b.triggerMutex.RUnlock()

	for k, _ := range b.triggers {
		triggers = append(triggers, k)
	}
	return triggers
}

func int64ToByte(i int64) []byte {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, uint64(i))
	return b
}

func byteToInt64(b []byte) int64 {
	i := int64(binary.LittleEndian.Uint64(b))
	return i
}

func fmtDuration(d time.Duration) string {
	day := d / time.Hour * 24
	d -= day * time.Hour * 24

	h := d / time.Hour
	d -= h * time.Hour

	m := d / time.Minute
	d -= m * time.Minute

	return fmt.Sprintf("%02d days, %02d hours, %02d minutes", day, h, m)
}

func makeMockMessage(nick, message string) *hbot.Message {
	return &hbot.Message{
		&irc.Message{
			&irc.Prefix{
				nick,
				"",
				"",
			},
			"PRIVMSG",
			[]string{},
			message,
			true,
		},
		message,
		time.Now(),
		"",
		nick,
	}
}

func makeMockBot(nick string) *hbot.Bot {
	return &hbot.Bot{
		Nick: nick,
		Host: "foo",
	}
}

func partChannel(irc *hbot.Bot, ch, msg string) {
	partString := "PART " + ch + " " + msg
	irc.Send(partString)
}
