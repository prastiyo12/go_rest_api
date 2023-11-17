package IssueController

import (
	"go_rest_api/repositories/vote/Issue"
	"log"

	"github.com/gofiber/fiber/v2"
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
		err     error
		payload Issue.IssueRequest
	)
	log.Println(payload)
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
		err     error
		payload Issue.IssueUpdateRequest
	)
	log.Println(payload)
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
