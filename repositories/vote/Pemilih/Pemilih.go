package Pemilih

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

type PemilihRequest struct {
	ID             uuid.UUID `json:"id"`
	CompanyId      uuid.UUID `json:"company_id"`
	TpsId          uuid.UUID `json:"tps_id"`
	Name           string    `json:"name"`
	Phone          string    `json:"phone"`
	IdentityNumber string    `json:"identity_number"`
	// BirthDate      time.Time `json:"birth_date"`
	// BirthPlace     string    `json:"birth_place"`
	// Gender         string    `json:"gender"`
	// Address        string    `json:"address"`
	// ProvinceId     uuid.UUID `json:"province_id"`
	CityId     string `json:"city_id"`
	DistrictId string `json:"district_id"`
	VillageId  string `json:"village_id"`
	// Rt             string    `json:"rt"`
	// Rw             string    `json:"rw"`
	// Longitude      float64   `json:"longitude"`
	// Latitude       float64   `json:"latitude"`
	// UserTypeId     uuid.UUID `json:"user_type_id"`
	// UserId         uuid.UUID `json:"user_id"`
	PathPhoto string `json:"path_photo"`
	// Info1          string    `json:"info1"`
	// Info2          string    `json:"info2"`
	// Info3          string    `json:"info3"`
	Status    bool      `json:"status"`
	CreatedBy uuid.UUID `json:"created_by"`
	CreatedAt time.Time `json:"created_at"`
}

type PemilihUpdateRequest struct {
	CompanyId      uuid.UUID `json:"company_id"`
	TpsId          uuid.UUID `json:"tps_id"`
	Name           string    `json:"name"`
	Phone          string    `json:"phone"`
	IdentityNumber string    `json:"identity_number"`
	// BirthDate    time.Time `json:"birth_date"`
	// BirthPlace   string    `json:"birth_place"`
	// Gender       string    `json:"gender"`
	// Address      string    `json:"address"`
	// ProvinceId   uuid.UUID `json:"province_id"`
	CityId     string `json:"city_id"`
	DistrictId string `json:"district_id"`
	VillageId  string `json:"village_id"`
	// Rt           string    `json:"rt"`
	// Rw           string    `json:"rw"`
	// Longitude    float64   `json:"longitude"`
	// Latitude     float64   `json:"latitude"`
	// UserTypeId   uuid.UUID `json:"user_type_id"`
	// UserId       uuid.UUID `json:"user_id"`
	PathPhoto string `json:"path_photo"`
	// Info1        string    `json:"info1"`
	// Info2        string    `json:"info2"`
	// Info3        string    `json:"info3"`
	Status    bool      `json:"status"`
	UpdatedBy uuid.UUID `json:"updated_by"`
	UpdatedAt time.Time `json:"updated_at"`
}

type PemilihResponse struct {
	ID             uuid.UUID `json:"id"`
	CompanyId      uuid.UUID `json:"company_id"`
	TpsId          uuid.UUID `json:"tps_id"`
	TpsName        string    `json:"tps_name"`
	TpsCode        string    `json:"tps_code"`
	Name           string    `json:"name"`
	Phone          string    `json:"phone"`
	IdentityNumber string    `json:"identity_number"`
	// BirthDate      time.Time `json:"birth_date"`
	// BirthPlace     string    `json:"birth_place"`
	// Gender         string    `json:"gender"`
	// Address        string    `json:"address"`
	// ProvinceId     uuid.UUID `json:"province_id"`
	CityId       string `json:"city_id"`
	City         string `json:"city"`
	DistrictId   string `json:"district_id"`
	District     string `json:"district"`
	VillageId    string `json:"village_id"`
	UrbanVillage string `json:"urban_village"`
	// Rt             string    `json:"rt"`
	// Rw             string    `json:"rw"`
	// Longitude      float64   `json:"longitude"`
	// Latitude       float64   `json:"latitude"`
	// UserTypeId     uuid.UUID `json:"user_type_id"`
	// UserId         uuid.UUID `json:"user_id"`
	PathPhoto string `json:"path_photo"`
	// Info1          string    `json:"info1"`
	// Info2          string    `json:"info2"`
	// Info3          string    `json:"info3"`
	Status    bool      `json:"status"`
	CreatedBy uuid.UUID `json:"created_by"`
	UpdatedBy uuid.UUID `json:"updated_by"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func GetAll(c *fiber.Ctx) (u []*PemilihResponse, tRow, tPages int, error error) {
	user := c.Locals("user").(core.UserResponse)
	page, _ := strconv.Atoi(c.Query("page"))
	rows, _ := strconv.Atoi(c.Query("rows"))
	dir := c.Query("dir")
	sort := c.Query("sort")
	searchName := c.Query("name")
	// searchCampaignTarget := c.Query("campaign_target")
	searchTpsId := c.Query("tps_id")
	searchPhone := c.Query("phone")
	searchIdentityNumber := c.Query("identity_number")
	searchGender := c.Query("gender")
	qStatePage := "SELECT c.*, t.code as tps_code, t.name as tps_name, address_cities.city,address_districts.district, address_villages.urban_village"

	qStateTotal := "SELECT COUNT(*) as total_data "

	qState := " FROM pemilihs c left join tps t on c.tps_id = t.id "
	qState = qState + " left join address_cities  on c.city_id = address_cities.id"
	qState = qState + " left join address_districts  on c.district_id = address_districts.id"
	qState = qState + " left join address_villages  on c.village_id = address_villages.id"
	qState = qState + " WHERE c.company_id = '" + user.CompanyId.String() + "' AND c.deleted_at is NULL"

	if searchGender != "" {
		qState = qState + " AND lower(c.gender) = '" + strings.ToLower(searchGender) + "'"
	}

	if searchTpsId != "" {
		qState = qState + " AND lower(c.tps_id) = '" + strings.ToLower(searchTpsId) + "'"
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

func (c *PemilihRequest) Create() (*PemilihRequest, error) {
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
