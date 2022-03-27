package Config

import (
	"fmt"
	"gorm-db/Models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {

	dsn := "host=localhost user=root password=myrefuge dbname=myrefugeDb port=5432 sslmode=disable TimeZone=america/new_york"

	// Creating a connection to the database
	// Config.DB, err = gorm.Open("mysql", Config.DBUrl((Config.BuildDBConfig())))
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("status : ", err)
	}

	// defer Config.DB.Close()
	// defer DB.Close()

	// run the migrations: todo struct
	// Config.DB.AutoMigrate(&Models.BaselineException{})
	database.AutoMigrate(&Models.Verse{})
	database.AutoMigrate(&Models.Category{})

	// database.Create(&Models.Category{Name: "Pride"})
	// database.Create(&Models.Verse{Book: "John", Chapter: 3, Verse: 16, Text: "For God so loved the world that He gave His Only Begotten Son...", Translation_Name: "NKJV", CategoryID: 1})

	// Populate the the DB variable with our database instance. We will use this variable in our handlers to get access to the database.
	DB = database
}
