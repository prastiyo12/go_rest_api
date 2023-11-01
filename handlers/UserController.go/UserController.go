package UserController

import (
	"go_rest_api/models/entity"

	"github.com/gofiber/fiber/v2"
)

// @Summary		Get Profile
// @Description	Ambil data profile user login
// @Tags			User Credential
// @Accept			json
// @Produce		json
//
//	@Security		ApiKeyAuth
//
// @Router			/api/v1/me [get]
func GetProfile(c *fiber.Ctx) error {
	user := c.Locals("user").(entity.UserResponse)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"user": user}})
}
