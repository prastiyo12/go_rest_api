package RoleController

import (
	"go_rest_api/repositories/core/Role"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// @Summary		Role
// @Description	Ambil data Role
// @Tags			Role
// @Accept			json
// @Produce		json
//
// @Security		ApiKeyAuth
//
// @Router			/api/v1/user-role [get]
func GetAll(c *fiber.Ctx) error {
	data, err := Role.GetAll(c)

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

// @Summary		Role By ID
// @Description	Ambil data Role By ID
// @Tags			Role
// @Accept			json
// @Produce		json
// @Param			id	path	int	true	"Role ID"
//
// @Security		ApiKeyAuth
//
// @Router			/api/v1/user-role/{id} [get]
func GetById(c *fiber.Ctx) error {
	var Id = c.Params("id")

	data, err := Role.GetDataByID(Id)
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

// @Summary		Role Create
// @Description	Simpan data Role
// @Tags			Role
// @Accept			json
// @Produce		json
// @Param			user_role	body	Role.RoleInput	true	"Role Create"
// @Security		ApiKeyAuth
// @Router			/api/v1/user-role [post]
func Create(c *fiber.Ctx) error {
	var (
		err   error
		input Role.RoleInput
		model Role.RoleRequest
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

	model.ID = uuid.New()
	model.CompanyId = input.CompanyId
	model.Name = input.Name
	model.Status = true
	model.CreatedAt = time.Now()

	_, err = model.Create()
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

// @Summary		Role Update
// @Description	Simpan data Role
// @Tags			Role
// @Accept			json
// @Produce		json
// @Param			user_roles	body	Role.RoleInputUpdate	true	"Role Update"
// @Param			id			path	string								true	"Role ID"
// @Security		ApiKeyAuth
// @Router			/api/v1/user-role/{id} [post]
func Update(c *fiber.Ctx) error {
	var (
		err   error
		input Role.RoleInputUpdate
		model Role.RoleUpdateRequest
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
	model.Name = input.Name
	model.Status = input.Status
	model.UpdatedAt = time.Now()

	_, err = model.Update(Id)
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

// @Summary		Role Delete
// @Description	Simpan data Role
// @Tags			Role
// @Accept			json
// @Produce		json
// @Param			user_roles	body	Role.RoleUpdateRequest	true	"Role Update"
// @Param			id			path	string								true	"Role ID"
// @Security		ApiKeyAuth
// @Router			/api/v1/user-role/delete/{id} [post]
func Delete(c *fiber.Ctx) error {
	var (
		err error
	)
	var Id = c.Params("id")
	err = Role.Delete(Id)
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
