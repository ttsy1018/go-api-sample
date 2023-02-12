package main

import (
	"myapp/db"
	"github.com/jinzhu/gorm"
	"myapp/models"
)

func migrate(dbCon *gorm.DB) {
	dbCon.AutoMigrate(
		&models.Todo{},
		&models.User{},
	)
}

func main() {
	dbCon := db.Init()

	defer db.CloseDB(dbCon)

	migrate(dbCon)
}