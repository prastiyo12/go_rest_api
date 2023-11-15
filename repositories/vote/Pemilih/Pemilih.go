package Pemilih

import (
	"strings"
	"time"

	"go_rest_api/database"
	"go_rest_api/models/vote"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type PemilihRequest struct {
	ID             uuid.UUID  `json:"id"`
	CompanyId      uuid.UUID  `json:"company_id"`
	TpsId          uuid.UUID  `json:"tps_id"`
	Name           string     `json:"name"`
	Phone          string     `json:"phone"`
	IdentityNumber string     `json:"idvote_number"`
	BirthDate      time.Time  `json:"birth_date"`
	BirthPlace     string     `json:"birth_place"`
	Gender         string     `json:"gender"`
	Address        string     `json:"address"`
	ProvinceId     uuid.UUID  `json:"province_id"`
	CityId         uuid.UUID  `json:"city_id"`
	DistrictId     uuid.UUID  `json:"district_id"`
	VillageId      uuid.UUID  `json:"village_id"`
	Rt             string     `json:"rt"`
	Rw             string     `json:"rw"`
	Longitude      float64    `json:"longitude"`
	Latitude       float64    `json:"latitude"`
	UserTypeId     uuid.UUID  `json:"user_type_id"`
	UserId         uuid.UUID  `json:"user_id"`
	PathPhoto      string     `json:"path_photo"`
	Info1          string     `json:"info1"`
	Info2          string     `json:"info2"`
	Info3          string     `json:"info3"`
	Status         *bool      `json:"status"`
	CreatedBy      uuid.UUID  `json:"created_by"`
	CreatedAt      *time.Time `json:"created_at"`
}

type PemilihUpdateRequest struct {
	CompanyId    uuid.UUID  `json:"company_id"`
	TpsId        uuid.UUID  `json:"tps_id"`
	Name         string     `json:"name"`
	Phone        string     `json:"phone"`
	IdvoteNumber string     `json:"idvote_number"`
	BirthDate    time.Time  `json:"birth_date"`
	BirthPlace   string     `json:"birth_place"`
	Gender       string     `json:"gender"`
	Address      string     `json:"address"`
	ProvinceId   uuid.UUID  `json:"province_id"`
	CityId       uuid.UUID  `json:"city_id"`
	DistrictId   uuid.UUID  `json:"district_id"`
	VillageId    uuid.UUID  `json:"village_id"`
	Rt           string     `json:"rt"`
	Rw           string     `json:"rw"`
	Longitude    float64    `json:"longitude"`
	Latitude     float64    `json:"latitude"`
	UserTypeId   uuid.UUID  `json:"user_type_id"`
	UserId       uuid.UUID  `json:"user_id"`
	PathPhoto    string     `json:"path_photo"`
	Info1        string     `json:"info1"`
	Info2        string     `json:"info2"`
	Info3        string     `json:"info3"`
	Status       *bool      `json:"status"`
	UpdatedBy    uuid.UUID  `json:"updated_by"`
	UpdatedAt    *time.Time `json:"updated_at"`
}

func GetAll(c *fiber.Ctx) (u []*vote.Pemilih, error error) {
	keyword := c.Query("q")

	sqlState := "SELECT * FROM pemilihs "
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

func GetDataByID(id string) (c *vote.Pemilih, err error) {
	err = database.DB.Table("pemilihs").Where("id = ?", id).Find(&c).Error
	if err != nil {
		return &vote.Pemilih{}, err
	}
	return c, nil
}

func (c *PemilihRequest) Store() (*PemilihRequest, error) {
	var err = database.DB.Table("pemilihs").Create(&c).Error
	if err != nil {
		return &PemilihRequest{}, err
	}
	return c, nil
}

func (u *PemilihUpdateRequest) Update(id string) (*PemilihUpdateRequest, error) {
	var err = database.DB.Table("pemilihs").Where("id = ?", id).Updates(&u).Error
	if err != nil {
		return u, err
	}
	return u, nil
}

func Delete(id string) error {
	var err = database.DB.Table("pemilihs").Where("id = ?", id).Delete(vote.Pemilih{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *PemilihUpdateRequest) UpdateStatus(id string) (*PemilihUpdateRequest, error) {
	var err = database.DB.Table("pemilihs").Where("id = ?", id).Update("status", u.Status).Error
	if err != nil {
		return &PemilihUpdateRequest{}, err
	}
	return u, nil
}
