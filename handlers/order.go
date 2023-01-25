package handlers

import (
	"strconv"
	"strings"

	"logify/database"
	"logify/models"
	"logify/utils"

	"github.com/gofiber/fiber/v2"
	uuid "github.com/satori/go.uuid"
)

func AllOrders(c *fiber.Ctx) error {
	var orders []models.Order

	database.Database().Set("gorm:auto_preload", true).Find(&orders)

	return c.JSON(orders)
}

func GetPendingOrders(c *fiber.Ctx) error {
	var orders []models.Order

	pending := c.Query("pending")

	bearToken := c.Get("Authorization")

	arr := strings.Split(bearToken, " ")

	id, _ := utils.ParseJwt(arr[1])

	userId, _ := uuid.FromString(id)

	database.Database().Where("order_status = ? AND user_id = ?", pending, userId).Find(&orders)

	return c.JSON(orders)
}

func GetProcessingOrders(c *fiber.Ctx) error {
	var orders []models.Order

	processing := c.Query("processing")

	bearToken := c.Get("Authorization")

	arr := strings.Split(bearToken, " ")

	id, _ := utils.ParseJwt(arr[1])

	userId, _ := uuid.FromString(id)

	database.Database().Where("order_status  = ? AND user_id = ?", processing, userId).Find(&orders)

	return c.JSON(orders)
}

func GetTransitOrders(c *fiber.Ctx) error {
	var orders []models.Order

	transit := c.Query("transit")

	bearToken := c.Get("Authorization")

	arr := strings.Split(bearToken, " ")

	id, _ := utils.ParseJwt(arr[1])

	userId, _ := uuid.FromString(id)

	database.Database().Where("order_status  = ? AND user_id = ?", transit, userId).Find(&orders)

	return c.JSON(orders)
}

func GetDeliveredOrders(c *fiber.Ctx) error {
	var orders []models.Order

	delivered := c.Query("delivered")

	bearToken := c.Get("Authorization")

	arr := strings.Split(bearToken, " ")

	id, _ := utils.ParseJwt(arr[1])

	userId, _ := uuid.FromString(id)

	database.Database().Where("order_status  = ? AND user_id = ?", delivered, userId).Find(&orders)

	return c.JSON(orders)
}

func GetCancelOrders(c *fiber.Ctx) error {
	var orders []models.Order

	cancel := c.Query("cancel")

	bearToken := c.Get("Authorization")

	arr := strings.Split(bearToken, " ")

	id, _ := utils.ParseJwt(arr[1])

	userId, _ := uuid.FromString(id)

	database.Database().Where("order_status  = ? AND user_id = ?", cancel, userId).Find(&orders)

	return c.JSON(orders)
}

func UpdateOrderStatus(c *fiber.Ctx) error {
	id, _ := uuid.FromString(c.Params("id"))

	bearToken := c.Get("Authorization")

	arr := strings.Split(bearToken, " ")

	uuidString, _ := utils.ParseJwt(arr[1])

	userId, _ := uuid.FromString(uuidString)

	orders := models.Order{
		Id:     id,
		UserId: userId,
	}

	if err := c.BodyParser(&orders); err != nil {
		return err
	}

	database.Database().Model(&orders).Updates(orders)

	return c.JSON(fiber.Map{
		"message": "Updated successfully",
	})
}

func CreateOrder(c *fiber.Ctx) error {
	var orderDto fiber.Map

	if err := c.BodyParser(&orderDto); err != nil {
		return err
	}

	list := orderDto["order_items"].([]interface{})

	orderItems := make([]models.OrderItem, len(list))

	for i, item := range list {

		productId := item.(map[string]interface{})["product_id"].(string)

		price := item.(map[string]interface{})["price"].(string)

		productName := item.(map[string]interface{})["product_name"].(string)

		image := item.(map[string]interface{})["image"].(string)

		discount := item.(map[string]interface{})["discount"].(string)

		quantity := item.(map[string]interface{})["quantity"].(string)

		taxAmount := item.(map[string]interface{})["tax_amount"].(string)

		newPrice, _ := strconv.ParseFloat(price, 64)
		newQuantity, _ := strconv.ParseFloat(quantity, 64)

		id, _ := uuid.FromString(productId)

		orderItems[i] = models.OrderItem{
			ProductId:   id,
			Price:       newPrice,
			ProductName: productName,
			Image:       image,
			Discount:    discount,
			Quantity:    newQuantity,
			TaxAmount:   taxAmount,
		}
	}

	cookie := c.Cookies("jwt")

	id, _ := utils.ParseJwt(cookie)

	userId, _ := uuid.FromString(id)

	order := models.Order{
		UserId:            userId,
		DeliveryAddressId: orderDto["delivery_address_id"].(string),
		OrderAmount:       orderDto["order_amount"].(string),
		OrderStatus:       orderDto["order_status"].(string),
		PaymentStatus:     orderDto["payment_status"].(string),
		TotalTaxAmount:    orderDto["total_tax_amount"].(string),
		OrderItems:        orderItems,
	}

	database.Database().Create(&order)

	return c.JSON(order)
}
func UpdateOrder(c *fiber.Ctx) error {
	id, _ := uuid.FromString(c.Params("id"))

	var orderDto fiber.Map

	if err := c.BodyParser(&orderDto); err != nil {
		return err
	}

	list := orderDto["order_items"].([]interface{})

	orderItems := make([]models.OrderItem, len(list))

	for i, item := range list {

		productId := item.(map[string]interface{})["product_id"].(string)

		price := item.(map[string]interface{})["price"].(string)

		productName := item.(map[string]interface{})["product_name"].(string)

		image := item.(map[string]interface{})["image"].(string)

		discount := item.(map[string]interface{})["discount"].(string)

		quantity := item.(map[string]interface{})["quantity"].(string)
		
		taxAmount := item.(map[string]interface{})["tax_amount"].(string)

		newPrice, _ := strconv.ParseFloat(price, 64)
		newQuantity, _ := strconv.ParseFloat(quantity, 64)

		id, _ := uuid.FromString(productId)

		orderItems[i] = models.OrderItem{
			ProductId:   id,
			Price:       newPrice,
			ProductName: productName,
			Image:       image,
			Discount:    discount,
			Quantity:    newQuantity,
			TaxAmount:   taxAmount,
		}
	}

	cookie := c.Cookies("jwt")

	userId, _ := utils.ParseJwt(cookie)

	uId, _ := uuid.FromString(userId)

	order := models.Order{
		Id:                id,
		UserId:            uId,
		DeliveryAddressId: orderDto["delivery_address_id"].(string),
		OrderAmount:       orderDto["order_amount"].(string),
		OrderStatus:       orderDto["order_status"].(string),
		PaymentStatus:     orderDto["payment_status"].(string),
		TotalTaxAmount:    orderDto["total_tax_amount"].(string),
		OrderItems:        orderItems,
	}

	database.Database().Updates(&order)

	return c.JSON(order)
}

func DeleteOrder(c *fiber.Ctx) error {
	id, _ := uuid.FromString(c.Params("id"))

	order := models.Order{
		Id: id,
	}
	database.Database().Delete(&order)

	return c.JSON(fiber.Map{
		"message": "Order Deleted successfully",
	})

}

func ClearOrder(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	id, _ := utils.ParseJwt(cookie)

	userId, _ := uuid.FromString(id)

	order := models.Order{
		UserId: userId,
	}

	database.Database().Where("user_id = ?", userId).Unscoped().Delete(&order)

	return c.JSON(order)
}
