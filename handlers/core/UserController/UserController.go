package UserController

import (
	"go_rest_api/models/core"
	"go_rest_api/repositories/core/User"
	"log"

	"github.com/gofiber/fiber/v2"
)

// @Summary		Get User
// @Description	Ambil data profile user login
// @Tags			User Credential
// @Accept			json
// @Produce		json
//
//	@Security		ApiKeyAuth
//
// @Router			/api/v1/me [get]
func GetUser(c *fiber.Ctx) error {
	user := c.Locals("user").(core.UserResponse)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"user": user}})
}

//	@Summary		User
//	@Description	Ambil data User
//	@Tags			User
//	@Accept			json
//	@Produce		json
//
//	@Security		ApiKeyAuth
//
//	@Router			/api/v1/user [get]

func GetAll(c *fiber.Ctx) error {
	var err error
	return err
}

// @Summary		User By ID
// @Description	Ambil data User By ID
// @Tags			User
// @Accept			json
// @Produce		json
// @Param			id	path	int	true	"User ID"
//
// @Security		ApiKeyAuth
//
// @Router			/api/v1/user/{id} [get]
func GetById(c *fiber.Ctx) error {
	var err error
	return err
}

// @Summary		User Create
// @Description	Simpan data User
// @Tags			User
// @Accept			json
// @Produce		json
// @Param			user	body	User.UserRequest	true	"User Create"
// @Security		ApiKeyAuth
// @Router			/api/v1/user [post]
func Create(c *fiber.Ctx) error {
	var (
		err     error
		payload User.UserRequest
	)
	log.Println(payload)
	return err
}

// @Summary		User Update
// @Description	Simpan data User
// @Tags			User
// @Accept			json
// @Produce		json
// @Param			user	body	User.UserUpdateRequest	true	"User Update"
// @Param			id			path	int								true	"User ID"
// @Security		ApiKeyAuth
// @Router			/api/v1/user/{id} [post]
func Update(c *fiber.Ctx) error {
	var (
		err     error
		payload User.UserUpdateRequest
	)
	log.Println(payload)
	return err
}

// @Summary		User Delete
// @Description	Simpan data User
// @Tags			User
// @Accept			json
// @Produce		json
// @Param			user	body	User.UserUpdateRequest	true	"User Update"
// @Param			id			path	int								true	"User ID"
// @Security		ApiKeyAuth
// @Router			/api/v1/user/delete/{id} [post]
func Delete(c *fiber.Ctx) error {
	var (
		err     error
		payload User.UserRequest
	)
	log.Println(payload)
	return err
}
