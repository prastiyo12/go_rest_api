package Campaign

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

type CampaignRequest struct {
	ID             uuid.UUID `json:"id"`
	CompanyId      uuid.UUID `json:"company_id"  validate:"required"`
	ProfileId      uuid.UUID `json:"profile_id" validate:"required"`
	Campaign       string    `json:"campaign" validate:"required"`
	CampaignTarget int       `json:"campaign_target"`
	CampaignAmount int       `json:"campaign_amount"`
	TpsId          uuid.UUID `json:"tps_id"`
	CampaignStatus int       `json:"campaign_status"`
	Status         bool      `json:"status"`
	CreatedBy      uuid.UUID `json:"created_by"`
	CreatedAt      time.Time `json:"created_at"`
}

type CampaignUpdateRequest struct {
	ID             uuid.UUID `json:"id"`
	CompanyId      uuid.UUID `json:"company_id"`
	ProfileId      uuid.UUID `json:"profile_id"`
	Campaign       string    `json:"campaign"`
	CampaignTarget int       `json:"campaign_target"`
	CampaignAmount int       `json:"campaign_amount"`
	TpsId          uuid.UUID `json:"tps_id"`
	CampaignStatus int       `json:"campaign_status"`
	Status         bool      `json:"status"`
	UpdatedBy      uuid.UUID `json:"updated_by"`
	UpdatedAt      time.Time `json:"updated_at"`
}

func GetAll(c *fiber.Ctx) (u []*vote.Campaign, tRow, tPages int, error error) {
	user := c.Locals("user").(core.UserResponse)
	page, _ := strconv.Atoi(c.Query("page"))
	rows, _ := strconv.Atoi(c.Query("rows"))
	dir := c.Query("dir")
	sort := c.Query("sort")
	searchCampaign := c.Query("campaign")
	// searchCampaignTarget := c.Query("campaign_target")
	searchTpsId := c.Query("tps_id")
	searchCampaignStatus := c.Query("campaign_status")
	qStatePage := "SELECT * "

	qStateTotal := "SELECT COUNT(*) as total_data "

	qState := " FROM campaigns WHERE company_id = '" + user.CompanyId.String() + "' AND deleted_at is NULL"

	if searchCampaignStatus != "" {
		qState = qState + " AND lower(campaign_status) = '" + strings.ToLower(searchCampaignStatus) + "'"
	}

	if searchTpsId != "" {
		qState = qState + " AND lower(tps_id) = '" + strings.ToLower(searchTpsId) + "'"
	}

	if searchCampaign != "" {
		qState = qState + " AND lower(campaign) like '%" + strings.ToLower(searchCampaign) + "%'"
	}

	qStateTotal = qStateTotal + qState
	if err := database.DB.Raw(qStateTotal).Scan(&tRow).Error; err != nil {

		return u, tRow, tPages, err
	}

	if dir != "" {
		qState = qState + " ORDER BY " + dir + " " + sort
	} else {
		qState = qState + " ORDER BY created_at DESC "
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

func GetDataByID(id string) (c *vote.Campaign, err error) {
	err = database.DB.Table("campaigns").Where("id = ?", id).Find(&c).Error
	if err != nil {
		return &vote.Campaign{}, err
	}
	return c, nil
}

func (c *CampaignRequest) Create() (*CampaignRequest, error) {
	var err = database.DB.Table("campaigns").Create(&c).Error
	if err != nil {
		return &CampaignRequest{}, err
	}
	return c, nil
}

func (u *CampaignUpdateRequest) Update(id string) (*CampaignUpdateRequest, error) {
	var err = database.DB.Table("campaigns").Where("id = ?", id).Updates(&u).Error
	if err != nil {
		return u, err
	}
	return u, nil
}

func Delete(id string) error {
	var err = database.DB.Table("campaigns").Where("id = ?", id).Delete(vote.Campaign{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *CampaignUpdateRequest) UpdateStatus(id string) (*CampaignUpdateRequest, error) {
	var err = database.DB.Table("campaigns").Where("id = ?", id).Update("status", u.Status).Error
	if err != nil {
		return &CampaignUpdateRequest{}, err
	}
	return u, nil
}
