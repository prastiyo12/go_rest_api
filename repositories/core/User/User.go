package User

import (
	"math"
	"strconv"
	"strings"
	"time"

	"go_rest_api/database"
	"go_rest_api/models/core"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type UserRequest struct {
	ID             uuid.UUID `json:"id"`
	Name           string    `json:"name"`
	Email          string    `json:"email"`
	Password       string    `json:"password"`
	Role           uuid.UUID `json:"role"`
	CompanyId      uuid.UUID `json:"company_id"`
	Phone          string    `json:"phone"`
	Photo          string    `json:"photo"`
	FirebaseToken  string    `json:"firebase_token"`
	ActivationCode string    `json:"activation_code"`
	Status         bool      `json:"status"`
	CreatedAt      time.Time `json:"created_at"`
}

type UserUpdateRequest struct {
	Name           string    `json:"name"`
	Email          string    `json:"email"`
	Password       string    `json:"password"`
	Role           uuid.UUID `json:"role"`
	CompanyId      uuid.UUID `json:"company_id"`
	Phone          string    `json:"phone"`
	Photo          string    `json:"photo"`
	FirebaseToken  string    `json:"firebase_token"`
	ActivationCode string    `json:"activation_code"`
	Status         bool      `json:"status"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type UserResponse struct {
	ID             uuid.UUID `json:"id"`
	Name           string    `json:"name"`
	Email          string    `json:"email"`
	Password       string    `json:"password"`
	Role           uuid.UUID `json:"role"`
	CompanyId      uuid.UUID `json:"company_id"`
	CompanyName    string    `json:"company_name"`
	RoleName       string    `json:"role_name"`
	Phone          string    `json:"phone"`
	Photo          string    `json:"photo"`
	FirebaseToken  string    `json:"firebase_token"`
	ActivationCode string    `json:"activation_code"`
	Status         bool      `json:"status"`
	CreatedAt      time.Time `json:"created_at"`
}

func GetAll(c *fiber.Ctx) (u []*UserResponse, tRow, tPages int, error error) {
	page, _ := strconv.Atoi(c.Query("page"))
	rows, _ := strconv.Atoi(c.Query("rows"))
	dir := c.Query("dir")
	sort := c.Query("sort")
	searchName := c.Query("name")
	searchEmail := c.Query("email")
	searchPhone := c.Query("phone")
	qStatePage := "SELECT u.*, c.name as company_name, ur.name as role_name "

	qStateTotal := "SELECT COUNT(*) as total_data "

	qState := " FROM users u left join company c on u.company_id = c.id"
	qState = qState + " left join user_roles ur on u.role = ur.id"
	qState = qState + " WHERE c.deleted_at is NULL"

	if searchName != "" {
		qState = qState + " AND lower(c.name) like '%" + strings.ToLower(searchName) + "%'"
	}

	if searchEmail != "" {
		qState = qState + " AND lower(c.email) like '%" + strings.ToLower(searchEmail) + "%'"
	}

	if searchPhone != "" {
		qState = qState + " AND lower(c.phone) like '%" + strings.ToLower(searchPhone) + "%'"
	}

	qStateTotal = qStateTotal + qState
	if err := database.DB.Raw(qStateTotal).Scan(&tRow).Error; err != nil {

		return u, tRow, tPages, err
	}

	if dir != "" {
		qState = qState + " ORDER BY " + dir + " " + sort
	} else {
		qState = qState + " ORDER BY c.created_at DESC "
	}

	qState = qStatePage + qState
	start := 0
	if page > 1 {
		start = page
		if tRow <= rows {
			start = 0
		}
	}
	tPages = int(math.Ceil(float64(tRow) / float64(rows)))
	qState = qState + " OFFSET " + strconv.Itoa(start)
	qState = qState + " LIMIT " + strconv.Itoa(rows)

	if err := database.DB.Raw(qState).Scan(&u).Error; err != nil {
		return u, tRow, int(tPages), err
	}

	return u, tRow, int(tPages), nil
}

func GetDataByID(id string) (c *core.User, err error) {
	err = database.DB.Table("users").Where("id = ?", id).Find(&c).Error
	if err != nil {
		return &core.User{}, err
	}
	return c, nil
}

func (c *UserRequest) Create() (*UserRequest, error) {
	var err = database.DB.Table("users").Create(&c).Error
	if err != nil {
		return &UserRequest{}, err
	}
	return c, nil
}

func (u *UserUpdateRequest) Update(id string) (*UserUpdateRequest, error) {
	var err = database.DB.Table("users").Where("id = ?", id).Updates(&u).Error
	if err != nil {
		return u, err
	}
	return u, nil
}

func Delete(id string) error {
	var err = database.DB.Table("users").Where("id = ?", id).Delete(core.Company{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *UserUpdateRequest) UpdateStatus(id string) (*UserUpdateRequest, error) {
	var err = database.DB.Table("users").Where("id = ?", id).Update("status", u.Status).Error
	if err != nil {
		return &UserUpdateRequest{}, err
	}
	return u, nil
}
