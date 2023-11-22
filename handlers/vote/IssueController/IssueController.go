package IssueController

import (
	"go_rest_api/models/core"
	"go_rest_api/repositories/vote/Issue"
	"log"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// @Summary		Issue
// @Description	Ambil data Issue
// @Tags			Issue
// @Accept			json
// @Produce		json
//
// @Security		ApiKeyAuth
//
// @Router			/api/v1/issue [get]
func GetAll(c *fiber.Ctx) error {
	data, totalRows, totalPages, err := Issue.GetAll(c)

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

// @Summary		Issue By ID
// @Description	Ambil data Issue By ID
// @Tags			Issue
// @Accept			json
// @Produce		json
// @Param			id	path	int	true	"Issue ID"
//
// @Security		ApiKeyAuth
//
// @Router			/api/v1/issue/{id} [get]
func GetById(c *fiber.Ctx) error {
	var err error
	return err
}

// @Summary		Issue Create
// @Description	Simpan data Issue
// @Tags			Issue
// @Accept			json
// @Produce		json
// @Param			issue	body	Issue.IssueRequest	true	"Issue Create"
// @Security		ApiKeyAuth
// @Router			/api/v1/issue [post]
func Create(c *fiber.Ctx) error {
	var (
		err   error
		input Issue.IssueRequest
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

// @Summary		Issue Update
// @Description	Simpan data Issue
// @Tags			Issue
// @Accept			json
// @Produce		json
// @Param			issue	body	Issue.IssueUpdateRequest	true	"Issue Update"
// @Param			id			path	int								true	"Issue ID"
// @Security		ApiKeyAuth
// @Router			/api/v1/issue/{id} [post]
func Update(c *fiber.Ctx) error {
	var (
		err   error
		input Issue.IssueUpdateRequest
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

// @Summary		Issue Delete
// @Description	Simpan data Issue
// @Tags			Issue
// @Accept			json
// @Produce		json
// @Param			issue	body	Issue.IssueUpdateRequest	true	"Issue Update"
// @Param			id			path	int								true	"Issue ID"
// @Security		ApiKeyAuth
// @Router			/api/v1/issue/delete/{id} [post]
func Delete(c *fiber.Ctx) error {
	var (
		err     error
		payload Issue.IssueRequest
	)
	log.Println(payload)
	return err
}
