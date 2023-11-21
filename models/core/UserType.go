package core

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserType struct {
	gorm.Model
	ID        uuid.UUID `gorm:"type:uuid;primary_key" json:"id"`
	CompanyId uuid.UUID `gorm:"type:uuid" json:"company_id"`
	Name      string    `gorm:"type:varchar(100);not null" json:"name"`
	Status    bool      `gorm:"not null;default:false" json:"status"`
	CreatedAt time.Time `gorm:"not null;default:now()" json:"created_at"`
	UpdatedAt time.Time `gorm:"not null;default:now()" json:"updated_at"`
}

func (tipe *UserType) BeforeCreate(tx *gorm.DB) (err error) {
	// UUID version 4
	tipe.ID = uuid.New()
	return
}
