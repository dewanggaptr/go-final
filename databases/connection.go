package database

import (
	"final-project/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = "localhost"
	user     = "putra"
	password = "test123"
	dbPort   = "5432"
	dbname   = "go_social"
	db       *gorm.DB
)

func DBInit() {
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, user, password, dbname, dbPort)

	// url := "root:@tcp(localhost:5432)/mini-shop"

	db, err := gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		log.Fatal("error connection to dataase :", err)
	}

	fmt.Println("connecting to database")

	db.Debug().AutoMigrate(&models.User{}, &models.Post{})

	// DB = db
	// return db, nil
}

func GetDB() *gorm.DB {
	return db
}
