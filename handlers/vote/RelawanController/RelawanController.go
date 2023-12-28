package RelawanController

import (
	"go_rest_api/models/core"
	"go_rest_api/repositories/core/Profile"
	"go_rest_api/repositories/core/SettingData"
	"go_rest_api/repositories/core/User"
	"go_rest_api/repositories/vote/Pemilih"
	"go_rest_api/repositories/vote/Relawan"
	"log"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// @Summary		Relawan
// @Description	Ambil data Relawan
// @Tags			Relawan
// @Accept			json
// @Produce		json
//
// @Security		ApiKeyAuth
//
// @Router			/api/v1/relawan [get]
func GetAll(c *fiber.Ctx) error {
	data, totalRows, totalPages, err := Relawan.GetAll(c)

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

// @Summary		Relawan By ID
// @Description	Ambil data Relawan By ID
// @Tags			Relawan
// @Accept			json
// @Produce		json
// @Param			id	path	int	true	"Relawan ID"
//
// @Security		ApiKeyAuth
//
// @Router			/api/v1/relawan/{id} [get]
func GetById(c *fiber.Ctx) error {
	var err error
	return err
}

// @Summary		Relawan Create
// @Description	Simpan data Relawan
// @Tags			Relawan
// @Accept			json
// @Produce		json
// @Param			issue	body	Relawan.RelawanRequest	true	"Relawan Create"
// @Security		ApiKeyAuth
// @Router			/api/v1/relawan [post]
func Create(c *fiber.Ctx) error {
	var (
		err     error
		input   Relawan.RelawanRequest
		profile Profile.ProfileRequest
		login   User.UserRequest
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

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("123456"), bcrypt.DefaultCost)
	if err != nil {
		err := c.Status(fiber.StatusBadRequest).SendString(err.Error())
		if err != nil {
			return nil
		}
	}
	settingRoleRelawan, err := SettingData.GetSettingValue("role_relawan")
	if err != nil {
		err := c.Status(fiber.StatusBadRequest).SendString(err.Error())
		if err != nil {
			return nil
		}
	}
	login.ID = uuid.New()
	login.Name = input.Name
	login.Email = input.Email
	login.Password = string(hashedPassword)
	login.Role = uuid.MustParse(settingRoleRelawan.Value)
	login.CompanyId = user.CompanyId
	login.Phone = input.Phone
	login.Status = true
	login.CreatedAt = time.Now()
	_, err = login.Create()
	if err != nil {
		err := c.Status(fiber.StatusBadRequest).SendString(err.Error())
		if err != nil {
			return nil
		}
	}

	profile.ID = uuid.New()
	profile.Status = true
	profile.UserId = login.ID
	profile.IdentityNumber = input.IdentityNumber
	profile.BirthDate = input.BirthDate
	profile.BirthPlace = input.BirthPlace
	profile.Gender = input.Gender
	profile.Address = input.Address
	profile.ProvinceId = input.ProvinceId
	profile.CityId = input.CityId
	profile.DistrictId = input.DistrictId
	profile.VillageId = input.DistrictId
	profile.Rt = input.Rt
	profile.Rw = input.Rw
	profile.UserTypeId = input.UserTypeId
	profile.CompanyId = user.CompanyId
	profile.CreatedAt = time.Now()
	profile.CreatedBy = user.ID

	_, err = profile.Create()
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

// @Summary		Relawan Update
// @Description	Simpan data Relawan
// @Tags			Relawan
// @Accept			json
// @Produce		json
// @Param			issue	body	Relawan.RelawanUpdateRequest	true	"Relawan Update"
// @Param			id			path	int								true	"Relawan ID"
// @Security		ApiKeyAuth
// @Router			/api/v1/pemilih/{id} [post]
func Update(c *fiber.Ctx) error {
	var (
		err     error
		input   Relawan.RelawanUpdateRequest
		profile Profile.ProfileUpdateRequest
		login   User.UserUpdateRequest
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

	login.Name = input.Name
	login.Email = input.Email
	login.Phone = input.Phone
	login.Status = true
	login.UpdatedAt = time.Now()
	_, err = login.Update(input.UserId.String())
	if err != nil {
		err := c.Status(fiber.StatusBadRequest).SendString(err.Error())
		if err != nil {
			return nil
		}
	}

	var Id = c.Params("id")
	profile.Status = true
	profile.UserId = input.UserId
	profile.IdentityNumber = input.IdentityNumber
	profile.BirthDate = input.BirthDate
	profile.BirthPlace = input.BirthPlace
	profile.Gender = input.Gender
	profile.Address = input.Address
	profile.ProvinceId = input.ProvinceId
	profile.CityId = input.CityId
	profile.DistrictId = input.DistrictId
	profile.VillageId = input.DistrictId
	profile.Rt = input.Rt
	profile.Rw = input.Rw
	profile.UserTypeId = input.UserTypeId
	profile.CompanyId = user.CompanyId
	profile.UpdatedAt = time.Now()
	profile.UpdatedBy = user.ID

	_, err = profile.Update(Id)
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

// @Summary		Relawan Delete
// @Description	Simpan data Relawan
// @Tags			Relawan
// @Accept			json
// @Produce		json
// @Param			issue	body	Relawan.RelawanUpdateRequest	true	"Relawan Update"
// @Param			id			path	int								true	"Relawan ID"
// @Security		ApiKeyAuth
// @Router			/api/v1/relawan/delete/{id} [post]
func Delete(c *fiber.Ctx) error {
	var (
		err     error
		payload Pemilih.PemilihRequest
	)
	log.Println(payload)
	return err
}
