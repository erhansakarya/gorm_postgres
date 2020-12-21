package migration

import (
	"github.com/erhansakarya/gorm_postgres/config"
	"github.com/erhansakarya/gorm_postgres/model"
)

// Migrate creates tables
func Migrate() {
	config.GetDB().AutoMigrate(&model.User{})
}
