package router

import (
	"go_rest_api/config"
	"go_rest_api/handlers/core/CompanyController"
	"go_rest_api/handlers/core/LoginController"
	"go_rest_api/handlers/core/MenuController"
	"go_rest_api/handlers/core/OptionController"
	"go_rest_api/handlers/core/ProfileController"
	"go_rest_api/handlers/core/RoleController"
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
	v1.Get("/menu", MenuController.GetMenu)
	v1.Post("/logout", LoginController.LogoutUser)

	//campaign
	campaign := v1.Group("/campaign")
	campaign.Get("/", CampaignController.GetAll)
	campaign.Get("/:id", CampaignController.GetById)
	campaign.Post("/", CampaignController.Create)
	campaign.Post("/delete", CampaignController.Delete)
	campaign.Post("/:id", CampaignController.Update)

	//company
	company := v1.Group("/company")
	company.Get("/", CompanyController.GetAll)
	company.Get("/:id", CompanyController.GetById)
	company.Post("/", CompanyController.Create)
	company.Post("/delete", CompanyController.Delete)
	company.Post("/:id", CompanyController.Update)

	//dapil
	dapil := v1.Group("/dapil")
	dapil.Get("/", DapilController.GetAll)
	dapil.Get("/:id", DapilController.GetById)
	dapil.Post("/", DapilController.Create)
	dapil.Post("/delete", DapilController.Delete)
	dapil.Post("/:id", DapilController.Update)

	//issue
	issue := v1.Group("/issue")
	issue.Get("/", IssueController.GetAll)
	issue.Get("/:id", IssueController.GetById)
	issue.Post("/", IssueController.Create)
	issue.Post("/delete", IssueController.Delete)
	issue.Post("/:id", IssueController.Update)

	//pemilu
	pemilu := v1.Group("/pemilu")
	pemilu.Get("/", PemiluController.GetAll)
	pemilu.Get("/:id", PemiluController.GetById)
	pemilu.Post("/", PemiluController.Create)
	pemilu.Post("/delete", PemiluController.Delete)
	pemilu.Post("/:id", PemiluController.Update)

	//tps
	tps := v1.Group("/tps")
	tps.Get("/", TpsController.GetAll)
	tps.Get("/dapil", TpsController.GetAllDapil)
	tps.Get("/dapil-area", TpsController.GetAllDapilArea)
	tps.Get("/tps", TpsController.GetAllTps)
	tps.Get("/:id", TpsController.GetById)
	tps.Post("/", TpsController.Create)
	tps.Post("/delete", TpsController.Delete)
	tps.Post("/:id", TpsController.Update)

	//profile
	profile := v1.Group("/profile")
	profile.Get("/", ProfileController.GetAll)
	profile.Get("/:id", ProfileController.GetById)
	profile.Post("/", ProfileController.Create)
	profile.Post("/delete/:id", ProfileController.Delete)
	profile.Post("/:id", ProfileController.Update)

	// user management
	user := v1.Group("/user")
	user.Get("/", UserController.GetAll)
	user.Get("/:id", UserController.GetById)
	user.Post("/", UserController.Create)
	user.Post("/delete", UserController.Delete)
	user.Post("/upload-image/:id", UserController.UploadImage)
	user.Post("/:id", UserController.Update)

	// user role
	role := v1.Group("/user-role")
	role.Get("/", RoleController.GetAll)
	role.Get("/:id", RoleController.GetById)
	role.Post("/", RoleController.Create)
	role.Post("/delete", RoleController.Delete)
	role.Post("/:id", RoleController.Update)

	//menu
	menu := v1.Group("/menu")
	menu.Get("/all-menu", MenuController.GetAll)
	menu.Get("/:id", MenuController.GetById)
	menu.Post("/", MenuController.Create)
	menu.Post("/delete", MenuController.Delete)
	menu.Post("/:id", MenuController.Update)

	//option
	option := v1.Group("/option")
	option.Get("/province", OptionController.GetProvince)
	option.Get("/city", OptionController.GetCity)
	option.Get("/district", OptionController.GetDistrict)
	option.Get("/village", OptionController.GetVillage)
	option.Get("/user-type", ProfileController.GetAllUserType)
	option.Get("/users", ProfileController.GetAllUser)
}
