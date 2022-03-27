package main

import (
	"gorm-db/Config"
	"gorm-db/Routes"
)

func main() {

	// Connect to Db
	Config.Connect()

	// Setup routes
	r := Routes.SetupRouter()

	// Running
	r.Run()

}
