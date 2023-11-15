package Issue

import (
	"strings"
	"time"

	"go_rest_api/database"
	"go_rest_api/models/vote"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type IssueRequest struct {
	ID            uuid.UUID  `json:"id"`
	CompanyId     uuid.UUID  `json:"company_id"`
	Issue         string     `json:"issue"`
	IssueMaker    string     `json:"issue_maker"`
	IssueSolution string     `json:"issue_solution"`
	Status        *bool      `json:"status"`
	CreatedBy     uuid.UUID  `json:"created_by"`
	CreatedAt     *time.Time `json:"created_at"`
}

type IssueUpdateRequest struct {
	CompanyId     uuid.UUID  `json:"company_id"`
	Issue         string     `json:"issue"`
	IssueMaker    string     `json:"issue_maker"`
	IssueSolution string     `json:"issue_solution"`
	Status        *bool      `json:"status"`
	UpdatedBy     uuid.UUID  `json:"updated_by"`
	UpdatedAt     *time.Time `json:"updated_at"`
}

func GetAll(c *fiber.Ctx) (u []*vote.Issue, error error) {
	keyword := c.Query("q")

	sqlState := "SELECT * FROM issues "
	if keyword != "" {
		sqlState = sqlState + "WHERE LOWER(issue) LIKE '%" + strings.ToLower(keyword) + "%'"
	}
	sqlState = sqlState + " ORDER BY created_at DESC"

	err := database.DB.Raw(sqlState).Scan(&u).Error
	if err != nil {
		return u, err
	}

	return u, nil
}

func GetDataByID(id string) (c *vote.Issue, err error) {
	err = database.DB.Table("issues").Where("id = ?", id).Find(&c).Error
	if err != nil {
		return &vote.Issue{}, err
	}
	return c, nil
}

func (c *IssueRequest) Store() (*IssueRequest, error) {
	var err = database.DB.Table("issues").Create(&c).Error
	if err != nil {
		return &IssueRequest{}, err
	}
	return c, nil
}

func (u *IssueUpdateRequest) Update(id string) (*IssueUpdateRequest, error) {
	var err = database.DB.Table("issues").Where("id = ?", id).Updates(&u).Error
	if err != nil {
		return u, err
	}
	return u, nil
}

func Delete(id string) error {
	var err = database.DB.Table("issues").Where("id = ?", id).Delete(vote.Issue{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *IssueUpdateRequest) UpdateStatus(id string) (*IssueUpdateRequest, error) {
	var err = database.DB.Table("issues").Where("id = ?", id).Update("status", u.Status).Error
	if err != nil {
		return &IssueUpdateRequest{}, err
	}
	return u, nil
}
