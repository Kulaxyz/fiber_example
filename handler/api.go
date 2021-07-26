package handler

import (
	"fmt"
	"github.com/gofiber/fiber/v2"

)

func Hello(c *fiber.Ctx) error {
	name := c.Locals("user_id")
	return c.JSON(fiber.Map{"status": "success", "message": fmt.Sprintf("Hello i'm ok, %v", name) , "data": nil})
}
