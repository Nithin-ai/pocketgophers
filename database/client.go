package database

import (
	"log"
	"shoppingcartcrud/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB
var err error

func DbConnection(connectionString string) {
	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		log.Println("Error in connecting to database....")
	}

	db.AutoMigrate(&models.Product{})
	log.Println("Database auto-migration finished....")
}
