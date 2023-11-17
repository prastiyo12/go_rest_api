package Role

import (
	"strings"
	"time"

	"go_rest_api/database"
	"go_rest_api/models/core"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type RoleInput struct {
	CompanyId uuid.UUID `json:"company_id"`
	Name      string    `json:"name"`
	Status    bool      `json:"status"`
}

type RoleInputUpdate struct {
	Name   string `json:"name"`
	Status bool   `json:"status"`
}

type RoleRequest struct {
	ID        uuid.UUID `json:"id"`
	CompanyId uuid.UUID `json:"company_id"`
	Name      string    `json:"name"`
	Status    bool      `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

type RoleUpdateRequest struct {
	Name      string    `json:"name"`
	Status    bool      `json:"status"`
	UpdatedAt time.Time `json:"updated_at"`
}

func GetAll(c *fiber.Ctx) (u []*core.UserRole, error error) {
	keyword := c.Query("q")

	qState := "SELECT * FROM user_roles "
	if keyword != "" {
		qState = qState + "WHERE LOWER(name) LIKE '%" + strings.ToLower(keyword) + "%'"
	}
	qState = qState + " ORDER BY created_at DESC"

	err := database.DB.Raw(qState).Scan(&u).Error
	if err != nil {
		return u, err
	}

	return u, nil
}

func GetDataByID(id string) (c *core.UserRole, err error) {
	err = database.DB.Table("user_roles").Where("id = ?", id).Find(&c).Error
	if err != nil {
		return &core.UserRole{}, err
	}
	return c, nil
}

func (c *RoleRequest) Create() (*RoleRequest, error) {
	var err = database.DB.Table("user_roles").Create(&c).Error
	if err != nil {
		return &RoleRequest{}, err
	}
	return c, nil
}

func (u *RoleUpdateRequest) Update(id string) (*RoleUpdateRequest, error) {
	var err = database.DB.Table("user_roles").Where("id = ?", id).Updates(&u).Error
	if err != nil {
		return u, err
	}
	return u, nil
}

func Delete(id string) error {
	var err = database.DB.Table("user_roles").Where("id = ?", id).Delete(core.UserRole{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *RoleUpdateRequest) UpdateStatus(id string) (*RoleUpdateRequest, error) {
	var err = database.DB.Table("menus").Where("id = ?", id).Update("status", u.Status).Error
	if err != nil {
		return &RoleUpdateRequest{}, err
	}
	return u, nil
}
