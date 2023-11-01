package router

import (
	"go_rest_api/config"
	"go_rest_api/handlers/LoginController"
	"go_rest_api/handlers/UserController.go"

	"go_rest_api/middleware"

	_ "go_rest_api/docs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func SetupRoutes(app *fiber.App) {
	// grouping
	app.Get("/docs/*", swagger.HandlerDefault)
	app.Get("/", func(c *fiber.Ctx) error {
		message := config.Config("APP_NAME")
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"Welcome": message})
	})
	api := app.Group("/api")
	auth := api.Group("/auth")
	auth.Post("/login", LoginController.SignInUser)
	auth.Post("/sign-up", LoginController.SignUpUser)
	// app.Post("/forgot-password", LoginController.Login)

	// api with middleware
	v1 := api.Group("/v1")
	jwt := middleware.DeserializeUser
	v1.Use(jwt)

	//login system
	v1.Get("/me", UserController.GetProfile)
	v1.Post("/logout", LoginController.LogoutUser)

	// user management
	// apiWithMiddlware.Get("/", UserController.GetAllUsers)
	// apiWithMiddlware.Get("/:id", UserController.GetSingleUser)
	// apiWithMiddlware.Post("/", UserController.CreateUser)
	// apiWithMiddlware.Put("/:id", UserController.UpdateUser)
	// apiWithMiddlware.Delete("/:id", UserController.DeleteUserByID)
}
