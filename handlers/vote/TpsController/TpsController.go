package TpsController

import (
	"go_rest_api/repositories/vote/Tps"
	"log"

	"github.com/gofiber/fiber/v2"
)

// @Summary		Tps
// @Description	Ambil data Tps
// @Tags			Tps
// @Accept			json
// @Produce		json
//
// @Security		ApiKeyAuth
//
// @Router			/api/v1/company [get]
func GetAll(c *fiber.Ctx) error {
	var err error
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
	var err error
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
		err     error
		payload Tps.TpsRequest
	)
	log.Println(payload)
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
		err     error
		payload Tps.TpsUpdateRequest
	)
	log.Println(payload)
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
		err     error
		payload Tps.TpsRequest
	)
	log.Println(payload)
	return err
}
