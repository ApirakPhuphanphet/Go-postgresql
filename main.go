package main

import (
	"github.com/ApirakPhuphanphet/Go-postgresql/db"
	"github.com/ApirakPhuphanphet/Go-postgresql/handler"
	"github.com/ApirakPhuphanphet/Go-postgresql/models"
	"github.com/gofiber/fiber/v2"
)

func router(app *fiber.App) {
	app.Post("/user/signup", handler.Signup)
	app.Post("/user/signin", handler.Signin)
	app.Get("/user/:id", handler.GetUser)
	app.Put("/user/:id", handler.UpdateUser)
	app.Delete("/user/:id", handler.DeleteUser)

	app.Post("/product/create", handler.CreateProduct)
	app.Get("/product", handler.GetAllProduct)
	app.Get("/product/:id", handler.GetProductById)
	app.Put("/product/:id", handler.UpdateProduct)
	app.Delete("/product/:id", handler.DeleteProduct)
}

func main() {
	app := fiber.New()

	DB := db.DatabaseConnection()

	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Product{})

	router(app)

	app.Listen(":8000")
}
