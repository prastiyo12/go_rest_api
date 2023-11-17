package OptionController

import (
	"go_rest_api/repositories/core/Option"

	"github.com/gofiber/fiber/v2"
)

// @Summary		Province
// @Description	Ambil data Province
// @Tags			Option
// @Accept			json
// @Produce		json
//
// @Security		ApiKeyAuth
//
// @Router			/api/v1/option/province [get]
func GetProvince(c *fiber.Ctx) error {
	data, err := Option.GetProvince(c)

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

// @Summary		Cities
// @Description	Ambil data City
// @Tags			Option
// @Accept			json
// @Produce		json
// @Param			province_id	query	string	true	"Province ID"
//
// @Security		ApiKeyAuth
//
// @Router			/api/v1/option/city [get]
func GetCity(c *fiber.Ctx) error {
	data, err := Option.GetCity(c)
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

// @Summary		District
// @Description	Ambil data District
// @Tags			Option
// @Accept			json
// @Produce		json
// @Param			city_id	query	string	true	"City ID"
//
// @Security		ApiKeyAuth
//
// @Router			/api/v1/option/district [get]
func GetDistrict(c *fiber.Ctx) error {
	data, err := Option.GetDistrict(c)
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

// @Summary		Village
// @Description	Ambil data Village
// @Tags			Option
// @Accept			json
// @Produce		json
// @Param			district_id	query	string	true	"District ID"
//
// @Security		ApiKeyAuth
//
// @Router			/api/v1/option/village [get]
func GetVillage(c *fiber.Ctx) error {
	data, err := Option.GetVillage(c)
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
