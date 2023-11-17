package TpsController

import (
	"go_rest_api/models/core"
	"go_rest_api/repositories/vote/Tps"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// @Summary		Tps
// @Description	Ambil data Tps
// @Tags			Tps
// @Accept			json
// @Produce		json
//
// @Security		ApiKeyAuth
//
// @Router			/api/v1/tps [get]
func GetAll(c *fiber.Ctx) error {
	data, totalRows, totalPages, err := Tps.GetAll(c)

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

// @Summary		Tps By ID
// @Description	Ambil data Tps By ID
// @Tags			Tps
// @Accept			json
// @Produce		json
// @Param			id	path	int	true	"Tps ID"
//
// @Security		ApiKeyAuth
//
// @Router			/api/v1/tps/{id} [get]
func GetById(c *fiber.Ctx) error {
	var Id = c.Params("id")

	data, err := Tps.GetDataByID(Id)
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

// @Summary		Tps Create
// @Description	Simpan data Tps
// @Tags			Tps
// @Accept			json
// @Produce		json
// @Param			tps	body	Tps.TpsRequest	true	"Tps Create"
// @Security		ApiKeyAuth
// @Router			/api/v1/tps [post]
func Create(c *fiber.Ctx) error {
	var (
		err   error
		input Tps.TpsRequest
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

// @Summary		Tps Update
// @Description	Simpan data Tps
// @Tags			Tps
// @Accept			json
// @Produce		json
// @Param			tps	body	Tps.TpsUpdateRequest	true	"Tps Update"
// @Param			id			path	int								true	"Tps ID"
// @Security		ApiKeyAuth
// @Router			/api/v1/tps/{id} [post]
func Update(c *fiber.Ctx) error {
	var (
		err   error
		input Tps.TpsUpdateRequest
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

// @Summary		Tps Delete
// @Description	Simpan data Tps
// @Tags			Tps
// @Accept			json
// @Produce		json
// @Param			tps	body	Tps.TpsUpdateRequest	true	"Tps Update"
// @Param			id			path	int								true	"Tps ID"
// @Security		ApiKeyAuth
// @Router			/api/v1/tps/delete/{id} [post]
func Delete(c *fiber.Ctx) error {
	var (
		err error
	)
	var Id = c.Params("id")
	err = Tps.Delete(Id)
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

func GetAllDapil(c *fiber.Ctx) error {
	data, err := Tps.GetAllDapil(c)
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

func GetAllDapilArea(c *fiber.Ctx) error {
	data, err := Tps.GetAllDapilArea(c)
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
