package handler

import (
	"fmt"
	"time"

	"github.com/ApirakPhuphanphet/Go-postgresql/db"
	"github.com/ApirakPhuphanphet/Go-postgresql/models"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

var DB = db.DatabaseConnection()

func Signup(c *fiber.Ctx) error {
	newUser := new(models.User)
	start := time.Now()

	if err := c.BodyParser(newUser); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	bytes, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), 14)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	newUser.Password = string(bytes)
	// Check if user already exist
	var existingUser models.User

	DB.Where("username = ?", newUser.Username).First(&existingUser)

	if existingUser.Username != "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "user already exist"})
	}

	result := DB.Create(&newUser)

	fmt.Print("Time: " + time.Since(start).String())
	return c.Status(fiber.StatusCreated).JSON(result)
}

func Signin(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	var dbUser models.User
	DB.Where("username = ?", user.Username).First(&dbUser)
	err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password))
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "incorrect password"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Sign in success"})
}

func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User
	DB.First(&user, id)
	if user.Username == "" {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "user not found"})
	}
	DB.Delete(&user)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Delete success"})
}

func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User
	DB.First(&user, id)
	if user.Username == "" {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "user not found"})
	}
	user.Password = ""
	return c.Status(fiber.StatusOK).JSON(user)
}

func UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var existUser models.User
	var UpdateUser models.User

	DB.First(&existUser, id)
	// Check if user already exist
	if existUser.Username == "" {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "user not found"})
	}
	if err := c.BodyParser(&UpdateUser); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	// Check if password match
	if err := bcrypt.CompareHashAndPassword([]byte(existUser.Password), []byte(UpdateUser.Password)); err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"error": "password not match"})
	}

	UpdateUser.Password = string([]byte(existUser.Password))
	DB.Model(&existUser).Updates(UpdateUser)

	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "update data"})
}
