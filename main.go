package main

import (
	"fmt"
	"log"

	"github.com/rizqirenaldy27/invoice-esb/infrastructure/config"
	"github.com/rizqirenaldy27/invoice-esb/infrastructure/database"
	"github.com/rizqirenaldy27/invoice-esb/interface/http"
)

func main() {

	fmt.Println("Starting API")
	db, err := database.InitDB()
	if err != nil {
		log.Fatal(err)
		return
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
		return
	}

	defer sqlDB.Close()

	port := config.AppConfig.SystemConfig.Port

	fmt.Printf("Start API in port %s", port)

	app := http.InitRouter(db)

	app.Listen(":" + port)
}
