package PemilihController

import (
	"go_rest_api/models/core"
	"go_rest_api/repositories/vote/Pemilih"
	"log"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// @Summary		Pemilih
// @Description	Ambil data Pemilih
// @Tags			Pemilih
// @Accept			json
// @Produce		json
//
// @Security		ApiKeyAuth
//
// @Router			/api/v1/pemilih [get]
func GetAll(c *fiber.Ctx) error {
	data, totalRows, totalPages, err := Pemilih.GetAll(c)

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

// @Summary		Pemilih By ID
// @Description	Ambil data Pemilih By ID
// @Tags			Pemilih
// @Accept			json
// @Produce		json
// @Param			id	path	int	true	"Pemilih ID"
//
// @Security		ApiKeyAuth
//
// @Router			/api/v1/pemilih/{id} [get]
func GetById(c *fiber.Ctx) error {
	var err error
	return err
}

// @Summary		Pemilih Create
// @Description	Simpan data Pemilih
// @Tags			Pemilih
// @Accept			json
// @Produce		json
// @Param			issue	body	Pemilih.IssueRequest	true	"Pemilih Create"
// @Security		ApiKeyAuth
// @Router			/api/v1/pemilih [post]
func Create(c *fiber.Ctx) error {
	var (
		err   error
		input Pemilih.PemilihRequest
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
	input.UserId = user.ID
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

// @Summary		Pemilih Update
// @Description	Simpan data Pemilih
// @Tags			Pemilih
// @Accept			json
// @Produce		json
// @Param			issue	body	Pemilih.PemilihUpdateRequest	true	"Issue Update"
// @Param			id			path	int								true	"Issue ID"
// @Security		ApiKeyAuth
// @Router			/api/v1/pemilih/{id} [post]
func Update(c *fiber.Ctx) error {
	var (
		err   error
		input Pemilih.PemilihUpdateRequest
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

// @Summary		Pemilih Delete
// @Description	Simpan data Pemilih
// @Tags			Pemilih
// @Accept			json
// @Produce		json
// @Param			issue	body	Pemilih.PemilihUpdateRequest	true	"Pemilih Update"
// @Param			id			path	int								true	"Pemilih ID"
// @Security		ApiKeyAuth
// @Router			/api/v1/pemilih/delete/{id} [post]
func Delete(c *fiber.Ctx) error {
	var (
		err     error
		payload Pemilih.PemilihRequest
	)
	log.Println(payload)
	return err
}
