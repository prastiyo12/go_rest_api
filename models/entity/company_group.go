package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CompanyGroup struct {
	gorm.Model
	ID        uuid.UUID  `gorm:"type:uuid;primary_key"`
	Name      string     `gorm:"type:varchar(100);not null" json:"name"`
	Status    *bool      `gorm:"not null;default:false"`
	CreatedBy uuid.UUID  `gorm:"type:uuid" json:"created_by"`
	UpdatedBy uuid.UUID  `gorm:"type:uuid" json:"updated_by"`
	CreatedAt *time.Time `gorm:"not null;default:now()"`
	UpdatedAt *time.Time `gorm:"not null;default:now()"`
}

func (company *CompanyGroup) BeforeCreate(tx *gorm.DB) (err error) {
	// UUID version 4
	company.ID = uuid.New()
	return
}
