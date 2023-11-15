package Company

import (
	"strings"
	"time"

	"go_rest_api/database"
	"go_rest_api/models/core"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type CompanyRequest struct {
	ID        uuid.UUID `json:"id"`
	GroupId   uuid.UUID `json:"group_id"`
	Code      int       `json:"code"`
	Name      string    `json:"name"`
	PathPhoto string    `json:"path_photo"`
	DapilId   uuid.UUID `json:"dapil_id"`
	Status    bool      `json:"status"`
	CreatedBy uuid.UUID `json:"created_by"`
	UpdatedBy uuid.UUID `json:"updated_by"`
	CreatedAt time.Time `json:"created_at"`
}

type CompanyUpdateRequest struct {
	GroupId   uuid.UUID `json:"group_id"`
	Code      int       `json:"code"`
	Name      string    `json:"name"`
	PathPhoto string    `json:"path_photo"`
	DapilId   uuid.UUID `json:"dapil_id"`
	Status    bool      `json:"status"`
	UpdatedBy uuid.UUID `json:"updated_by"`
	UpdatedAt time.Time `json:"updated_at"`
}

func GetAll(c *fiber.Ctx) (u []*core.Company, error error) {
	keyword := c.Query("q")

	sqlState := "SELECT * FROM companies "
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

func GetDataByID(id string) (c *core.Company, err error) {
	err = database.DB.Table("companies").Where("id = ?", id).Find(&c).Error
	if err != nil {
		return &core.Company{}, err
	}
	return c, nil
}

func (c *CompanyRequest) Create() (*CompanyRequest, error) {
	var err = database.DB.Table("companies").Create(&c).Error
	if err != nil {
		return &CompanyRequest{}, err
	}
	return c, nil
}

func (u *CompanyUpdateRequest) Update(id string) (*CompanyUpdateRequest, error) {
	var err = database.DB.Table("companies").Where("id = ?", id).Updates(&u).Error
	if err != nil {
		return u, err
	}
	return u, nil
}

func Delete(id string) error {
	var err = database.DB.Table("companies").Where("id = ?", id).Delete(core.Company{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *CompanyUpdateRequest) UpdateStatus(id string) (*CompanyUpdateRequest, error) {
	var err = database.DB.Table("companies").Where("id = ?", id).Update("status", u.Status).Error
	if err != nil {
		return &CompanyUpdateRequest{}, err
	}
	return u, nil
}
