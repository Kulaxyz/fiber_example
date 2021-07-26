package handler

import (
	_ "errors"
	"github.com/Kulaxyz/partner/repository"
	"github.com/gofiber/fiber/v2"
	_ "gorm.io/gorm"
	"net/mail"

	helper "github.com/Kulaxyz/partner/helpers"
	model "github.com/Kulaxyz/partner/models"
)

type LoginInput struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

func Login(c *fiber.Ctx) error {
	var input LoginInput

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Error on login request parse", "data": err})
	}
	username := input.Email
	pass := input.Password

	var user *model.User
	var noUserErr error
	if _, invalidEmail := mail.ParseAddress(username); invalidEmail == nil {
		user, noUserErr = repository.GetUserByEmail(username)
	} else {
		user, noUserErr = repository.GetUserByUsername(username)
	}

	if noUserErr != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "User not found", "data": noUserErr})
	}

	if !helper.IsSame(pass, user.Password) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "Invalid password", "data": nil})
	}

	t, err := helper.NewToken(user)

	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Success login", "data": t})
}

