package core

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SettingData struct {
	gorm.Model
	Name        string     `gorm:"type:varchar(100);not null;primary_key"`
	Value       string     `gorm:"type:varchar(100);" json:"value"`
	Table       string     `gorm:"type:varchar(100);" json:"table"`
	Field       string     `gorm:"type:varchar(100);" json:"field"`
	Description string     `gorm:"type:text;" json:"description"`
	CompanyId   uuid.UUID  `gorm:"type:uuid" json:"company_id"`
	CreatedAt   *time.Time `gorm:"not null;default:now()"`
	UpdatedAt   *time.Time `gorm:"not null;default:now()"`
}
