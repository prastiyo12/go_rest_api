package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Pemilu struct {
	gorm.Model
	ID                 uuid.UUID  `gorm:"type:uuid;primary_key"`
	CompanyId          uuid.UUID  `gorm:"type:uuid"`
	TpsId              uuid.UUID  `gorm:"type:uuid" json:"tps_id"`
	TotalVoters        int        `gorm:"not null" json:"total_voters"`
	TotalVotersCompany int        `gorm:"not null" json:"total_voters_company"`
	TotalVotersGroup   int        `gorm:"not null" json:"total_voters_group"`
	TotalVotersOther   int        `gorm:"not null" json:"total_voters_other"`
	PathPhoto          string     `gorm:"type:varchar(255)" json:"path_photo"`
	Longitude          float64    `gorm:"type:float8" json:"longitude"`
	Latitude           float64    `gorm:"type:float8" json:"latitude"`
	Info1              string     `gorm:"type:varchar(255)" json:"info1"`
	Info2              string     `gorm:"type:varchar(255)" json:"info2"`
	Info3              string     `gorm:"type:varchar(255)" json:"info3"`
	Status             *bool      `gorm:"not null;default:false"`
	CreatedBy          uuid.UUID  `gorm:"type:uuid" json:"created_by"`
	UpdatedBy          uuid.UUID  `gorm:"type:uuid" json:"updated_by"`
	CreatedAt          *time.Time `gorm:"not null;default:now()"`
	UpdatedAt          *time.Time `gorm:"not null;default:now()"`
}

func (profile *Pemilu) BeforeCreate(tx *gorm.DB) (err error) {
	// UUID version 4
	profile.ID = uuid.New()
	return
}
