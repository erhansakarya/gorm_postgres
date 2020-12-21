package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq" // postgresql driver
)

// DB handler
var db *gorm.DB

// Init connects to db
func Init() {
	var err error
	db, err = gorm.Open("postgres", "test.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect to the db")
	}
}

// GetDB returns DB handler
func GetDB() *gorm.DB {
	return db
}
