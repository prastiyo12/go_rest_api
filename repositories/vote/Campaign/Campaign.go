package Campaign

import (
	"strings"
	"time"

	"go_rest_api/database"
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

func GetAll(c *fiber.Ctx) (u []*vote.Campaign, error error) {
	keyword := c.Query("q")

	qState := "SELECT * FROM campaigns "
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
