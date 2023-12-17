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
	totalPemilih, _ := Dashboard.GetTotalPemilih(c)
	totalPemilihCalge, _ := Dashboard.GetVoteResult(c)
	companies, _ := Dashboard.GetCompany(c)
	news, _ := Dashboard.GetNews(c)

	res := Dashboard.SummaryResponse{}
	res.TotalTps = totalTps.TotalTps
	res.TotalPemilih = totalPemilih.TotalPemilih
	res.TotalPemilihCaleg = totalPemilihCalge.TotalPemilihCaleg
	res.CompanyName = companies.Name
	res.PathPhoto = companies.PathPhoto
	res.News = news

	totalP, _ := strconv.Atoi(totalPemilih.TotalPemilih)
	totalPC, _ := strconv.Atoi(totalPemilihCalge.TotalPemilihCaleg)
	tPC := (totalP - totalPC)
	res.TotalPesaing = strconv.Itoa(tPC)

	if err != nil {
		err := c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": fiber.StatusBadRequest, "message": err.Error()})
		if err != nil {
			return err
		}
		return err
	}

	err = c.Status(fiber.StatusOK).JSON(fiber.Map{"code": fiber.StatusOK, "message": "Data Stored.", "data": res})
	if err != nil {
		return err
	}
	return err
}
