package SummaryController

import (
	"go_rest_api/repositories/vote/Dashboard"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// @Summary		Summary
// @Description	Ambil data Dashboard
// @Tags			Tps
// @Accept			json
// @Produce		json
//
// @Security		ApiKeyAuth
//
// @Router			/api/v1/dashboard [get]
func GetSummary(c *fiber.Ctx) error {
	totalTps, err := Dashboard.GetTotalTps(c)
	totalPemilih, err := Dashboard.GetTotalPemilih(c)
	totalPemilihCalge, err := Dashboard.GetVoteResult(c)

	res := Dashboard.SummaryResponse{}
	res.TotalTps = totalTps.TotalTps
	res.TotalPemilih = totalPemilih.TotalPemilih
	res.TotalPemilihCaleg = totalPemilihCalge.TotalPemilihCaleg

	totalP, err := strconv.Atoi(totalPemilih.TotalPemilih)
	totalPC, err := strconv.Atoi(totalPemilihCalge.TotalPemilihCaleg)
	tPC := (totalP - totalPC)
	res.TotalPesaing = strconv.Itoa(tPC)

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
