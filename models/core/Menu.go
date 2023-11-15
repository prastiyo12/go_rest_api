package core

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Menu struct {
	gorm.Model
	ID        uuid.UUID `gorm:"type:uuid;primary_key"`
	CompanyId uuid.UUID `gorm:"type:uuid" json:"company_id"`
	ParentId  uuid.UUID `gorm:"type:uuid" json:"parent_id"`
	Name      string    `gorm:"type:varchar(100);not null"`
	Path      string    `gorm:"type:varchar(100);not null"`
	Icon      string    `gorm:"type:varchar(100);not null"`
	Translate string    `gorm:"type:varchar(100);not null"`
	Status    bool      `gorm:"not null;default:false"`
	CreatedAt time.Time `gorm:"not null;default:now()"`
	UpdatedAt time.Time `gorm:"not null;default:now()"`
}

func (model *Menu) BeforeCreate(tx *gorm.DB) (err error) {
	// UUID version 4
	model.ID = uuid.New()
	return
}
