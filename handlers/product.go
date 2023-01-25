package handlers

import (
	"logify/database"
	"logify/models"
	"logify/utils"

	"github.com/gofiber/fiber/v2"
	uuid "github.com/satori/go.uuid"
)

func AllProducts(c *fiber.Ctx) error {

	var products []models.Product

	database.Database().Find(&products)

	return c.JSON(products)
}

func GetLatestProducts(c *fiber.Ctx) error {
	var products []models.Product

	database.Database().Order("id desc").Limit(14).Find(&products)

	return c.JSON(products)
}

func GetBrandProducts(c *fiber.Ctx) error {
	var products []models.Product

	brandId, _ := uuid.FromString(c.Query("brandId"))

	database.Database().Where("brand_id = ?", brandId).Find(&products)

	return c.JSON(products)
}

func GetCategoryProducts(c *fiber.Ctx) error {
	var products []models.Product

	categoryId, _ := uuid.FromString(c.Query("categoryId"))

	database.Database().Where("category_id = ?", categoryId).Find(&products)

	return c.JSON(products)
}

func CreateProduct(c *fiber.Ctx) error {
	var req map[string]string

	if err := c.BodyParser(&req); err != nil {
		return err
	}

	cookie := c.Cookies("jwt")

	id, _ := utils.ParseJwt(cookie)

	categoryId := req["category_id"]

	brandId := req["brand_id"]

	userId, _ := uuid.FromString(id)

	categoryid, _ := uuid.FromString(categoryId)

	brandid, _ := uuid.FromString(brandId)

	product := models.Product{
		UserID:        userId,
		CategoryID:    categoryid,
		BrandID:       brandid,
		Name:          req["name"],
		Thumbnail:     req["thumbnail"],
		Details:       req["details"],
		UnitPrice:     req["unit_price"],
		PurchasePrice: req["purchase_price"],
		Tax:           req["tax"],
		TaxType:       req["tax_type"],
		Discount:      req["discount"],
		DiscountType:  req["Discount_type"],
		CurrentStock:  req["current_stock"],
		Status:        req["status"],
	}

	database.Database().Create(&product)

	return c.JSON(product)
}

func GetProduct(c *fiber.Ctx) error {
	id, _ := uuid.FromString(c.Params("id"))

	product := models.Product{
		Id: id,
	}

	database.Database().Find(&product)

	return c.JSON(product)
}

func UpdateProduct(c *fiber.Ctx) error {
	id, _ := uuid.FromString(c.Params("id"))

	var req map[string]string

	if err := c.BodyParser(&req); err != nil {
		return err
	}

	cookie := c.Cookies("jwt")

	userID, _ := utils.ParseJwt(cookie)

	categoryId := req["category_id"]

	brandId := req["brand_id"]

	userId, _ := uuid.FromString(userID)
	categoryid, _ := uuid.FromString(categoryId)
	brandid, _ := uuid.FromString(brandId)

	product := models.Product{
		Id:            id,
		UserID:        userId,
		CategoryID:    categoryid,
		BrandID:       brandid,
		Name:          req["name"],
		Thumbnail:     req["thumbnail"],
		Details:       req["details"],
		UnitPrice:     req["unit_price"],
		PurchasePrice: req["purchase_price"],
		Tax:           req["tax"],
		TaxType:       req["tax_type"],
		Discount:      req["discount"],
		DiscountType:  req["Discount_type"],
		CurrentStock:  req["current_stock"],
		Status:        req["status"],
	}
	database.Database().Model(&product).Updates(product)

	return c.JSON(product)
}

func DeleteProduct(c *fiber.Ctx) error {
	id, _ := uuid.FromString(c.Params("id"))

	product := models.Product{
		Id: id,
	}

	database.Database().Delete(&product)

	return nil
}
