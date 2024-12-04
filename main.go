package main

import (
	"fmt"

	"github.com/ApirakPhuphanphet/Go-postgresql/db"
	"github.com/ApirakPhuphanphet/Go-postgresql/models"
)

func main() {
	fmt.Println("Hello, universe!")

	DB := db.DatabaseConnection()

	DB.AutoMigrate(&models.User{})

	// Create
	user := models.User{Username: "Michael", Password: "1111"}
	DB.Create(&user)
}
