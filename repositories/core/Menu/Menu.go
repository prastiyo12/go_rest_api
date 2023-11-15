package Menu

import (
	"strings"
	"time"

	"go_rest_api/database"
	"go_rest_api/models/core"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type MenuRequest struct {
	ID        uuid.UUID `json:"id"`
	CompanyId uuid.UUID `json:"company_id"`
	ParentId  uuid.UUID `json:"parent_id"`
	Name      string    `json:"name"`
	Path      string    `json:"path"`
	Icon      string    `json:"icon"`
	Translate string    `json:"translate"`
	Status    bool      `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type MenuUpdateRequest struct {
	CompanyId uuid.UUID `json:"company_id"`
	ParentId  uuid.UUID `json:"parent_id"`
	Name      string    `json:"name"`
	Path      string    `json:"path"`
	Icon      string    `json:"icon"`
	Translate string    `json:"translate"`
	Status    bool      `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func GetAll(c *fiber.Ctx) (u []*core.Menu, error error) {
	keyword := c.Query("q")

	qState := "SELECT * FROM menus "
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

func GetDataByID(id string) (c *core.Menu, err error) {
	err = database.DB.Table("menus").Where("id = ?", id).Find(&c).Error
	if err != nil {
		return &core.Menu{}, err
	}
	return c, nil
}

func (c *MenuRequest) Create() (*MenuRequest, error) {
	var err = database.DB.Table("menus").Create(&c).Error
	if err != nil {
		return &MenuRequest{}, err
	}
	return c, nil
}

func (u *MenuUpdateRequest) Update(id string) (*MenuUpdateRequest, error) {
	var err = database.DB.Table("menus").Where("id = ?", id).Updates(&u).Error
	if err != nil {
		return u, err
	}
	return u, nil
}

func Delete(id string) error {
	var err = database.DB.Table("menus").Where("id = ?", id).Delete(core.Menu{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *MenuUpdateRequest) UpdateStatus(id string) (*MenuUpdateRequest, error) {
	var err = database.DB.Table("menus").Where("id = ?", id).Update("status", u.Status).Error
	if err != nil {
		return &MenuUpdateRequest{}, err
	}
	return u, nil
}

func GetAllMenu(c *fiber.Ctx) (u []*core.Menu, error error) {
	keyword := c.Query("q")

	qState := "SELECT  FROM menus "
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
