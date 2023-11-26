package PemiluController

import (
	"go_rest_api/models/core"
	"go_rest_api/repositories/vote/Pemilu"
	"log"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// @Summary		Pemilu
// @Description	Ambil data Pemilu
// @Tags			Pemilu
// @Accept			json
// @Produce		json
//
// @Security		ApiKeyAuth
//
// @Router			/api/v1/pemilu [get]
func GetAll(c *fiber.Ctx) error {
	data, totalRows, totalPages, err := Pemilu.GetAll(c)

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

// @Summary		Pemilu By ID
// @Description	Ambil data Pemilu By ID
// @Tags			Pemilu
// @Accept			json
// @Produce		json
// @Param			id	path	int	true	"Pemilu ID"
//
// @Security		ApiKeyAuth
//
// @Router			/api/v1/pemilu/{id} [get]
func GetById(c *fiber.Ctx) error {
	var err error
	return err
}

// @Summary		Pemilu Create
// @Description	Simpan data Pemilu
// @Tags			Pemilu
// @Accept			json
// @Produce		json
// @Param			pemilu	body	Pemilu.PemiluRequest	true	"Pemilu Create"
// @Security		ApiKeyAuth
// @Router			/api/v1/pemilu [post]
func Create(c *fiber.Ctx) error {
	var (
		err   error
		input Pemilu.PemiluRequest
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
	input.CompanyId = user.CompanyId
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

// @Summary		Pemilu Update
// @Description	Simpan data Pemilu
// @Tags			Pemilu
// @Accept			json
// @Produce		json
// @Param			pemilu	body	Pemilu.PemiluUpdateRequest	true	"Pemilu Update"
// @Param			id			path	int								true	"Pemilu ID"
// @Security		ApiKeyAuth
// @Router			/api/v1/pemilu/{id} [post]
func Update(c *fiber.Ctx) error {
	var (
		err   error
		input Pemilu.PemiluUpdateRequest
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

// @Summary		Pemilu Delete
// @Description	Simpan data Pemilu
// @Tags			Pemilu
// @Accept			json
// @Produce		json
// @Param			pemilu	body	Pemilu.PemiluUpdateRequest	true	"Pemilu Update"
// @Param			id			path	int								true	"Pemilu ID"
// @Security		ApiKeyAuth
// @Router			/api/v1/pemilu/delete/{id} [post]
func Delete(c *fiber.Ctx) error {
	var (
		err     error
		payload Pemilu.PemiluRequest
	)
	log.Println(payload)
	return err
}
