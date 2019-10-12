package bitbot

import (
	"fmt"
	"strings"
	"time"

	"github.com/whyrusleeping/hellabot"
	bolt "go.etcd.io/bbolt"
	log "gopkg.in/inconshreveable/log15.v2"
)

var TrackIdleUsers = NamedTrigger{
	ID:   "trackIdleUsers",
	Help: "Passive, non-interactive, experimental trigger. Monitors time since last activity for all users in channel. Works like /whois.",
	Condition: func(irc *hbot.Bot, m *hbot.Message) bool {
		return m.Command == "PRIVMSG"
	},
	Action: func(irc *hbot.Bot, m *hbot.Message) bool {
		err := b.TrackIdleUsers(m)
		if err != nil {
			log.Error(err.Error())
		}
		return false // keep processing triggers
	},
}

func (b Bot) TrackIdleUsers(m *hbot.Message) error {
	err := b.DB.Update(func(tx *bolt.Tx) error {
		now := int64ToByte(time.Now().Unix())
		bucket, err := tx.CreateBucketIfNotExists([]byte(m.From))
		if err != nil {
			return err
		}
		err = bucket.Put([]byte("last_message_time"), []byte(now))
		return err
	})
	return err
}

var ReportIdleUsers = NamedTrigger{
	ID: "reportIdleUsers",
	Condition: func(irc *hbot.Bot, m *hbot.Message) bool {
		return m.Command == "PRIVMSG" && strings.HasPrefix(m.Content, "!idle")
	},
	Action: func(irc *hbot.Bot, m *hbot.Message) bool {
		args := strings.Split(m.Content, " ")
		if len(args) < 2 {
			irc.Reply(m, "Please specify a nick to lookup")
			return true
		}
		log.Info(fmt.Sprintf("Getting idle time for %s", args[1]))
		report, err := b.GetUserIdleTime(args[1])
		if err != nil {
			irc.Reply(m, "Unable to lookup idle time for that user")
			return true
		}
		if report == "" {
			irc.Reply(m, "I have not seen that user yet.")
			return true
		}
		irc.Reply(m, fmt.Sprintf("%s has been idle for %s", args[1], report))
		return true
	},
}

func (b Bot) GetUserIdleTime(nick string) (string, error) {
	var val []byte
	err := b.DB.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(nick))
		c := bucket.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			log.Info(fmt.Sprintf("key=%s, value=%s\n", k, v))
		}
		val = bucket.Get([]byte("last_message_time"))
		return nil
	})
	if err != nil {
		return "", err
	}
	if val == nil { // bbolt returns nil on nonexistent keys
		return "", err
	}
	i := byteToInt64(val)
	ts := time.Unix(i, 0)
	elapsed := fmtDuration(time.Since(ts))
	return elapsed, err
}
