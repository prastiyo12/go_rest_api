package core

import (
	"time"
)

// Address Provinces
type AddressProvince struct {
	ID        string    `gorm:"type:varchar(100);primary_key" json:"id"`
	CountryId string    `gorm:"type:varchar(100);not null" json:"country_id"`
	Province  string    `gorm:"type:varchar(255);not null" json:"province"`
	CreatedAt time.Time `gorm:"not null;default:now()"`
	UpdatedAt time.Time `gorm:"not null;default:now()"`
}

// Address Cities
type AddressCity struct {
	ID         string    `gorm:"type:varchar(100);primary_key" json:"id"`
	ProvinceId string    `gorm:"type:varchar(100);not null" json:"province_id"`
	City       string    `gorm:"type:varchar(255);not null" json:"city"`
	CreatedAt  time.Time `gorm:"not null;default:now()"`
	UpdatedAt  time.Time `gorm:"not null;default:now()"`
}

// Address Districts
type AddressDistrict struct {
	ID        string    `gorm:"type:varchar(100);primary_key" json:"id"`
	CityId    string    `gorm:"type:varchar(100);not null" json:"city_id"`
	District  string    `gorm:"type:varchar(255);not null" json:"district"`
	CreatedAt time.Time `gorm:"not null;default:now()"`
	UpdatedAt time.Time `gorm:"not null;default:now()"`
}

// Address Villages
type AddressVillage struct {
	ID           string    `gorm:"type:varchar(100);primary_key" json:"id"`
	DistrictId   string    `gorm:"type:varchar(100);not null" json:"district_id"`
	UrbanVillage string    `gorm:"type:varchar(255);not null" json:"urban_village"`
	CreatedAt    time.Time `gorm:"not null;default:now()"`
	UpdatedAt    time.Time `gorm:"not null;default:now()"`
}
