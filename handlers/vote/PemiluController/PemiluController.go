package PemiluController

import (
	"go_rest_api/repositories/vote/Pemilu"
	"log"

	"github.com/gofiber/fiber/v2"
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
	var err error
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
		err     error
		payload Pemilu.PemiluRequest
	)
	log.Println(payload)
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
		err     error
		payload Pemilu.PemiluUpdateRequest
	)
	log.Println(payload)
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
