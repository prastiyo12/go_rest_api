package Pemilu

import (
	"math"
	"strconv"
	"time"

	"go_rest_api/database"
	"go_rest_api/models/core"
	"go_rest_api/models/vote"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type PemiluRequest struct {
	ID                 uuid.UUID `json:"id"`
	CompanyId          uuid.UUID `json:"company_id"`
	TpsId              uuid.UUID `json:"tps_id"`
	TotalVoters        int       `json:"total_voters"`
	TotalVotersCompany int       `json:"total_voters_company"`
	TotalVotersGroup   int       `json:"total_voters_group"`
	TotalVotersOther   int       `json:"total_voters_other"`
	PathPhoto          string    `json:"path_photo"`
	Longitude          float64   `json:"longitude"`
	Latitude           float64   `json:"latitude"`
	Status             bool      `json:"status"`
	CreatedBy          uuid.UUID `json:"created_by"`
	CreatedAt          time.Time `json:"created_at"`
}

type PemiluUpdateRequest struct {
	CompanyId          uuid.UUID `json:"company_id"`
	TpsId              uuid.UUID `json:"tps_id"`
	TotalVoters        int       `json:"total_voters"`
	TotalVotersCompany int       `json:"total_voters_company"`
	TotalVotersGroup   int       `json:"total_voters_group"`
	TotalVotersOther   int       `json:"total_voters_other"`
	PathPhoto          string    `json:"path_photo"`
	Longitude          float64   `json:"longitude"`
	Latitude           float64   `json:"latitude"`
	Status             bool      `json:"status"`
	UpdatedBy          uuid.UUID `json:"updated_by"`
	UpdatedAt          time.Time `json:"updated_at"`
}

type PemiluResponse struct {
	ID                 uuid.UUID `json:"id"`
	CompanyId          uuid.UUID `json:"company_id"`
	TpsId              uuid.UUID `json:"tps_id"`
	TpsName            string    `json:"tps_name"`
	TotalVoters        int       `json:"total_voters"`
	TotalVotersCompany int       `json:"total_voters_company"`
	TotalVotersGroup   int       `json:"total_voters_group"`
	TotalVotersOther   int       `json:"total_voters_other"`
	PathPhoto          string    `json:"path_photo"`
	Longitude          float64   `json:"longitude"`
	Latitude           float64   `json:"latitude"`
	Status             bool      `json:"status"`
	CreatedBy          uuid.UUID `json:"created_by"`
	CreatedDate        string    `json:"created_date"`
}

func GetAll(c *fiber.Ctx) (u []*PemiluResponse, tRow, tPages int, error error) {
	user := c.Locals("user").(core.UserResponse)
	page, _ := strconv.Atoi(c.Query("page"))
	rows, _ := strconv.Atoi(c.Query("rows"))
	dir := c.Query("dir")
	sort := c.Query("sort")
	// searchName := c.Query("name")
	// searchCampaignTarget := c.Query("campaign_target")
	// searchTpsId := c.Query("tps_id")
	// searchPhone := c.Query("phone")
	// searchIdentityNumber := c.Query("identity_number")
	// searchGender := c.Query("gender")
	qStatePage := "SELECT obj.*, t.code as tps_code, t.name as tps_name, to_char(obj.created_at,'YYYY-MM-DD HH24:MI:SS') as created_date "

	qStateTotal := "SELECT COUNT(*) as total_data "

	qState := " FROM pemilus obj left join tps t on obj.tps_id = t.id "
	qState = qState + " WHERE obj.company_id = '" + user.CompanyId.String() + "' AND obj.deleted_at is NULL"

	// if searchGender != "" {
	// 	qState = qState + " AND lower(c.gender) = '" + strings.ToLower(searchGender) + "'"
	// }

	// if searchTpsId != "" {
	// 	qState = qState + " AND lower(c.tps_id) = '" + strings.ToLower(searchTpsId) + "'"
	// }

	// if searchName != "" {
	// 	qState = qState + " AND lower(c.name) like '%" + strings.ToLower(searchName) + "%'"
	// }

	// if searchPhone != "" {
	// 	qState = qState + " AND lower(c.phone) like '%" + strings.ToLower(searchPhone) + "%'"
	// }

	// if searchIdentityNumber != "" {
	// 	qState = qState + " AND lower(c.identity_number) like '%" + strings.ToLower(searchIdentityNumber) + "%'"
	// }

	qStateTotal = qStateTotal + qState
	if err := database.DB.Raw(qStateTotal).Scan(&tRow).Error; err != nil {

		return u, tRow, tPages, err
	}

	if dir != "" {
		qState = qState + " ORDER BY " + dir + " " + sort
	} else {
		qState = qState + " ORDER BY obj.created_at DESC "
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

func GetDataByID(id string) (c *PemiluResponse, err error) {
	err = database.DB.Table("pemilus").Where("id = ?", id).Find(&c).Error
	if err != nil {
		return &PemiluResponse{}, err
	}
	return c, nil
}

func (c *PemiluRequest) Create() (*PemiluRequest, error) {
	var err = database.DB.Table("pemilus").Create(&c).Error
	if err != nil {
		return &PemiluRequest{}, err
	}
	return c, nil
}

func (u *PemiluUpdateRequest) Update(id string) (*PemiluUpdateRequest, error) {
	var err = database.DB.Table("pemilus").Where("id = ?", id).Updates(&u).Error
	if err != nil {
		return u, err
	}
	return u, nil
}

func Delete(id string) error {
	var err = database.DB.Table("pemilus").Where("id = ?", id).Delete(vote.Pemilu{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *PemiluUpdateRequest) UpdateStatus(id string) (*PemiluUpdateRequest, error) {
	var err = database.DB.Table("pemilus").Where("id = ?", id).Update("status", u.Status).Error
	if err != nil {
		return &PemiluUpdateRequest{}, err
	}
	return u, nil
}
