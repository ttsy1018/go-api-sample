package main

import (
	"myapp/db"
	"myapp/models"
	"github.com/jinzhu/gorm"
	"strconv"
	"fmt"
)

func todoSeeds(db *gorm.DB) error {
	for i := 0; i < 10; i++ {
		var userId int
		if i < 5 {
			userId = 1
		} else {
			userId = 2
		}
		todo := models.Todo {
			Title: "タイトル"+strconv.Itoa(i+1),
			Comment: "コメント"+strconv.Itoa(i+1),
			UserID: userId,
		}

		if err := db.Create(&todo).Error; err != nil {
			fmt.Printf("%+v", err)
		}
	}
	return nil
}

func main() {
	dbCon := db.Init()
	// dBを閉じる
	defer db.CloseDB(dbCon)

	// if err := userSeeds(dbCon); err != nil {
	// 	fmt.Printf("%+v", err)
    //     return
	// }
	
	if err := todoSeeds(dbCon); err != nil {
		fmt.Printf("%+v", err)
        return
	}
}