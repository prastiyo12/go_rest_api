package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DapilArea struct {
	gorm.Model
	DapilId     uuid.UUID  `gorm:"type:uuid;not null" json:"dapil_id"`
	CityIdId    uuid.UUID  `gorm:"type:uuid;not null" json:"city_id"`
	DistrictId  uuid.UUID  `gorm:"type:uuid;not null" json:"district_id"`
	VillageId   uuid.UUID  `gorm:"type:uuid;not null" json:"village_id"`
	TotalVoters int        `gorm:"not null" json:"total_voters"`
	Status      *bool      `gorm:"not null;default:true"`
	CreatedBy   uuid.UUID  `gorm:"type:uuid" json:"created_by"`
	UpdatedBy   uuid.UUID  `gorm:"type:uuid" json:"updated_by"`
	CreatedAt   *time.Time `gorm:"not null;default:now()"`
	UpdatedAt   *time.Time `gorm:"not null;default:now()"`
}
