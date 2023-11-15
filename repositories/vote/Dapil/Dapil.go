package Dapil

import (
	"strings"
	"time"

	"go_rest_api/database"
	"go_rest_api/models/vote"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type DapilRequest struct {
	ID        uuid.UUID  `json:"id"`
	Code      string     `json:"code"`
	Name      string     `json:"name"`
	Status    *bool      `json:"status"`
	CreatedBy uuid.UUID  `json:"created_by"`
	CreatedAt *time.Time `json:"created_at"`
}

type DapilUpdateRequest struct {
	Code      string     `json:"code"`
	Name      string     `json:"name"`
	Status    *bool      `json:"status"`
	UpdatedBy uuid.UUID  `json:"updated_by"`
	UpdatedAt *time.Time `json:"updated_at"`
}

func GetAll(c *fiber.Ctx) (u []*vote.Dapil, error error) {
	keyword := c.Query("q")

	sqlState := "SELECT * FROM dapils "
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

func GetDataByID(id string) (c *vote.Dapil, err error) {
	err = database.DB.Table("dapils").Where("id = ?", id).Find(&c).Error
	if err != nil {
		return &vote.Dapil{}, err
	}
	return c, nil
}

func (c *DapilRequest) Store() (*DapilRequest, error) {
	var err = database.DB.Table("dapils").Create(&c).Error
	if err != nil {
		return &DapilRequest{}, err
	}
	return c, nil
}

func (u *DapilUpdateRequest) Update(id string) (*DapilUpdateRequest, error) {
	var err = database.DB.Table("dapils").Where("id = ?", id).Updates(&u).Error
	if err != nil {
		return u, err
	}
	return u, nil
}

func Delete(id string) error {
	var err = database.DB.Table("dapils").Where("id = ?", id).Delete(vote.Dapil{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *DapilUpdateRequest) UpdateStatus(id string) (*DapilUpdateRequest, error) {
	var err = database.DB.Table("dapils").Where("id = ?", id).Update("status", u.Status).Error
	if err != nil {
		return &DapilUpdateRequest{}, err
	}
	return u, nil
}
