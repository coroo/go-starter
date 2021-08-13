package config

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	_ "github.com/joho/godotenv/autoload"
)

func ConnectDB() (c *gorm.DB, err error) {
	DB_CONNECTION := os.Getenv("DB_CONNECTION_ODS")
	DB_HOST := os.Getenv("DB_HOST_ODS")
	DB_PORT := os.Getenv("DB_PORT_ODS")
	DB_DATABASE := os.Getenv("DB_DATABASE_ODS")
	DB_USERNAME := os.Getenv("DB_USERNAME_ODS")
	DB_PASSWORD := os.Getenv("DB_PASSWORD_ODS")

	DB_TEST := os.Getenv("DB_TEST")
	DB_DETAIL := DB_USERNAME + ":" + DB_PASSWORD + "@tcp(" + DB_HOST + ":" + DB_PORT + ")/" + DB_DATABASE + "?parseTime=true"
	if DB_TEST != "" {
		DB_CONNECTION = "sqlite3"
		DB_DETAIL = DB_TEST
	}

	conn, err := gorm.Open(DB_CONNECTION, DB_DETAIL)
	if err != nil || conn == nil {
		fmt.Println("Error connecting to DB")
		fmt.Println(err.Error())
	}
	return conn, err
}

func ConnectDBSY() (c *gorm.DB, err error) {
	DB_CONNECTION := os.Getenv("DB_CONNECTION_SY")
	DB_HOST := os.Getenv("DB_HOST_SY")
	DB_PORT := os.Getenv("DB_PORT_SY")
	DB_DATABASE := os.Getenv("DB_DATABASE_SY")
	DB_USERNAME := os.Getenv("DB_USERNAME_SY")
	DB_PASSWORD := os.Getenv("DB_PASSWORD_SY")

	conn, err := gorm.Open(DB_CONNECTION, DB_USERNAME+":"+DB_PASSWORD+"@tcp("+DB_HOST+":"+DB_PORT+")/"+DB_DATABASE+"?parseTime=true")
	if err != nil || conn == nil {
		fmt.Println("Error connecting to DB")
		fmt.Println(err.Error())
	}
	return conn, err
}
