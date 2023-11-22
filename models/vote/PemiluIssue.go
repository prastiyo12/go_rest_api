package vote

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Issue struct {
	gorm.Model
	ID            uuid.UUID  `gorm:"type:uuid;primary_key" json:"id"`
	CompanyId     uuid.UUID  `gorm:"type:uuid" json:"company_id"`
	IssueTitle    string     `gorm:"type:varchar(255);not null" json:"issue_title"`
	Issue         string     `gorm:"type:varchar(255);not null" json:"issue"`
	IssueMaker    string     `gorm:"type:varchar(255);not null" json:"issue_maker"`
	IssueSolution string     `gorm:"type:varchar(255);not null" json:"issue_solution"`
	IssuePhoto    string     `gorm:"type:varchar(255);not null" json:"issue_photo"`
	Status        *bool      `gorm:"not null;default:false"`
	CreatedBy     uuid.UUID  `gorm:"type:uuid" json:"created_by"`
	UpdatedBy     uuid.UUID  `gorm:"type:uuid" json:"updated_by"`
	CreatedAt     *time.Time `gorm:"not null;default:now()"`
	UpdatedAt     *time.Time `gorm:"not null;default:now()"`
}

func (obj *Issue) BeforeCreate(tx *gorm.DB) (err error) {
	// UUID version 4
	obj.ID = uuid.New()
	return
}
