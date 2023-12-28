package Relawan

import (
	"math"
	"strconv"
	"strings"
	"time"

	"go_rest_api/database"
	"go_rest_api/models/core"
	"go_rest_api/models/vote"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type RelawanRequest struct {
	ID             uuid.UUID `json:"id"`
	CompanyId      uuid.UUID `json:"company_id"`
	Name           string    `json:"name"`
	Email          string    `json:"email"`
	Phone          string    `json:"phone"`
	IdentityNumber string    `json:"identity_number"`
	BirthDate      time.Time `json:"birth_date"`
	BirthPlace     string    `json:"birth_place"`
	Gender         string    `json:"gender"`
	Address        string    `json:"address"`
	ProvinceId     string    `json:"province_id"`
	CityId         string    `json:"city_id"`
	DistrictId     string    `json:"district_id"`
	VillageId      string    `json:"village_id"`
	Rt             string    `json:"rt"`
	Rw             string    `json:"rw"`
	Longitude      float64   `json:"longitude"`
	Latitude       float64   `json:"latitude"`
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

type RelawanUpdateRequest struct {
	CompanyId      uuid.UUID `json:"company_id"`
	Email          string    `json:"email"`
	Name           string    `json:"name"`
	Phone          string    `json:"phone"`
	IdentityNumber string    `json:"identity_number"`
	BirthDate      time.Time `json:"birth_date"`
	BirthPlace     string    `json:"birth_place"`
	Gender         string    `json:"gender"`
	Address        string    `json:"address"`
	ProvinceId     string    `json:"province_id"`
	CityId         string    `json:"city_id"`
	DistrictId     string    `json:"district_id"`
	VillageId      string    `json:"village_id"`
	Rt             string    `json:"rt"`
	Rw             string    `json:"rw"`
	Longitude      float64   `json:"longitude"`
	Latitude       float64   `json:"latitude"`
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

type RelawanResponse struct {
	ID             uuid.UUID `json:"id"`
	CompanyId      uuid.UUID `json:"company_id"`
	Name           string    `json:"name"`
	Phone          string    `json:"phone"`
	Email          string    `json:"email"`
	IdentityNumber string    `json:"identity_number"`
	BirthDate      time.Time `json:"birth_date"`
	BirthPlace     string    `json:"birth_place"`
	Gender         string    `json:"gender"`
	Address        string    `json:"address"`
	ProvinceId     uuid.UUID `json:"province_id"`
	Province       string    `json:"province"`
	CityId         string    `json:"city_id"`
	City           string    `json:"city"`
	DistrictId     string    `json:"district_id"`
	District       string    `json:"district"`
	VillageId      string    `json:"village_id"`
	UrbanVillage   string    `json:"urban_village"`
	Rt             string    `json:"rt"`
	Rw             string    `json:"rw"`
	Longitude      float64   `json:"longitude"`
	Latitude       float64   `json:"latitude"`
	UserTypeId     uuid.UUID `json:"user_type_id"`
	UserId         uuid.UUID `json:"user_id"`
	PathPhoto      string    `json:"path_photo"`
	Info1          string    `json:"info1"`
	Info2          string    `json:"info2"`
	Info3          string    `json:"info3"`
	Status         bool      `json:"status"`
	CreatedBy      uuid.UUID `json:"created_by"`
	UpdatedBy      uuid.UUID `json:"updated_by"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

func GetAll(c *fiber.Ctx) (u []*RelawanResponse, tRow, tPages int, error error) {
	user := c.Locals("user").(core.UserResponse)
	page, _ := strconv.Atoi(c.Query("page"))
	rows, _ := strconv.Atoi(c.Query("rows"))
	dir := c.Query("dir")
	sort := c.Query("sort")
	searchName := c.Query("name")
	searchProvinceId := c.Query("province_id")
	searchCityId := c.Query("city_id")
	searchDisctrictId := c.Query("district_id")
	searchVillageId := c.Query("village_id")
	searchPhone := c.Query("phone")
	searchIdentityNumber := c.Query("identity_number")
	searchGender := c.Query("gender")
	qStatePage := "SELECT c.*, u.phone, t.email,address_provinces.province, address_cities.city,address_districts.district, address_villages.urban_village"

	qStateTotal := "SELECT COUNT(*) as total_data "

	qState := " FROM profile c left join users u on c.user_id = u.id"
	qState = qState + " left join address_provinces  on c.province_id = address_provinces.id"
	qState = qState + " left join address_cities  on c.city_id = address_cities.id"
	qState = qState + " left join address_districts  on c.district_id = address_districts.id"
	qState = qState + " left join address_villages  on c.village_id = address_villages.id"
	qState = qState + " WHERE c.company_id = '" + user.CompanyId.String() + "' AND c.deleted_at is NULL"

	if searchProvinceId != "" {
		qState = qState + " AND c.province_id = '" + searchProvinceId + "'"
	}

	if searchCityId != "" {
		qState = qState + " AND c.city_id = '" + searchCityId + "'"
	}

	if searchDisctrictId != "" {
		qState = qState + " AND c.district_id = '" + searchDisctrictId + "'"
	}

	if searchVillageId != "" {
		qState = qState + " AND c.village_id = '" + searchVillageId + "'"
	}

	if searchGender != "" {
		qState = qState + " AND lower(c.gender) = '" + strings.ToLower(searchGender) + "'"
	}

	if searchName != "" {
		qState = qState + " AND lower(c.name) like '%" + strings.ToLower(searchName) + "%'"
	}

	if searchPhone != "" {
		qState = qState + " AND lower(c.phone) like '%" + strings.ToLower(searchPhone) + "%'"
	}

	if searchIdentityNumber != "" {
		qState = qState + " AND lower(c.identity_number) like '%" + strings.ToLower(searchIdentityNumber) + "%'"
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

func GetDataByID(id string) (c *vote.Pemilih, err error) {
	err = database.DB.Table("pemilihs").Where("id = ?", id).Find(&c).Error
	if err != nil {
		return &vote.Pemilih{}, err
	}
	return c, nil
}
