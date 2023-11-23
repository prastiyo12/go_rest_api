package Issue

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

type IssueRequest struct {
	ID            uuid.UUID `json:"id"`
	CompanyId     uuid.UUID `json:"company_id"`
	Issue         string    `json:"issue"`
	IssueTitle    string    `json:"issue_title"`
	IssuePhoto    string    `json:"issue_photo"`
	IssueMaker    string    `json:"issue_maker"`
	IssueSolution string    `json:"issue_solution"`
	Status        bool      `json:"status"`
	CreatedBy     uuid.UUID `json:"created_by"`
	CreatedAt     time.Time `json:"created_at"`
}

type IssueUpdateRequest struct {
	CompanyId     uuid.UUID `json:"company_id"`
	Issue         string    `json:"issue"`
	IssueTitle    string    `json:"issue_title"`
	IssuePhoto    string    `json:"issue_photo"`
	IssueMaker    string    `json:"issue_maker"`
	IssueSolution string    `json:"issue_solution"`
	Status        bool      `json:"status"`
	UpdatedBy     uuid.UUID `json:"updated_by"`
	UpdatedAt     time.Time `json:"updated_at"`
}

func GetAll(c *fiber.Ctx) (u []*vote.Issue, tRow, tPages int, error error) {
	user := c.Locals("user").(core.UserResponse)
	page, _ := strconv.Atoi(c.Query("page"))
	rows, _ := strconv.Atoi(c.Query("rows"))
	dir := c.Query("dir")
	sort := c.Query("sort")
	searchIssue := c.Query("issue")
	searchIssueMaker := c.Query("issue_maker")
	qStatePage := "SELECT * "

	qStateTotal := "SELECT COUNT(*) as total_data "

	qState := " FROM issues WHERE company_id = '" + user.CompanyId.String() + "' AND deleted_at is NULL"

	if searchIssueMaker != "" {
		qState = qState + " AND lower(issue_maker) = '" + strings.ToLower(searchIssueMaker) + "'"
	}

	if searchIssue != "" {
		qState = qState + " AND (lower(issue) like '%" + strings.ToLower(searchIssue) + "%'  "
		qState = qState + " OR lower(issue_title) like '%" + strings.ToLower(searchIssue) + "%' ) "
	}

	qStateTotal = qStateTotal + qState
	//fmt.Println("qStateTotal :", qStateTotal)
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

func GetDataByID(id string) (c *vote.Issue, err error) {
	err = database.DB.Table("issues").Where("id = ?", id).Find(&c).Error
	if err != nil {
		return &vote.Issue{}, err
	}
	return c, nil
}

func (c *IssueRequest) Create() (*IssueRequest, error) {
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
