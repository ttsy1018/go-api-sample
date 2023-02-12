package main

import (
	"fmt"
	"myapp/db"
	"myapp/models"
	"strconv"

	"golang.org/x/crypto/bcrypt"
	"github.com/jinzhu/gorm"
)


func userSeeds(db *gorm.DB) error {
	for i := 0; i < 10; i++ {
		hash, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
		user := models.User {
			Name: "ユーザー"+strconv.Itoa(i+1),
			Email: "sample"+strconv.Itoa(i+1)+"@gmail.com",
			Password: string(hash),
		}

		if err := db.Create(&user).Error; err != nil {
			fmt.Printf("%+v", err)
		}
	}
	return nil
}


func main() {
	dbCon := db.Init()
	// dBを閉じる
	defer db.CloseDB(dbCon)

	if err := userSeeds(dbCon); err != nil {
		fmt.Printf("%+v", err)
        return
	}
}