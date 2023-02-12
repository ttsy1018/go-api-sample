package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"os"
)

var (
	DB *gorm.DB // このグローバル変数を使わないとぬるぽになるので注意
	err error
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

	// parseTime: trueでtime.Time型を取得できる
	CONNECT := USER + ":" + PASS + "@" + HOST + "/" + DBNAME + "?parseTime=true"

	// DB接続
	DB, err = gorm.Open(DBMS, CONNECT)
	if err != nil {
		panic(err)
	}

	return DB
}

func CloseDB(db *gorm.DB) {
	sqlDb := db.DB()
	if err = sqlDb.Close(); err != nil {
		panic(err)
	}
}
