package core

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserType struct {
	gorm.Model
	ID        uuid.UUID `gorm:"type:uuid;primary_key"`
	CompanyId uuid.UUID `gorm:"type:uuid" json:"company_id"`
	Name      string    `gorm:"type:varchar(100);not null"`
	Status    bool      `gorm:"not null;default:false"`
	CreatedAt time.Time `gorm:"not null;default:now()"`
	UpdatedAt time.Time `gorm:"not null;default:now()"`
}

func (tipe *UserType) BeforeCreate(tx *gorm.DB) (err error) {
	// UUID version 4
	tipe.ID = uuid.New()
	return
}
