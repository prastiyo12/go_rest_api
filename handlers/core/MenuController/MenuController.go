package MenuController

import (
	"go_rest_api/repositories/core/Menu"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// @Summary		Menu
// @Description	Ambil data Menu
// @Tags			Menu
// @Accept			json
// @Produce		json
//
// @Security		ApiKeyAuth
//
// @Router			/api/v1/pemilu [get]
func GetAll(c *fiber.Ctx) error {
	data, err := Menu.GetAll(c)

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

// @Summary		Menu By ID
// @Description	Ambil data Menu By ID
// @Tags			Menu
// @Accept			json
// @Produce		json
// @Param			id	path	int	true	"Menu ID"
//
// @Security		ApiKeyAuth
//
// @Router			/api/v1/menu/{id} [get]
func GetById(c *fiber.Ctx) error {
	var Id = c.Params("id")

	data, err := Menu.GetDataByID(Id)
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

// @Summary		Menu Create
// @Description	Simpan data Menu
// @Tags			Menu
// @Accept			json
// @Produce		json
// @Param			menu	body	Menu.MenuRequest	true	"Menu Create"
// @Security		ApiKeyAuth
// @Router			/api/v1/menu [post]
func Create(c *fiber.Ctx) error {
	var (
		err   error
		input Menu.MenuRequest
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

// @Summary		Menu Update
// @Description	Simpan data Menu
// @Tags			Menu
// @Accept			json
// @Produce		json
// @Param			menu	body	Menu.MenuUpdateRequest	true	"Menu Update"
// @Param			id			path	int								true	"Menu ID"
// @Security		ApiKeyAuth
// @Router			/api/v1/menu/{id} [post]
func Update(c *fiber.Ctx) error {
	var (
		err   error
		input Menu.MenuUpdateRequest
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

// @Summary		Menu Delete
// @Description	Simpan data Menu
// @Tags			Menu
// @Accept			json
// @Produce		json
// @Param			menu	body	Menu.MenuUpdateRequest	true	"Menu Update"
// @Param			id			path	int								true	"Menu ID"
// @Security		ApiKeyAuth
// @Router			/api/v1/menu/delete/{id} [post]
func Delete(c *fiber.Ctx) error {
	var (
		err error
	)
	var Id = c.Params("id")
	err = Menu.Delete(Id)
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

func GetMenu(c *fiber.Ctx) error {
	var err error
	return err
}
