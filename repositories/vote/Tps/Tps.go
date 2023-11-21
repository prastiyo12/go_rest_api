package Tps

import (
	"log"
	"math"
	"strconv"
	"strings"
	"time"

	"go_rest_api/database"
	"go_rest_api/models/vote"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type TpsRequest struct {
	ID          uuid.UUID `json:"id"`
	DapilId     uuid.UUID `json:"dapil_id"`
	DapilAreaId int       `json:"dapil_area_id"`
	Rw          string    `json:"rw"`
	Code        string    `json:"code"`
	Name        string    `json:"name"`
	Status      bool      `json:"status"`
	CreatedBy   uuid.UUID `json:"created_by"`
	CreatedAt   time.Time `json:"created_at"`
}

type TpsUpdateRequest struct {
	DapilId     uuid.UUID `json:"dapil_id"`
	DapilAreaId int       `json:"dapil_area_id"`
	Rw          string    `json:"rw"`
	Code        string    `json:"code"`
	Name        string    `json:"name"`
	Status      bool      `json:"status"`
	UpdatedBy   uuid.UUID `json:"updated_by"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type TpsResponse struct {
	ID           uuid.UUID  `json:"id"`
	DapilId      uuid.UUID  `json:"dapil_id"`
	DapilCode    string     `json:"dapil_code"`
	DapilName    string     `json:"dapil_name"`
	DapilAreaId  int        `json:"dapil_area_id"`
	City         string     `json:"city"`
	District     string     `json:"district"`
	UrbanVillage string     `json:"urban_village"`
	Rw           string     `json:"rw"`
	Code         string     `json:"code"`
	Name         string     `json:"name"`
	Status       *bool      `json:"status"`
	CreatedBy    uuid.UUID  `json:"created_by"`
	UpdatedBy    uuid.UUID  `json:"updated_by"`
	CreatedAt    *time.Time `json:"created_at"`
	UpdatedAt    *time.Time `json:"updated_at"`
}

type DapilOption struct {
	Id   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type DapilAreaOption struct {
	Id           int       `json:"id"`
	DapilId      uuid.UUID `json:"dapil_id"`
	ProvinceId   string    `json:"province_id"`
	CityId       string    `json:"city_id"`
	DistrictId   string    `json:"district_id"`
	District     string    `json:"district"`
	UrbanVillage string    `json:"urban_village"`
	VillageId    string    `json:"village_id"`
	TotalVoters  int       `json:"total_voters"`
}

func GetAll(c *fiber.Ctx) (u []*TpsResponse, tRow, tPages int, error error) {
	page, _ := strconv.Atoi(c.Query("page"))
	rows, _ := strconv.Atoi(c.Query("rows"))
	dir := c.Query("dir")
	sort := c.Query("sort")
	searchName := c.Query("name")
	searchCode := c.Query("code")
	searchRw := c.Query("rw")
	qStatePage := "SELECT c.*, t.code as dapil_code, t.name as dapil_name, v.city, w.district, x.urban_village "

	qStateTotal := "SELECT COUNT(*) as total_data "

	qState := " FROM tps c left join dapils t on c.dapil_id = t.id"
	qState = qState + " left join dapil_areas u on c.dapil_area_id = u.id"
	qState = qState + " left join address_cities v on u.city_id = v.id"
	qState = qState + " left join address_districts w on u.district_id = w.id"
	qState = qState + " left join address_villages x on u.village_id = x.id"
	qState = qState + " WHERE c.deleted_at is NULL"

	if searchName != "" {
		qState = qState + " AND lower(c.name) like '%" + strings.ToLower(searchName) + "%'"
	}

	if searchCode != "" {
		qState = qState + " AND lower(c.code) like '%" + strings.ToLower(searchCode) + "%'"
	}

	if searchRw != "" {
		qState = qState + " AND lower(c.rw) like '%" + strings.ToLower(searchRw) + "%'"
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

func GetDataByID(id string) (c *vote.Tps, err error) {
	err = database.DB.Table("tps").Where("id = ?", id).Find(&c).Error
	if err != nil {
		return &vote.Tps{}, err
	}
	return c, nil
}

func (c *TpsRequest) Create() (*TpsRequest, error) {
	log.Println(c)
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

func Delete(id uuid.UUID) error {
	var err = database.DB.Table("tps").Where("id = ?", id).Delete(&vote.Tps{}).Error
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

func GetAllDapil(c *fiber.Ctx) (u []*DapilOption, error error) {
	qState := "SELECT id, name FROM dapils WHERE deleted_at is null "
	if err := database.DB.Raw(qState).Scan(&u).Error; err != nil {
		return u, err
	}

	return u, nil
}

func GetAllDapilArea(c *fiber.Ctx) (u []*DapilAreaOption, error error) {
	dapilId := c.Query("dapil_id")
	qState := "SELECT a.*, b.district, c.urban_village FROM dapil_areas a"
	qState = qState + " left join address_districts b on a.district_id = b.id"
	qState = qState + " left join address_villages c on a.village_id = c.id"
	qState = qState + " WHERE dapil_id = '" + dapilId + "' "
	if err := database.DB.Raw(qState).Scan(&u).Error; err != nil {
		return u, err
	}

	return u, nil
}

func GetAllTps(c *fiber.Ctx) (u []*DapilOption, error error) {
	cityId := c.Query("city_id")
	qState := "SELECT t.id, t.name FROM tps t JOIN dapil_areas d on t.dapil_id = d.dapil_id WHERE t.deleted_at is null "
	if cityId != "" {
		qState = qState + " AND d.city_id = '" + cityId + "'"
	}
	qState = qState + " GROUP BY t.id"
	if err := database.DB.Raw(qState).Scan(&u).Error; err != nil {
		return u, err
	}

	return u, nil
}
