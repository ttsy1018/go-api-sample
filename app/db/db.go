package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	gorm "github.com/jinzhu/gorm"
	godotenv "github.com/joho/godotenv"
	"os"
)

// DBの接続
func Init() *gorm.DB {
	// .envを読み込む
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
		return nil
	}

	// 接続情報を定義
	DBMS := "mysql"
	USER := os.Getenv("MYSQL_USER")
	PASS := os.Getenv("MYSQL_PASSWORD")
	HOST := "tcp(" + os.Getenv("MYSQL_HOST") + ":" + os.Getenv("MYSQL_PORT") + ")"
	DBNAME := os.Getenv("MYSQL_DATABASE")

	CONNECT := USER + ":" + PASS + "@" + HOST + "/" + DBNAME

	// DB接続
	db, err := gorm.Open(DBMS, CONNECT)
	if err != nil {
		panic(err)
	}

	return db
}
