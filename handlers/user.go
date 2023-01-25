package handlers

import (
	"shoppy/database"
	"shoppy/models"
	"shoppy/utils"
	"strings"

	"github.com/gofiber/fiber/v2"
	uuid "github.com/satori/go.uuid"
)

func UpdateInfo(c *fiber.Ctx) error {
	var req map[string]string

	if err := c.BodyParser(&req); err != nil {
		return err
	}

	bearToken := c.Get("Authorization")

	arr := strings.Split(bearToken, " ")

	userId, err := utils.ParseJwt(arr[1])
	if err != nil {
		c.JSON(409)
		return c.JSON(fiber.Map{"success": false, "message": "Invalid token"})
	}

	id, err := uuid.FromString(userId)
	if err != nil {
		c.JSON(500)
		return c.JSON(fiber.Map{"message": "Internal server error", "success": "false"})
	}

	user := models.User{
		Id:        id,
		FirstName: req["f_name"],
		LastName:  req["l_name"],
		Email:     req["email"],
		Phone:     req["phone"],
	}

	user.SetPassword(req["password"])
	
	database.Database().Model(&user).Updates(user)

	return c.JSON(fiber.Map{
		"message": "user updated successfully",
	})
}
