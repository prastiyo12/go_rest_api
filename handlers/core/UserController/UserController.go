package UserController

import (
	"go_rest_api/models/core"
	"go_rest_api/repositories/core/User"
	"log"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
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
	data, totalRows, totalPages, err := User.GetAll(c)

	if err != nil {
		err := c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": fiber.StatusBadRequest, "message": err.Error()})
		if err != nil {
			return err
		}
		return err
	}

	err = c.Status(fiber.StatusOK).JSON(fiber.Map{"code": fiber.StatusOK, "message": "Data Stored.", "data": data, "totalRow": totalRows, "totalPages": totalPages})
	if err != nil {
		return err
	}
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
		err   error
		input User.UserRequest
	)

	if err := c.BodyParser(&input); err != nil {
		err := c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": fiber.StatusBadRequest, "message": err.Error()})
		if err != nil {
			return nil
		}
	}

	validate := validator.New()
	errValidate := validate.Struct(input)
	if errValidate != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "failed",
			"error":   errValidate.Error(),
		})
	}

	input.ID = uuid.New()
	input.Status = true
	input.CreatedAt = time.Now()

	_, err = input.Create()
	if err != nil {
		err := c.Status(fiber.StatusBadRequest).SendString(err.Error())
		if err != nil {
			return nil
		}
	}

	err = c.Status(fiber.StatusOK).JSON(fiber.Map{"code": fiber.StatusOK, "message": "success"})
	if err != nil {
		return nil
	}
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
		err   error
		input User.UserUpdateRequest
	)

	if err := c.BodyParser(&input); err != nil {
		err := c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": fiber.StatusBadRequest, "message": err.Error()})
		if err != nil {
			return nil
		}
	}

	validate := validator.New()
	errValidate := validate.Struct(input)
	if errValidate != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "failed",
			"error":   errValidate.Error(),
		})
	}

	var Id = c.Params("id")
	input.UpdatedAt = time.Now()

	_, err = input.Update(Id)
	if err != nil {
		err := c.Status(fiber.StatusBadRequest).SendString(err.Error())
		if err != nil {
			return nil
		}
	}

	err = c.Status(fiber.StatusOK).JSON(fiber.Map{"code": fiber.StatusOK, "message": "success"})
	if err != nil {
		return nil
	}
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
