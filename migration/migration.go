package migration

import (
	"fmt"

	"github.com/erhansakarya/gorm_sqlite/model"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq" // postgresql driver
)

var db *gorm.DB

// InitialMigration creates tables and connect to db
func InitialMigration() {
	db, err := gorm.Open("postgres", "test.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect to the db")
	}
	defer db.Close()

	db.AutoMigrate(&model.User{})
}
