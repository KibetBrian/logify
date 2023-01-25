package middlewares

import (
	"errors"
	"strings"

	"shoppy/database"
	"shoppy/models"
	"shoppy/utils"

	"github.com/gofiber/fiber/v2"
	uuid "github.com/satori/go.uuid"
)

func IsAuthenticated(c *fiber.Ctx) error {
	bearToken := c.Get("Authorization")

	arr := strings.Split(bearToken, " ")

	if _, err := utils.ParseJwt(arr[1]); err != nil {
		c.Status(fiber.StatusUnauthorized)

		return c.JSON(fiber.Map{
			"message": "invalid token",
		})
	}

	return c.Next()
}

func IsAuthorized(c *fiber.Ctx, page string) error {
	cookie := c.Cookies("jwt")

	userId, err := utils.ParseJwt(cookie)
	if err != nil {
		c.JSON(500)
		return c.JSON(fiber.Map{"message": "Internal server error", "success": false})
	}

	id, err := uuid.FromString(userId)
	if err != nil {
		c.JSON(500)
		return c.JSON(fiber.Map{"message": "Internal server error", "success": "false"})
	}

	user := models.User{
		Id: id,
	}

	database.Database().Preload("Role").Find(&user)

	role := models.Role{
		Id: user.RoleId,
	}

	database.Database().Preload("Permissions").Find(&role)
	if c.Method() == "GET" {
		for _, permission := range role.Permissions {
			if permission.Name == "view_"+page || permission.Name == "edit_"+page {
				return nil
			}
		}
	}

	for _, permission := range role.Permissions {
		if permission.Name == "edit_"+page {
			return nil
		}
	}

	c.Status(fiber.StatusUnauthorized)
	
	return errors.New("unauthorized")
}
