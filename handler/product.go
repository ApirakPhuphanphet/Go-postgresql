package handler

import (
	"strconv"
	"time"

	"github.com/ApirakPhuphanphet/Go-postgresql/models"
	"github.com/gofiber/fiber/v2"
)

func CreateProduct(c *fiber.Ctx) error {
	now := time.Now()

	var newProduct models.Product
	var existProduct models.Product

	if err := c.BodyParser(&newProduct); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err,
		})
	}
	DB.Where("name = ?", newProduct.Name).First(&existProduct)

	if existProduct.Name != "" {
		return c.Status(400).JSON(fiber.Map{"error": `Product already exists`})
	}

	newProduct.CreateAt = string(now.Format("2006-01-02 15:04:05"))
	newProduct.UpdateAt = string(now.Format("2006-01-02 15:04:05"))

	result := DB.Create(&newProduct)

	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": result.Error,
		})
	}

	return c.Status(201).JSON(newProduct)
}

func GetAllProduct(c *fiber.Ctx) error {
	var products []models.Product
	DB.Find(&products)

	return c.Status(200).JSON(products)
}

func GetProductById(c *fiber.Ctx) error {
	id := c.Params("id")
	var product models.Product

	DB.Where("id = ?", id).First(&product)

	if product.Name == "" {
		return c.Status(404).JSON(fiber.Map{
			"error": "Product not found",
		})
	}

	return c.Status(200).JSON(product)
}

func UpdateProduct(c *fiber.Ctx) error {
	// convet id to int
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err,
		})
	}
	var existProduct models.Product
	var updateProduct models.Product
	now := time.Now()

	if err := c.BodyParser(&updateProduct); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err,
		})
	}

	DB.Where("name = ?", updateProduct.Name).First(&existProduct)

	// Check if product name has been used
	if existProduct.Name != "" && uint64(existProduct.ID) != id {
		return c.Status(400).JSON(fiber.Map{"error": `Product name has been used`})
	}

	DB.Where("id = ?", id).First(&existProduct)

	if existProduct.Name == "" {
		return c.Status(404).JSON(fiber.Map{
			"error": "Product not found",
		})
	}

	updateProduct.UpdateAt = string(now.Format("2006-01-02 15:04:05"))
	DB.Model(&existProduct).Updates(updateProduct)

	return c.Status(200).JSON(existProduct)
}

func DeleteProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	var product models.Product

	DB.Where("id = ?", id).First(&product)

	if product.Name == "" {
		return c.Status(404).JSON(fiber.Map{
			"error": "Product not found",
		})
	}

	DB.Delete(&product)

	return c.Status(200).JSON(fiber.Map{
		"message": "Product deleted successfully",
	})
}
