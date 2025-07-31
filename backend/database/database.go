package database

import (
	"fmt"
	"log"
	"to-do-app/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "host=localhost user=postgres password=postgres port=5432 sslmode=disable Timezone=UTC"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	err = db.AutoMigrate(&model.Todo{})
	if err != nil {
		log.Fatal("Failed to migrate model: ", err)
	}

	DB = db
	fmt.Println("Database connected successfully")
}
