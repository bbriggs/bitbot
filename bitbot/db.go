package bitbot

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type DB struct {
	DB *gorm.DB
}

func newDB(c DBConfig) (*gorm.DB, error) {
	connString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", c.Host, c.Port, c.User, c.Name, c.Pass, c.SSLMode)
	db, err := gorm.Open("postgres", connString)
	return db, err
}
