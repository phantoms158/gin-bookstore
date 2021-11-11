package databases

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/phantoms158/gin-bookstore/models"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open("sqlite3", "test.db")

	if err != nil {
	  panic("Failed to connect to database!")
	}

	database.AutoMigrate(&models.Book{})

  	DB = database
}