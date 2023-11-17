package main

import (
	"go_rest_api/config"
	"go_rest_api/database"
	"go_rest_api/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	_ "github.com/lib/pq"
)

//	@title			API Documentation
//	@description	API DOCUMENTATION
//	@termsOfService	http://swagger.io/terms/
//	@contact.name	API Support
//	@contact.email	prastiyo.beka12@gmail.com
//	@host			localhost:3000
//	@BasePath		/api/v1

//	@securitydefinitions.apikey	ApiKeyAuth
//	@in							header
//	@name						Authorization

//	@externalDocs.description	OpenAPI
//	@externalDocs.url			https://swagger.io/resources/open-api/
//	@BasePath

func main() {
	database.Connect()
	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New())
	router.SetupRoutes(app)
	// handle unavailable route
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})
	app.Listen(config.Config("APP_PORT"))
}
