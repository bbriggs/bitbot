package bitbot

import (
	"time"

	"github.com/whyrusleeping/hellabot"
	bolt "go.etcd.io/bbolt"
	log "gopkg.in/inconshreveable/log15.v2"
)

type DB struct {
	DB *bolt.DB
}


func newDB() (*bolt.DB, error) {
	db, err := bolt.Open(".bolt.db", 0666, nil)
	return db, err
}

var TrackIdleUsers = hbot.Trigger{
	func(irc *hbot.Bot, m *hbot.Message) bool {
		return m.Command == "PRIVMSG"
	},
	func(irc *hbot.Bot, m *hbot.Message) bool {
		err := b.TrackIdleUsers(m)
		if err != nil {
			log.Error(err.Error())
		}
		return false  // keep processing triggers
	},
}

func (b Bot) TrackIdleUsers(m *hbot.Message) error {
	err := b.DB.Update(func(tx *bolt.Tx) error {
		now := int64ToByte(time.Now().Unix())
		bucket, err := tx.CreateBucketIfNotExists([]byte(m.From))
		if err != nil {
			return err
		}
		err = bucket.Put([]byte("last_message"), []byte(now))
		return err
	})
	return err
}
