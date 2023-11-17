package DapilController

import (
	"go_rest_api/models/core"
	"go_rest_api/repositories/vote/Dapil"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// @Summary		Dapil
// @Description	Ambil data Dapil
// @Tags			Dapil
// @Accept			json
// @Produce		json
// @Param			page	query	string	true	"Page"
// @Param			rows	query	string	true	"Rows"
// @Param			dir	query	string	true	"Dir"
// @Param			sort	query	string	true	"Sort"
// @Param			code	query	string	false	"Code"
// @Param			name	query	string	false	"Name"
// @Security		ApiKeyAuth
//
// @Router			/api/v1/dapil [get]
func GetAll(c *fiber.Ctx) error {
	data, totalRows, totalPages, err := Dapil.GetAll(c)

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

// @Summary		Dapil By ID
// @Description	Ambil data Dapil By ID
// @Tags			Dapil
// @Accept			json
// @Produce		json
// @Param			id	path	int	true	"Dapil ID"
//
// @Security		ApiKeyAuth
//
// @Router			/api/v1/dapil/{id} [get]
func GetById(c *fiber.Ctx) error {
	data, totalRows, totalPages, err := Dapil.GetDataByID(c)
	if err != nil {
		err := c.Status(fiber.StatusBadRequest).SendString(err.Error())
		if err != nil {
			return nil
		}
	}

	err = c.Status(fiber.StatusOK).JSON(fiber.Map{"code": fiber.StatusOK, "message": "Data Stored.", "data": data, "totalRow": totalRows, "totalPages": totalPages})
	if err != nil {
		return nil
	}
	return err
}

// @Summary		Dapil Create
// @Description	Simpan data Dapil
// @Tags			Dapil
// @Accept			json
// @Produce		json
// @Param			dapil	body	Dapil.DapilInput	true	"Dapil Create"
// @Security		ApiKeyAuth
// @Router			/api/v1/dapil [post]
func Create(c *fiber.Ctx) error {
	var (
		err   error
		input Dapil.DapilInput
		model Dapil.DapilRequest
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

	model.ID = uuid.New()
	model.Code = input.Code
	model.Name = input.Name
	model.TotalVoters = input.TotalVoters
	model.Status = true
	model.CreatedAt = time.Now()
	model.CreatedBy = user.ID
	bulkAreas := []Dapil.DapilAreaInput{}
	for _, val := range input.Areas {
		villages, _ := Dapil.GetDapilVillage(val.DistrictId)
		for _, vil := range villages {
			area := Dapil.DapilAreaInput{}
			area.DapilId = model.ID
			area.ProvinceId = input.ProvinceId
			area.CityId = input.CityId
			area.DistrictId = val.DistrictId
			area.VillageId = vil.ID
			area.TotalVoters = 0
			area.Status = true
			area.CreatedBy = user.ID
			area.CreatedAt = time.Now()
			bulkAreas = append(bulkAreas, area)
		}
	}

	_, err = model.Create()
	if err != nil {
		err := c.Status(fiber.StatusBadRequest).SendString(err.Error())
		if err != nil {
			return nil
		}
	}

	_, err = Dapil.CreateBulkArea(bulkAreas)
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

// @Summary		Dapil Update
// @Description	Simpan data Dapil
// @Tags			Dapil
// @Accept			json
// @Produce		json
// @Param			dapil	body	Dapil.DapilInputUpdate	true	"Dapil Update"
// @Param			id			path	string								true	"Dapil ID"
// @Security		ApiKeyAuth
// @Router			/api/v1/dapil/{id} [post]
func Update(c *fiber.Ctx) error {
	var (
		err   error
		input Dapil.DapilInput
		model Dapil.DapilUpdateRequest
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
	model.TotalVoters = input.TotalVoters
	model.Code = input.Code
	model.Name = input.Name
	model.UpdatedAt = time.Now()
	model.UpdatedBy = user.ID
	bulkAreas := []Dapil.DapilAreaInput{}
	for _, val := range input.Areas {
		villages, _ := Dapil.GetDapilVillage(val.DistrictId)
		for _, vil := range villages {
			area := Dapil.DapilAreaInput{}
			area.DapilId = uuid.MustParse(Id)
			area.ProvinceId = input.ProvinceId
			area.CityId = input.CityId
			area.DistrictId = val.DistrictId
			area.VillageId = vil.ID
			area.TotalVoters = 0
			area.Status = true
			area.CreatedBy = user.ID
			area.CreatedAt = time.Now()
			bulkAreas = append(bulkAreas, area)
		}
	}
	_, err = model.Update(Id)
	if err != nil {
		err := c.Status(fiber.StatusBadRequest).SendString(err.Error())
		if err != nil {
			return nil
		}
	}

	err = Dapil.DeleteArea(Id)
	if err != nil {
		err := c.Status(fiber.StatusBadRequest).SendString(err.Error())
		if err != nil {
			return nil
		}
	}

	_, err = Dapil.CreateBulkArea(bulkAreas)
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

// @Summary		Dapil Delete
// @Description	Simpan data Dapil
// @Tags			Dapil
// @Accept			json
// @Produce		json
// @Param			dapil	body	Dapil.DapilUpdateRequest	true	"Dapil Update"
// @Param			id			path	string								true	"Dapil ID"
// @Security		ApiKeyAuth
// @Router			/api/v1/dapil/delete/{id} [post]
func Delete(c *fiber.Ctx) error {
	var (
		err    error
		params Dapil.DapilResult
	)
	if err := c.BodyParser(&params); err != nil {
		err := c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": fiber.StatusBadRequest, "message": err.Error()})
		if err != nil {
			return nil
		}
	}

	err = Dapil.Delete(params.ID)
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
