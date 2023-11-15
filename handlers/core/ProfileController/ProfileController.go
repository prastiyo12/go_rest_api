package ProfileController

import (
	"go_rest_api/models/core"
	"go_rest_api/repositories/core/Profile"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// @Summary		Profile
// @Description	Ambil data Profile
// @Tags			Profile
// @Accept			json
// @Produce		json
//
// @Security		ApiKeyAuth
//
// @Router			/api/v1/pemilu [get]
func GetAll(c *fiber.Ctx) error {
	data, err := Profile.GetAll(c)

	if err != nil {
		err := c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": fiber.StatusBadRequest, "message": err.Error()})
		if err != nil {
			return err
		}
		return err
	}

	err = c.Status(fiber.StatusOK).JSON(fiber.Map{"code": fiber.StatusOK, "message": "success", "data": data})
	if err != nil {
		return err
	}
	return err
}

// @Summary		Profile By ID
// @Description	Ambil data Profile By ID
// @Tags			Profile
// @Accept			json
// @Produce		json
// @Param			id	path	int	true	"Profile ID"
//
// @Security		ApiKeyAuth
//
// @Router			/api/v1/profile/{id} [get]
func GetById(c *fiber.Ctx) error {
	var Id = c.Params("id")

	data, err := Profile.GetDataByID(Id)
	if err != nil {
		err := c.Status(fiber.StatusBadRequest).SendString(err.Error())
		if err != nil {
			return nil
		}
	}

	err = c.Status(fiber.StatusOK).JSON(fiber.Map{"code": fiber.StatusOK, "message": "success", "data": data})
	if err != nil {
		return nil
	}
	return err
}

// @Summary		Profile Create
// @Description	Simpan data Profile
// @Tags			Profile
// @Accept			json
// @Produce		json
// @Param			profile	body	Profile.ProfileRequest	true	"Profile Create"
// @Security		ApiKeyAuth
// @Router			/api/v1/profile [post]
func Create(c *fiber.Ctx) error {
	var (
		err   error
		input Profile.ProfileRequest
	)

	user := c.Locals("user").(core.UserResponse)
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
	input.CreatedBy = user.ID

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

// @Summary		Profile Update
// @Description	Simpan data Profile
// @Tags			Profile
// @Accept			json
// @Produce		json
// @Param			profile	body	Profile.ProfileUpdateRequest	true	"Profile Update"
// @Param			id			path	int								true	"Profile ID"
// @Security		ApiKeyAuth
// @Router			/api/v1/profile/{id} [post]
func Update(c *fiber.Ctx) error {
	var (
		err   error
		input Profile.ProfileUpdateRequest
	)
	user := c.Locals("user").(core.UserResponse)
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
	input.UpdatedBy = user.ID

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

// @Summary		Profile Delete
// @Description	Simpan data Profile
// @Tags			Profile
// @Accept			json
// @Produce		json
// @Param			profile	body	Profile.ProfileUpdateRequest	true	"Profile Update"
// @Param			id			path	int								true	"Profile ID"
// @Security		ApiKeyAuth
// @Router			/api/v1/profile/delete/{id} [post]
func Delete(c *fiber.Ctx) error {
	var (
		err error
	)
	var Id = c.Params("id")
	err = Profile.Delete(Id)
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
