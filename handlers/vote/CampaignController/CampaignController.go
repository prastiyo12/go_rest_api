package CampaignController

import (
	"go_rest_api/models/core"
	"go_rest_api/repositories/vote/Campaign"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// @Summary		Campaign
// @Description	Ambil data Campaign
// @Tags			Campaign
// @Accept			json
// @Produce		json
//
// @Security		ApiKeyAuth
//
// @Router			/api/v1/campaign [get]
func GetAll(c *fiber.Ctx) error {
	data, totalRows, totalPages, err := Campaign.GetAll(c)

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

// @Summary		Campaign By ID
// @Description	Ambil data Campaign By ID
// @Tags			Campaign
// @Accept			json
// @Produce		json
// @Param			id	path	int	true	"Campaign ID"
//
// @Security		ApiKeyAuth
//
// @Router			/api/v1/campaign/{id} [get]
func GetById(c *fiber.Ctx) error {
	var Id = c.Params("id")

	data, err := Campaign.GetDataByID(Id)
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

// @Summary		Campaign Create
// @Description	Simpan data Campaign
// @Tags			Campaign
// @Accept			json
// @Produce		json
// @Param			campaign	body	Campaign.CampaignRequest	true	"Campaign Create"
// @Security		ApiKeyAuth
// @Router			/api/v1/campaign [post]
func Create(c *fiber.Ctx) error {
	var (
		err   error
		input Campaign.CampaignRequest
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

// @Summary		Campaign Update
// @Description	Simpan data Campaign
// @Tags			Campaign
// @Accept			json
// @Produce		json
// @Param			campaign	body	Campaign.CampaignUpdateRequest	true	"Campaign Update"
// @Param			id			path	int								true	"Campaign ID"
// @Security		ApiKeyAuth
// @Router			/api/v1/campaign/{id} [post]
func Update(c *fiber.Ctx) error {
	var (
		err   error
		input Campaign.CampaignUpdateRequest
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

// @Summary		Campaign Delete
// @Description	Simpan data Campaign
// @Tags			Campaign
// @Accept			json
// @Produce		json
// @Param			campaign	body	Campaign.CampaignUpdateRequest	true	"Campaign Update"
// @Param			id			path	int								true	"Campaign ID"
// @Security		ApiKeyAuth
// @Router			/api/v1/campaign/delete/{id} [post]
func Delete(c *fiber.Ctx) error {
	var (
		err error
	)
	var Id = c.Params("id")
	err = Campaign.Delete(Id)
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
