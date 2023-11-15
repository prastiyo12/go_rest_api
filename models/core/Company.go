package core

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Company struct {
	gorm.Model
	ID        uuid.UUID `gorm:"type:uuid;primary_key"`
	GroupId   uuid.UUID `gorm:"type:uuid" json:"group_id"`
	Code      int       `gorm:"type:int8;not null" json:"code"`
	Name      string    `gorm:"type:varchar(100);not null" json:"name"`
	PathPhoto string    `gorm:"type:varchar(255)" json:"path_photo"`
	DapilId   uuid.UUID `gorm:"type:uuid" json:"dapil_id"`
	Status    bool      `gorm:"not null;default:false"`
	CreatedBy uuid.UUID `gorm:"type:uuid" json:"created_by"`
	UpdatedBy uuid.UUID `gorm:"type:uuid" json:"updated_by"`
	CreatedAt time.Time `gorm:"not null;default:now()"`
	UpdatedAt time.Time `gorm:"not null;default:now()"`
}

func (company *Company) BeforeCreate(tx *gorm.DB) (err error) {
	// UUID version 4
	company.ID = uuid.New()
	return
}
