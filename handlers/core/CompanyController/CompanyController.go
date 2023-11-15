package CompanyController

import (
	"go_rest_api/models/core"
	"go_rest_api/repositories/core/Company"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// @Summary		Company
// @Description	Ambil data Company
// @Tags			Company
// @Accept			json
// @Produce		json
// @Param           search    query   string  false  "name search"
// @Security		ApiKeyAuth
//
// @Router			/api/v1/company [get]
func GetAll(c *fiber.Ctx) error {
	data, err := Company.GetAll(c)

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

// @Summary		Company By ID
// @Description	Ambil data Company By ID
// @Tags			Company
// @Accept			json
// @Produce		json
// @Param			id	path	int	true	"Company ID"
//
// @Security		ApiKeyAuth
//
// @Router			/api/v1/company/{id} [get]
func GetById(c *fiber.Ctx) error {
	var Id = c.Params("id")

	data, err := Company.GetDataByID(Id)
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

// @Summary		Company Create
// @Description	Simpan data Company
// @Tags			Company
// @Accept			json
// @Produce		json
// @Param			company	body	Company.CompanyRequest	true	"Company Create"
// @Security		ApiKeyAuth
// @Router			/api/v1/company [post]
func Create(c *fiber.Ctx) error {
	var (
		err   error
		input Company.CompanyRequest
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

// @Summary		Company Update
// @Description	Simpan data Company
// @Tags			Company
// @Accept			json
// @Produce		json
// @Param			company	body	Company.CompanyUpdateRequest	true	"Company Update"
// @Param			id			path	int								true	"Company ID"
// @Security		ApiKeyAuth
// @Router			/api/v1/company/{id} [post]
func Update(c *fiber.Ctx) error {
	var (
		err   error
		input Company.CompanyUpdateRequest
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

// @Summary		Company Delete
// @Description	Simpan data Company
// @Tags			Company
// @Accept			json
// @Produce		json
// @Param			company	body	Company.CompanyUpdateRequest	true	"Company Update"
// @Param			id			path	int								true	"Company ID"
// @Security		ApiKeyAuth
// @Router			/api/v1/company/delete/{id} [post]
func Delete(c *fiber.Ctx) error {
	var (
		err error
	)
	var Id = c.Params("id")
	err = Company.Delete(Id)
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
