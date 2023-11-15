package router

import (
	"go_rest_api/config"
	"go_rest_api/handlers/core/CompanyController"
	"go_rest_api/handlers/core/LoginController"
	"go_rest_api/handlers/core/ProfileController"
	"go_rest_api/handlers/core/UserController"
	"go_rest_api/handlers/vote/CampaignController"
	"go_rest_api/handlers/vote/DapilController"
	"go_rest_api/handlers/vote/IssueController"
	"go_rest_api/handlers/vote/PemiluController"
	"go_rest_api/handlers/vote/TpsController"

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
	v1.Get("/me", UserController.GetUser)
	v1.Post("/logout", LoginController.LogoutUser)

	//campaign
	campaign := v1.Group("/campaign")
	campaign.Get("/", CampaignController.GetAll)
	campaign.Get("/:id", CampaignController.GetById)
	campaign.Post("/", CampaignController.Create)
	campaign.Post("/:id", CampaignController.Update)
	campaign.Post("/delete/:id", CampaignController.Delete)

	//company
	company := v1.Group("/company")
	company.Get("/", CompanyController.GetAll)
	company.Get("/:id", CompanyController.GetById)
	company.Post("/", CompanyController.Create)
	company.Post("/:id", CompanyController.Update)
	company.Post("/delete/:id", CompanyController.Delete)

	//dapil
	dapil := v1.Group("/dapil")
	dapil.Get("/", DapilController.GetAll)
	dapil.Get("/:id", DapilController.GetById)
	dapil.Post("/", DapilController.Create)
	dapil.Post("/:id", DapilController.Update)
	dapil.Post("/delete/:id", DapilController.Delete)

	//issue
	issue := v1.Group("/issue")
	issue.Get("/", IssueController.GetAll)
	issue.Get("/:id", IssueController.GetById)
	issue.Post("/", IssueController.Create)
	issue.Post("/:id", IssueController.Update)
	issue.Post("/delete/:id", IssueController.Delete)

	//pemilu
	pemilu := v1.Group("/pemilu")
	pemilu.Get("/", PemiluController.GetAll)
	pemilu.Get("/:id", PemiluController.GetById)
	pemilu.Post("/", PemiluController.Create)
	pemilu.Post("/:id", PemiluController.Update)
	pemilu.Post("/delete/:id", PemiluController.Delete)

	//tps
	tps := v1.Group("/tps")
	tps.Get("/", TpsController.GetAll)
	tps.Get("/:id", TpsController.GetById)
	tps.Post("/", TpsController.Create)
	tps.Post("/:id", TpsController.Update)
	tps.Post("/delete/:id", TpsController.Delete)

	//profile
	profile := v1.Group("/profile")
	profile.Get("/", ProfileController.GetAll)
	profile.Get("/:id", ProfileController.GetById)
	profile.Post("/", ProfileController.Create)
	profile.Post("/:id", ProfileController.Update)
	profile.Post("/delete/:id", ProfileController.Delete)

	// user management
	user := v1.Group("/user")
	user.Get("/", UserController.GetAll)
	user.Get("/:id", UserController.GetById)
	user.Post("/", UserController.Create)
	user.Put("/:id", UserController.Update)
	user.Delete("/:id", UserController.Delete)
}
