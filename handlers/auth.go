package handlers

import (
	"shoppy/database"
	"shoppy/models"
	"shoppy/utils"
	"time"

	"github.com/gofiber/fiber/v2"
	uuid "github.com/satori/go.uuid"
	"github.com/spf13/viper"
)

func Register(c fiber.Ctx) error {
	var req map[string]string

	if err := c.BodyParser(&req); err != nil {
		return err
	}

	id := uuid.NewV4()

	user := models.User{
		FirstName: req["f_name"],
		LastName:  req["l_name"],
		Email:     req["email"],
		Phone:     req["phone"],
		Id:        id,
	}

	user.SetPassword(req["password"])

	var count int64
	database.Database().Model(&models.User{}).Where("email = ?", req["email"]).Count(&count)

	if count != 0 {
		c.JSON(409)
		return c.JSON(fiber.Map{"success": false, "message": "email already exists"})
	}

	database.Database().Create(&user)

	uuidString := user.Id.String()

	token, err := utils.GenerateJwt(uuidString)
	if err != nil {
		c.JSON(500)
		return c.JSON(fiber.Map{"success": false, "message": "Internal server error"})
	}

	cookieExpiry := viper.Get("COOKIE_DURATION").(int)

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Duration(time.Duration(cookieExpiry).Hours())),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "Registered successfully",
		"token":   token,
		"user":    user,
	})
}

func Login(c *fiber.Ctx) error {
	var req map[string]string

	if err := c.BodyParser(&req); err != nil {
		return err
	}

	var user models.User

	var count int64
	database.Database().Model(&models.User{}).Where("email = ?", req["email"]).Count(&count)

	if count == 0 {
		c.JSON(404)
		return c.JSON(fiber.Map{"success": false, "message": "user not found"})
	}

	if err := user.ComparePassword(req["password"]); err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"success": false,
			"message": "email and password do not match",
		})
	}

	token, err := utils.GenerateJwt(user.Id.String())
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	cookieExpiry := viper.Get("COOKIE_DURATION").(int)

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Duration(time.Duration(cookieExpiry).Hours())),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "login successfully",
		"token":   token,
		"user":    user,
	})
}

func Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "logout success",
	})
}