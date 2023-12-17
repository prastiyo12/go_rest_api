package Dashboard

import (
	"go_rest_api/database"
	"go_rest_api/models/core"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type SummaryResponse struct {
	TotalTps          string `json:"total_tps"`
	TotalPemilih      string `json:"total_pemilih"`
	TotalPemilihCaleg string `json:"total_pemilih_caleg"`
	TotalPesaing      string `json:"total_pesaing"`
	CompanyName       string `json:"company_name"`
	PathPhoto         string `json:"path_photo"`
}

type TotalTpsRes struct {
	TotalTps string `json:"total_tps"`
}

type TotalPemilihRes struct {
	TotalPemilih string `json:"total_pemilih"`
}

type TotalPemilihCalegRes struct {
	TotalPemilihCaleg string `json:"total_pemilih_caleg"`
}

type CompanyRes struct {
	GroupId   uuid.UUID `json:"group_id"`
	Code      int       `json:"code"`
	Name      string    `json:"name"`
	PathPhoto string    `json:"path_photo"`
	DapilId   uuid.UUID `json:"dapil_id"`
	Status    bool      `json:"status"`
	UpdatedBy uuid.UUID `json:"updated_by"`
	UpdatedAt time.Time `json:"updated_at"`
}

func GetTotalTps(c *fiber.Ctx) (u TotalTpsRes, error error) {
	user := c.Locals("user").(core.UserResponse)
	qState := "SELECT count(*) as total_tps FROM tps t JOIN companies c ON t.dapil_id = c.dapil_id where c.id = '" + user.CompanyId.String() + "'"
	if err := database.DB.Raw(qState).Scan(&u).Error; err != nil {
		return u, err
	}
	return u, nil
}

func GetTotalPemilih(c *fiber.Ctx) (u TotalPemilihRes, error error) {
	user := c.Locals("user").(core.UserResponse)
	qState := "SELECT count(*) as total_pemilih FROM pemilihs where company_id = '" + user.CompanyId.String() + "'"
	if err := database.DB.Raw(qState).Scan(&u).Error; err != nil {
		return u, err
	}
	return u, nil
}

func GetVoteResult(c *fiber.Ctx) (u TotalPemilihCalegRes, error error) {
	user := c.Locals("user").(core.UserResponse)
	qState := "SELECT count(*) as total_pemilih_caleg FROM pemilus where company_id = '" + user.CompanyId.String() + "'"
	if err := database.DB.Raw(qState).Scan(&u).Error; err != nil {
		return u, err
	}
	return u, nil
}

func GetCompany(c *fiber.Ctx) (u CompanyRes, error error) {
	user := c.Locals("user").(core.UserResponse)
	qState := "SELECT * FROM companies where id = '" + user.CompanyId.String() + "'"
	if err := database.DB.Raw(qState).Scan(&u).Error; err != nil {
		return u, err
	}
	return u, nil
}
