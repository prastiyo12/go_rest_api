package core

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRolePermission struct {
	gorm.Model
	ID        uuid.UUID `gorm:"type:uuid;primary_key"`
	CompanyId uuid.UUID `gorm:"type:uuid" json:"company_id"`
	RoleId    uuid.UUID `gorm:"type:uuid" json:"role_id"`
	MenuId    uuid.UUID `gorm:"type:uuid" json:"menu_id"`
	OrderId   int       `gorm:"type:int" json:"order_id"`
	Status    bool      `gorm:"not null;default:false"`
	CreatedAt time.Time `gorm:"not null;default:now()"`
	UpdatedAt time.Time `gorm:"not null;default:now()"`
}

func (obj *UserRolePermission) BeforeCreate(tx *gorm.DB) (err error) {
	// UUID version 4
	obj.ID = uuid.New()
	return
}
