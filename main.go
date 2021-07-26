package main

import (
	"github.com/Kulaxyz/partner/database"
	"github.com/Kulaxyz/partner/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"
)

func main() {
	app := fiber.New()
	database.ConnectDB()

	app.Use(cors.New())
	router.SetupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
