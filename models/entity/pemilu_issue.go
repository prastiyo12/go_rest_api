package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Issue struct {
	gorm.Model
	ID            uuid.UUID  `gorm:"type:uuid;primary_key"`
	CompanyId     uuid.UUID  `gorm:"type:uuid"`
	Issue         string     `gorm:"type:varchar(255);not null" json:"issue"`
	IssueMaker    string     `gorm:"type:varchar(255);not null" json:"issue_maker"`
	IssueSolution string     `gorm:"type:varchar(255);not null" json:"issue_solution"`
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
