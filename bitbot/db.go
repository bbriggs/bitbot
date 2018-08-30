package bitbot

import (
	bolt "go.etcd.io/bbolt"
)

func newDB() (*bolt.DB, error) {
	db, err := bolt.Open(".bolt.db", 0666, nil)
	return db, err
}
