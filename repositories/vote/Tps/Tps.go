package Tps

import (
	"strings"
	"time"

	"go_rest_api/database"
	"go_rest_api/models/vote"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type TpsRequest struct {
	ID        uuid.UUID  `json:"id"`
	DapilId   uuid.UUID  `json:"dapil_id"`
	VillageId uuid.UUID  `json:"village_id"`
	Rw        string     `json:"rw"`
	Code      string     `json:"code"`
	Name      string     `json:"name"`
	Status    *bool      `json:"status"`
	CreatedBy uuid.UUID  `json:"created_by"`
	CreatedAt *time.Time `json:"created_at"`
}

type TpsUpdateRequest struct {
	DapilId   uuid.UUID  `json:"dapil_id"`
	VillageId uuid.UUID  `json:"village_id"`
	Rw        string     `json:"rw"`
	Code      string     `json:"code"`
	Name      string     `json:"name"`
	Status    *bool      `json:"status"`
	UpdatedBy uuid.UUID  `json:"updated_by"`
	UpdatedAt *time.Time `json:"updated_at"`
}

func GetAll(c *fiber.Ctx) (u []*vote.Tps, error error) {
	keyword := c.Query("q")

	sqlState := "SELECT * FROM tps "
	if keyword != "" {
		sqlState = sqlState + "WHERE LOWER(name) LIKE '%" + strings.ToLower(keyword) + "%'"
	}
	sqlState = sqlState + " ORDER BY created_at DESC"

	err := database.DB.Raw(sqlState).Scan(&u).Error
	if err != nil {
		return u, err
	}

	return u, nil
}

func GetDataByID(id string) (c *vote.Tps, err error) {
	err = database.DB.Table("tps").Where("id = ?", id).Find(&c).Error
	if err != nil {
		return &vote.Tps{}, err
	}
	return c, nil
}

func (c *TpsRequest) Store() (*TpsRequest, error) {
	var err = database.DB.Table("tps").Create(&c).Error
	if err != nil {
		return &TpsRequest{}, err
	}
	return c, nil
}

func (u *TpsUpdateRequest) Update(id string) (*TpsUpdateRequest, error) {
	var err = database.DB.Table("tps").Where("id = ?", id).Updates(&u).Error
	if err != nil {
		return u, err
	}
	return u, nil
}

func Delete(id string) error {
	var err = database.DB.Table("tps").Where("id = ?", id).Delete(vote.Tps{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *TpsUpdateRequest) UpdateStatus(id string) (*TpsUpdateRequest, error) {
	var err = database.DB.Table("tps").Where("id = ?", id).Update("status", u.Status).Error
	if err != nil {
		return &TpsUpdateRequest{}, err
	}
	return u, nil
}
