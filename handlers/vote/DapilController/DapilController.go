package DapilController

import (
	"go_rest_api/repositories/vote/Dapil"
	"log"

	"github.com/gofiber/fiber/v2"
)

// @Summary		Dapil
// @Description	Ambil data Dapil
// @Tags			Dapil
// @Accept			json
// @Produce		json
//
// @Security		ApiKeyAuth
//
// @Router			/api/v1/dapil [get]
func GetAll(c *fiber.Ctx) error {
	var err error
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
	var err error
	return err
}

// @Summary		Dapil Create
// @Description	Simpan data Dapil
// @Tags			Dapil
// @Accept			json
// @Produce		json
// @Param			dapil	body	Dapil.DapilRequest	true	"Dapil Create"
// @Security		ApiKeyAuth
// @Router			/api/v1/dapil [post]
func Create(c *fiber.Ctx) error {
	var (
		err     error
		payload Dapil.DapilRequest
	)
	log.Println(payload)
	return err
}

// @Summary		Dapil Update
// @Description	Simpan data Dapil
// @Tags			Dapil
// @Accept			json
// @Produce		json
// @Param			dapil	body	Dapil.DapilUpdateRequest	true	"Dapil Update"
// @Param			id			path	int								true	"Dapil ID"
// @Security		ApiKeyAuth
// @Router			/api/v1/dapil/{id} [post]
func Update(c *fiber.Ctx) error {
	var (
		err     error
		payload Dapil.DapilUpdateRequest
	)
	log.Println(payload)
	return err
}

// @Summary		Dapil Delete
// @Description	Simpan data Dapil
// @Tags			Dapil
// @Accept			json
// @Produce		json
// @Param			dapil	body	Dapil.DapilUpdateRequest	true	"Dapil Update"
// @Param			id			path	int								true	"Dapil ID"
// @Security		ApiKeyAuth
// @Router			/api/v1/dapil/delete/{id} [post]
func Delete(c *fiber.Ctx) error {
	var (
		err     error
		payload Dapil.DapilRequest
	)
	log.Println(payload)
	return err
}
