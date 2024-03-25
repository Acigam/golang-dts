package main

import (
	"final-project-acgm/database"
	"final-project-acgm/routers"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file: ", err)
	}

	port := os.Getenv("PORT")
	database.StartDB()
	r := routers.StartApp()
	r.Run(fmt.Sprintf(":%s", port))
}
