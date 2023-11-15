package Pemilu

import (
	"time"

	"go_rest_api/database"
	"go_rest_api/models/vote"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type PemiluRequest struct {
	ID                 uuid.UUID  `json:"id"`
	CompanyId          uuid.UUID  `json:"company_id"`
	TpsId              uuid.UUID  `json:"tps_id"`
	TotalVoters        int        `json:"total_voters"`
	TotalVotersCompany int        `json:"total_voters_company"`
	TotalVotersGroup   int        `json:"total_voters_group"`
	TotalVotersOther   int        `json:"total_voters_other"`
	PathPhoto          string     `json:"path_photo"`
	Longitude          float64    `json:"longitude"`
	Latitude           float64    `json:"latitude"`
	Info1              string     `json:"info1"`
	Info2              string     `json:"info2"`
	Info3              string     `json:"info3"`
	Status             *bool      `json:"status"`
	CreatedBy          uuid.UUID  `json:"created_by"`
	CreatedAt          *time.Time `json:"created_at"`
}

type PemiluUpdateRequest struct {
	CompanyId          uuid.UUID  `json:"company_id"`
	TpsId              uuid.UUID  `json:"tps_id"`
	TotalVoters        int        `json:"total_voters"`
	TotalVotersCompany int        `json:"total_voters_company"`
	TotalVotersGroup   int        `json:"total_voters_group"`
	TotalVotersOther   int        `json:"total_voters_other"`
	PathPhoto          string     `json:"path_photo"`
	Longitude          float64    `json:"longitude"`
	Latitude           float64    `json:"latitude"`
	Info1              string     `json:"info1"`
	Info2              string     `json:"info2"`
	Info3              string     `json:"info3"`
	Status             *bool      `json:"status"`
	UpdatedBy          uuid.UUID  `json:"updated_by"`
	UpdatedAt          *time.Time `json:"updated_at"`
}

func GetAll(c *fiber.Ctx) (u []*vote.Pemilu, error error) {
	// keyword := c.Query("q")

	sqlState := "SELECT * FROM pemilus "
	// if keyword != "" {
	// 	sqlState = sqlState + "WHERE LOWER(name) LIKE '%" + strings.ToLower(keyword) + "%'"
	// }
	sqlState = sqlState + " ORDER BY created_at DESC"

	err := database.DB.Raw(sqlState).Scan(&u).Error
	if err != nil {
		return u, err
	}

	return u, nil
}

func GetDataByID(id string) (c *vote.Pemilu, err error) {
	err = database.DB.Table("pemilus").Where("id = ?", id).Find(&c).Error
	if err != nil {
		return &vote.Pemilu{}, err
	}
	return c, nil
}

func (c *PemiluRequest) Store() (*PemiluRequest, error) {
	var err = database.DB.Table("pemilus").Create(&c).Error
	if err != nil {
		return &PemiluRequest{}, err
	}
	return c, nil
}

func (u *PemiluUpdateRequest) Update(id string) (*PemiluUpdateRequest, error) {
	var err = database.DB.Table("pemilus").Where("id = ?", id).Updates(&u).Error
	if err != nil {
		return u, err
	}
	return u, nil
}

func Delete(id string) error {
	var err = database.DB.Table("pemilus").Where("id = ?", id).Delete(vote.Pemilu{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *PemiluUpdateRequest) UpdateStatus(id string) (*PemiluUpdateRequest, error) {
	var err = database.DB.Table("pemilus").Where("id = ?", id).Update("status", u.Status).Error
	if err != nil {
		return &PemiluUpdateRequest{}, err
	}
	return u, nil
}
