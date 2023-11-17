package vote

import (
	"time"

	"github.com/google/uuid"
)

type DapilArea struct {
	DapilId     uuid.UUID `gorm:"type:uuid;not null" json:"dapil_id"`
	ProvinceId  string    `gorm:"varchar(100)" json:"province_id"`
	CityId      string    `gorm:"varchar(100)" json:"city_id"`
	DistrictId  string    `gorm:"varchar(100)" json:"district_id"`
	VillageId   string    `gorm:"varchar(100)" json:"village_id"`
	TotalVoters int       `gorm:"not null" json:"total_voters"`
	Status      bool      `gorm:"not null;default:true"`
	CreatedBy   uuid.UUID `gorm:"type:uuid" json:"created_by"`
	UpdatedBy   uuid.UUID `gorm:"type:uuid" json:"updated_by"`
	CreatedAt   time.Time `gorm:"not null;default:now()"`
	UpdatedAt   time.Time `gorm:"not null;default:now()"`
}
