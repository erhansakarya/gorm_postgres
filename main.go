package main

import (
	"fmt"

	"github.com/erhansakarya/gorm_postgres/config"
	"github.com/erhansakarya/gorm_postgres/migration"
	"github.com/erhansakarya/gorm_postgres/router"
)

func init() {
	config.Init()
	migration.Migrate()
}

func main() {
	fmt.Println("Go ORM")

	router.HandleRequests()
}
