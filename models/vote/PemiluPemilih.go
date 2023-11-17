package vote

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Pemilih struct {
	gorm.Model
	ID             uuid.UUID  `gorm:"type:uuid;primary_key"`
	CompanyId      uuid.UUID  `gorm:"type:uuid"`
	TpsId          uuid.UUID  `gorm:"type:uuid" json:"tps_id"`
	Name           string     `gorm:"type:varchar(100);not null" json:"name"`
	Phone          string     `gorm:"type:varchar(100);not null" json:"phone"`
	IdentityNumber string     `gorm:"type:varchar(20)" json:"identity_number"`
	BirthDate      time.Time  `gorm:"not null" json:"birth_date"`
	BirthPlace     string     `gorm:"type:varchar(255)" json:"birth_place"`
	Gender         string     `gorm:"type:varchar(255)" json:"gender"`
	Address        string     `gorm:"type:varchar(255)" json:"address"`
	ProvinceId     string     `gorm:"type:varchar(100)" json:"province_id"`
	CityId         string     `gorm:"type:varchar(100)" json:"city_id"`
	DistrictId     string     `gorm:"type:varchar(100)" json:"district_id"`
	VillageId      string     `gorm:"type:varchar(100)" json:"village_id"`
	Rt             string     `gorm:"type:varchar(10)" json:"rt"`
	Rw             string     `gorm:"type:varchar(10)" json:"rw"`
	Longitude      float64    `gorm:"type:float8" json:"longitude"`
	Latitude       float64    `gorm:"type:float8" json:"latitude"`
	UserTypeId     uuid.UUID  `gorm:"type:uuid" json:"user_type_id"`
	UserId         uuid.UUID  `gorm:"type:uuid;not null" json:"user_id"`
	PathPhoto      string     `gorm:"type:varchar(255)" json:"path_photo"`
	Info1          string     `gorm:"type:varchar(255)" json:"info1"`
	Info2          string     `gorm:"type:varchar(255)" json:"info2"`
	Info3          string     `gorm:"type:varchar(255)" json:"info3"`
	Status         *bool      `gorm:"not null;default:false"`
	CreatedBy      uuid.UUID  `gorm:"type:uuid" json:"created_by"`
	UpdatedBy      uuid.UUID  `gorm:"type:uuid" json:"updated_by"`
	CreatedAt      *time.Time `gorm:"not null;default:now()"`
	UpdatedAt      *time.Time `gorm:"not null;default:now()"`
}

func (profile *Pemilih) BeforeCreate(tx *gorm.DB) (err error) {
	// UUID version 4
	profile.ID = uuid.New()
	return
}
