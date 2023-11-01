package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Campaign struct {
	gorm.Model
	ID             uuid.UUID  `gorm:"type:uuid;primary_key"`
	CompanyId      uuid.UUID  `gorm:"type:uuid"`
	ProfileId      uuid.UUID  `gorm:"type:uuid"`
	Campaign       string     `gorm:"type:varchar(255);not null" json:"campaign"`
	CampaignTarget int        `gorm:"type:int8;not null" json:"campaign_target"`
	CampaignAmount int        `gorm:"type:int8;not null" json:"campaign_amount"`
	TpsId          uuid.UUID  `gorm:"type:uuid" json:"tps_id"`
	CampaignStatus int        `gorm:"not null;default:1"`
	Status         *bool      `gorm:"not null;default:false"`
	CreatedBy      uuid.UUID  `gorm:"type:uuid" json:"created_by"`
	UpdatedBy      uuid.UUID  `gorm:"type:uuid" json:"updated_by"`
	CreatedAt      *time.Time `gorm:"not null;default:now()"`
	UpdatedAt      *time.Time `gorm:"not null;default:now()"`
}

func (obj *Campaign) BeforeCreate(tx *gorm.DB) (err error) {
	// UUID version 4
	obj.ID = uuid.New()
	return
}
