package handlers

import (
	"bytes"
	"fmt"
	"html/template"
	"net/smtp"
	"os"
	"logify/database"
	"logify/models"
	"logify/utils"

	"github.com/gofiber/fiber/v2"
)

func Forgot(c *fiber.Ctx) error {
	var req map[string]string

	if err := c.BodyParser(&req); err != nil {
		return err
	}

	randStr, err := utils.GenerateRandomString(16)
	if err != nil {
		c.JSON(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{"message": "Internal server error"})
	}

	passwordReset := models.PasswordReset{
		Email: req["email"],
		Token: randStr,
	}

	database.Database().Create(&passwordReset)

	sender := os.Getenv("EMAIL_SENDER")
	pass := os.Getenv("EMAIL_PASSWORD")

	from := sender
	password := pass

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	to := []string{
		req["email"],
	}

	url := "http://localhost:8080/reset/" + randStr

	auth := smtp.PlainAuth("", from, password, smtpHost)

	t, _ := template.ParseFiles("template.html")

	var body bytes.Buffer

	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf("Subject: Reset Your Password\n%s\n\n", mimeHeaders)))

	t.Execute(&body, struct{ Message string }{Message: url})

	// Sending email.
	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, body.Bytes())
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"message": "Success",
	})
}

func Reset(c *fiber.Ctx) error {
	var req map[string]string

	if err := c.BodyParser(&req); err != nil {
		return err
	}

	token := c.Params("token")

	var tokenModel = models.PasswordReset{}

	var i int64
	database.Database().Where("token = ?", token).Count(&i).First(tokenModel)
	if i == 0 {
		c.JSON(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{"message": "invalid token"})
	}

	if req["email"] != tokenModel.Email {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Email do not match!",
		})
	}

	var passwordReset = models.PasswordReset{}

	if err := database.Database().Where("token = ?", req["token"]).Last(&passwordReset); err.Error != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Invalid token!",
		})
	}

	password, err := utils.HashPassword(req["password"])
	if err != nil {
		c.JSON(500)
		return c.JSON(fiber.Map{"message": "Error occured, try again later"})
	}

	database.Database().Model(&models.User{}).Where("email = ?", passwordReset.Email).Update("password", password)

	return c.JSON(fiber.Map{
		"message": "success",
	})
}
