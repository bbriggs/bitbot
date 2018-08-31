package bitbot

import (
	bolt "go.etcd.io/bbolt"
)

type DB struct {
	DB *bolt.DB
}


func newDB() (*bolt.DB, error) {
	db, err := bolt.Open(".bolt.db", 0666, nil)
	return db, err
}
