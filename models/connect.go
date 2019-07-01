package models

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	// postgres import
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

var db *gorm.DB

func init() {

	e := godotenv.Load()
	var username string
	var password string
	var dbName string
	var dbHost string

	if e != nil {
		fmt.Print(e)
		username = os.Getenv("test")
		password = os.Getenv("test143")
		dbName = os.Getenv("contacts")
		dbHost = os.Getenv("db")
	} else {
		username = os.Getenv("POSTGRES_USER")
		password = os.Getenv("POSTGRES_PASSWORD")
		dbName = os.Getenv("POSTGRES_DB")
		dbHost = os.Getenv("db_host")
	}

	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password)
	// fmt.Println(dbURI)

	conn, err := gorm.Open("postgres", dbURI)
	if err != nil {
		fmt.Print(err)
	}

	db = conn
	db.Debug().AutoMigrate(&Account{}, &Contact{})
}

// GetDB returns the database object
func GetDB() *gorm.DB {
	return db
}
