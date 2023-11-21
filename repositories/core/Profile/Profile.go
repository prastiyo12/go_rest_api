package Profile

import (
	"strings"
	"time"

	"go_rest_api/database"
	"go_rest_api/models/core"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type ProfileRequest struct {
	ID             uuid.UUID `json:"id"`
	CompanyId      uuid.UUID `json:"company_id"`
	Name           string    `json:"name"`
	IdentityNumber string    `json:"identity_number"`
	BirthDate      time.Time `json:"birth_date"`
	BirthPlace     string    `json:"birth_place"`
	Gender         string    `json:"gender"`
	Address        string    `json:"address"`
	ProvinceId     uuid.UUID `json:"province_id"`
	CityId         uuid.UUID `json:"city_id"`
	DistrictId     uuid.UUID `json:"district_id"`
	VillageId      uuid.UUID `json:"village_id"`
	Rt             string    `json:"rt"`
	Rw             string    `json:"rw"`
	UserTypeId     uuid.UUID `json:"user_type_id"`
	UserId         uuid.UUID `json:"user_id"`
	PathPhoto      string    `json:"path_photo"`
	Info1          string    `json:"info1"`
	Info2          string    `json:"info2"`
	Info3          string    `json:"info3"`
	Status         bool      `json:"status"`
	CreatedBy      uuid.UUID `json:"created_by"`
	CreatedAt      time.Time `json:"created_at"`
}

type ProfileUpdateRequest struct {
	CompanyId      uuid.UUID `json:"company_id"`
	Name           string    `json:"name"`
	IdentityNumber string    `json:"identity_number"`
	BirthDate      time.Time `json:"birth_date"`
	BirthPlace     string    `json:"birth_place"`
	Gender         string    `json:"gender"`
	Address        string    `json:"address"`
	ProvinceId     uuid.UUID `json:"province_id"`
	CityId         uuid.UUID `json:"city_id"`
	DistrictId     uuid.UUID `json:"district_id"`
	VillageId      uuid.UUID `json:"village_id"`
	Rt             string    `json:"rt"`
	Rw             string    `json:"rw"`
	UserTypeId     uuid.UUID `json:"user_type_id"`
	UserId         uuid.UUID `json:"user_id"`
	PathPhoto      string    `json:"path_photo"`
	Info1          string    `json:"info1"`
	Info2          string    `json:"info2"`
	Info3          string    `json:"info3"`
	Status         bool      `json:"status"`
	UpdatedBy      uuid.UUID `json:"updated_by"`
	UpdatedAt      time.Time `json:"updated_at"`
}

func GetAll(c *fiber.Ctx) (u []*core.Profile, error error) {
	keyword := c.Query("q")

	qState := "SELECT * FROM profiles "
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

func GetDataByID(id string) (c *core.Profile, err error) {
	err = database.DB.Table("profiles").Where("id = ?", id).Find(&c).Error
	if err != nil {
		return &core.Profile{}, err
	}
	return c, nil
}

func (c *ProfileRequest) Create() (*ProfileRequest, error) {
	var err = database.DB.Table("profiles").Create(&c).Error
	if err != nil {
		return &ProfileRequest{}, err
	}
	return c, nil
}

func (u *ProfileUpdateRequest) Update(id string) (*ProfileUpdateRequest, error) {
	var err = database.DB.Table("profiles").Where("id = ?", id).Updates(&u).Error
	if err != nil {
		return u, err
	}
	return u, nil
}

func Delete(id string) error {
	var err = database.DB.Table("profiles").Where("id = ?", id).Delete(core.Profile{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *ProfileUpdateRequest) UpdateStatus(id string) (*ProfileUpdateRequest, error) {
	var err = database.DB.Table("profiles").Where("id = ?", id).Update("status", u.Status).Error
	if err != nil {
		return &ProfileUpdateRequest{}, err
	}
	return u, nil
}

func GetAllUserTye(c *fiber.Ctx, companyId uuid.UUID) (u []*core.UserType, error error) {
	qState := "SELECT id, name FROM user_types WHERE deleted_at is null AND company_id = '" + companyId.String() + "' "
	if err := database.DB.Raw(qState).Scan(&u).Error; err != nil {
		return u, err
	}
	return u, nil
}

func GetAllUserRelawan(c *fiber.Ctx, companyId uuid.UUID) (u []*core.User, error error) {
	qState := "SELECT id, name FROM users WHERE deleted_at is null AND company_id = '" + companyId.String() + "' "
	if err := database.DB.Raw(qState).Scan(&u).Error; err != nil {
		return u, err
	}
	return u, nil
}
