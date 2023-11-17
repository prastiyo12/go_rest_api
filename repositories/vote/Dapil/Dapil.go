package Dapil

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

type DapilInput struct {
	Code        string  `json:"code"`
	Name        string  `json:"name"`
	TotalVoters int     `json:"total_voters"`
	ProvinceId  string  `json:"province_id"`
	CityId      string  `json:"city_id"`
	Areas       []DArea `json:"areas"`
}

type DArea struct {
	DistrictId string `json:"district_id"`
}

type ResDArea = []DArea

type DapilAreaInput struct {
	DapilId     uuid.UUID `json:"dapil_id"`
	ProvinceId  string    `json:"province_id"`
	CityId      string    `json:"city_id"`
	DistrictId  string    `json:"district_id"`
	VillageId   string    `json:"village_id"`
	TotalVoters int       `json:"total_voters"`
	Status      bool      `json:"status"`
	CreatedBy   uuid.UUID `json:"created_by"`
	UpdatedBy   uuid.UUID `json:"updated_by"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type DapilInputUpdate struct {
	TotalVoters int    `json:"total_voters"`
	Code        string `json:"code"`
	Name        string `json:"name"`
	Status      bool   `json:"status"`
}
type DapilRequest struct {
	ID          uuid.UUID `json:"id"`
	Code        string    `json:"code"`
	TotalVoters int       `json:"total_voters"`
	Name        string    `json:"name"`
	Status      bool      `json:"status"`
	CreatedBy   uuid.UUID `json:"created_by"`
	CreatedAt   time.Time `json:"created_at"`
}

type DapilUpdateRequest struct {
	Code        string    `json:"code"`
	TotalVoters int       `json:"total_voters"`
	Name        string    `json:"name"`
	Status      bool      `json:"status"`
	UpdatedBy   uuid.UUID `json:"updated_by"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type DapilResponse struct {
	ID          uuid.UUID `json:"id"`
	TotalVoters int       `json:"total_voters"`
	Code        string    `json:"code"`
	Name        string    `json:"name"`
	ProvinceId  string    `json:"province_id"`
	CityId      string    `json:"city_id"`
	Area        string    `json:"area"`
	Status      bool      `json:"status"`
	CreatedBy   uuid.UUID `json:"created_by"`
	UpdatedBy   uuid.UUID `json:"updated_by"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type DapilResult struct {
	ID          uuid.UUID `json:"id"`
	TotalVoters int       `json:"total_voters"`
	Code        string    `json:"code"`
	Name        string    `json:"name"`
	ProvinceId  string    `json:"province_id"`
	CityId      string    `json:"city_id"`
	Area        string    `json:"area"`
	Status      bool      `json:"status"`
	CreatedBy   uuid.UUID `json:"created_by"`
	UpdatedBy   uuid.UUID `json:"updated_by"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Areas       ResDArea  `json:"areas"`
}

type DapilVillageResult struct {
	DapilId      uuid.UUID `json:"id"`
	Name         string    `json:"name"`
	District     string    `json:"district"`
	UrbanVillage string    `json:"urban_village"`
}

func GetAll(c *fiber.Ctx) (result []DapilResult, tRow, tPages int, error error) {
	var u []*DapilResponse
	page, _ := strconv.Atoi(c.Query("page"))
	rows, _ := strconv.Atoi(c.Query("rows"))
	dir := c.Query("dir")
	sort := c.Query("sort")
	searchCode := c.Query("code")
	searchName := c.Query("name")
	qStatePage := "SELECT d.id,d.name, d.code, d.total_voters,d.status,da.province_id, da.city_id, string_agg(ad.district, ', ') AS area  "

	qStateTotal := "SELECT COUNT(*) as total_data "

	qState := " FROM (SELECT dapil_id,province_id, city_id, district_id from dapil_areas GROUP BY dapil_id,province_id, city_id, district_id) da "
	qState = qState + " left join address_districts ad on da.district_id = ad.id"
	qState = qState + " left join dapils d on da.dapil_id = d.id"
	qState = qState + " where d.deleted_at is null"

	if searchCode != "" {
		qState = qState + " AND lower(d.code) = '" + strings.ToLower(searchCode) + "'"
	}

	if searchName != "" {
		qState = qState + " AND lower(d.name) like '%" + strings.ToLower(searchName) + "%'"
	}

	qStateTotal = qStateTotal + qState
	//fmt.Println("qStateTotal :", qStateTotal)
	if err := database.DB.Raw(qStateTotal).Scan(&tRow).Error; err != nil {

		return nil, tRow, tPages, err
	}
	qState = qState + " group by d.id,d.name, d.code, d.total_voters, da.province_id, da.city_id"
	if dir != "" {
		qState = qState + " ORDER BY " + dir + " " + sort
	} else {
		qState = qState + " ORDER BY d.created_at DESC "
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
		return nil, tRow, int(tPages), err
	}

	for _, dt := range u {
		allAreas, _ := GetDataAreaByDapilId(dt.ID.String())
		dr := DapilResult{
			ID:          dt.ID,
			TotalVoters: dt.TotalVoters,
			Code:        dt.Code,
			Name:        dt.Name,
			ProvinceId:  dt.ProvinceId,
			CityId:      dt.CityId,
			Area:        dt.Area,
			Status:      dt.Status,
			CreatedBy:   dt.CreatedBy,
			UpdatedBy:   dt.UpdatedBy,
			CreatedAt:   dt.CreatedAt,
			UpdatedAt:   dt.UpdatedAt,
			Areas:       allAreas,
		}
		result = append(result, dr)

	}

	return result, tRow, int(tPages), nil
}

func GetDataAreaByDapilId(id string) (u ResDArea, err error) {

	qState := " SELECT district_id FROM dapil_areas WHERE dapil_id ='" + id + "' GROUP BY district_id"
	if err := database.DB.Raw(qState).Scan(&u).Error; err != nil {
		return u, err
	}

	return u, nil
}

func GetDataByID(c *fiber.Ctx) (u []DapilVillageResult, tRow, tPages int, error error) {
	page, _ := strconv.Atoi(c.Query("page"))
	rows, _ := strconv.Atoi(c.Query("rows"))
	dir := c.Query("dir")
	sort := c.Query("sort")
	dapilID := c.Params("id")
	searchName := c.Query("urban_village")
	qStatePage := "SELECT da.dapil_id,d.name, ad.district, av.urban_village  "

	qStateTotal := "SELECT COUNT(*) as total_data "

	qState := " FROM dapil_areas da  "
	qState = qState + " left join address_districts ad on da.district_id = ad.id"
	qState = qState + " left join address_villages av on da.village_id = av.id"
	qState = qState + " left join dapils d on da.dapil_id = d.id"
	qState = qState + " where d.deleted_at is null AND dapil_id = '" + dapilID + "'"

	if searchName != "" {
		qState = qState + " AND lower(av.urban_village) like '%" + strings.ToLower(searchName) + "%'"
	}

	qStateTotal = qStateTotal + qState
	if err := database.DB.Raw(qStateTotal).Scan(&tRow).Error; err != nil {

		return nil, tRow, tPages, err
	}
	if dir != "" {
		qState = qState + " ORDER BY " + dir + " " + sort
	} else {
		qState = qState + " ORDER BY ad.district ASC, av.urban_village ASC  "
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

func (c *DapilRequest) Create() (*DapilRequest, error) {
	var err = database.DB.Table("dapils").Create(&c).Error
	if err != nil {
		return &DapilRequest{}, err
	}
	return c, nil
}

func CreateBulkArea(c []DapilAreaInput) ([]DapilAreaInput, error) {
	var err = database.DB.Table("dapil_areas").CreateInBatches(c, 100).Error
	if err != nil {
		return []DapilAreaInput{}, err
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

func Delete(id uuid.UUID) error {
	var err = database.DB.Table("dapils").Where("id = ?", id).Delete(&vote.Dapil{}).Error
	if err != nil {
		return err
	}
	return nil
}

func DeleteArea(id string) error {
	var err error
	err = database.DB.Table("dapil_areas").
		Where("dapil_id = ?", id).
		Delete(vote.DapilArea{}).Error
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

func GetDapilVillage(districtId string) (u []*core.AddressVillage, error error) {
	qState := "SELECT * FROM address_villages "
	qState = qState + "WHERE district_id = '" + districtId + "'"
	qState = qState + " ORDER BY urban_village ASC"

	err := database.DB.Raw(qState).Scan(&u).Error
	if err != nil {
		return u, err
	}

	return u, nil
}
