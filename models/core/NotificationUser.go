package core

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type NotificationUser struct {
	gorm.Model
	ID        uuid.UUID  `gorm:"type:uuid;primary_key"`
	CompanyId uuid.UUID  `gorm:"type:uuid" json:"company_id"`
	ProfileId uuid.UUID  `gorm:"type:uuid" json:"profile_id"`
	Status    *bool      `gorm:"not null;default:false"`
	CreatedBy uuid.UUID  `gorm:"type:uuid" json:"created_by"`
	UpdatedBy uuid.UUID  `gorm:"type:uuid" json:"updated_by"`
	CreatedAt *time.Time `gorm:"not null;default:now()"`
	UpdatedAt *time.Time `gorm:"not null;default:now()"`
}

func (obj *NotificationUser) BeforeCreate(tx *gorm.DB) (err error) {
	// UUID version 4
	obj.ID = uuid.New()
	return
}
