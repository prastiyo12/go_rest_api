package Option

import (
	"go_rest_api/database"
	"go_rest_api/models/core"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func GetProvince(c *fiber.Ctx) (u []*core.AddressProvince, error error) {

	qState := "SELECT * FROM address_provinces "
	qState = qState + " ORDER BY province ASC"

	err := database.DB.Raw(qState).Scan(&u).Error
	if err != nil {
		return u, err
	}

	return u, nil
}

func GetCity(c *fiber.Ctx) (u []*core.AddressCity, error error) {
	paramsID := c.Query("province_id")
	search := c.Query("filter")
	qState := "SELECT * FROM address_cities WHERE city is not null"
	if paramsID != "" {
		qState = qState + " AND province_id = '" + paramsID + "'"
	}

	if search != "" {
		qState = qState + " AND lower(city) = '" + strings.ToLower(search) + "'"
	}
	qState = qState + " ORDER BY city ASC"
	err := database.DB.Raw(qState).Scan(&u).Error
	if err != nil {
		return u, err
	}

	return u, nil
}

func GetDistrict(c *fiber.Ctx) (u []*core.AddressDistrict, error error) {
	paramsID := c.Query("city_id")
	search := c.Query("filter")
	qState := "SELECT * FROM address_districts "
	if paramsID != "" {
		qState = qState + "WHERE city_id = '" + paramsID + "'"
	}
	if search != "" {
		qState = qState + " AND lower(district) = '" + strings.ToLower(search) + "'"
	}
	qState = qState + " ORDER BY district ASC"

	err := database.DB.Raw(qState).Scan(&u).Error
	if err != nil {
		return u, err
	}

	return u, nil
}

func GetVillage(c *fiber.Ctx) (u []*core.AddressVillage, error error) {
	paramsID := c.Query("district_id")
	search := c.Query("filter")
	qState := "SELECT * FROM address_villages "
	if paramsID != "" {
		qState = qState + "WHERE district_id = '" + paramsID + "'"
	}
	if search != "" {
		qState = qState + " AND lower(urban_village) = '" + strings.ToLower(search) + "'"
	}
	qState = qState + " ORDER BY urban_village ASC"

	err := database.DB.Raw(qState).Scan(&u).Error
	if err != nil {
		return u, err
	}

	return u, nil
}
