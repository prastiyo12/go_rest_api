package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Notification struct {
	gorm.Model
	ID           uuid.UUID  `gorm:"type:uuid;primary_key"`
	CompanyId    uuid.UUID  `gorm:"type:uuid" json:"company_id"`
	Notification string     `gorm:"type:varchar(100);not null" json:"name"`
	StartDate    time.Time  `gorm:"not null;default:now()" json:"start_date"`
	EndDate      time.Time  `gorm:"not null;default:now()" json:"end_date"`
	PathPhoto    string     `gorm:"type:varchar(255)" json:"path_photo"`
	Status       *bool      `gorm:"not null;default:false"`
	CreatedBy    uuid.UUID  `gorm:"type:uuid" json:"created_by"`
	UpdatedBy    uuid.UUID  `gorm:"type:uuid" json:"updated_by"`
	CreatedAt    *time.Time `gorm:"not null;default:now()"`
	UpdatedAt    *time.Time `gorm:"not null;default:now()"`
}

func (obj *Notification) BeforeCreate(tx *gorm.DB) (err error) {
	// UUID version 4
	obj.ID = uuid.New()
	return
}
