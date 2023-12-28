package SettingData

import (
	"time"

	"go_rest_api/database"

	"github.com/google/uuid"
)

type SettingDataRes struct {
	ID          uuid.UUID `json:"id"`
	CompanyId   uuid.UUID `json:"company_id"`
	Name        string    `json:"name"`
	Value       string    `json:"value"`
	Table       string    `json:"table"`
	Field       string    `json:"field"`
	Description string    `json:"description"`
	Status      bool      `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
}

func GetSettingValue(name string) (u SettingDataRes, error error) {
	qState := "SELECT * FROM setting_data WHERE name = '" + name + "'"
	err := database.DB.Raw(qState).Scan(&u).Error
	if err != nil {
		return u, err
	}

	return u, nil
}
