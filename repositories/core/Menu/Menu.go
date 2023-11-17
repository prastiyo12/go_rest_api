package Menu

import (
	"strings"
	"time"

	"go_rest_api/database"
	"go_rest_api/models/core"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type MenuRequest struct {
	ID        uuid.UUID `json:"id"`
	CompanyId uuid.UUID `json:"company_id"`
	ParentId  uuid.UUID `json:"parent_id"`
	Name      string    `json:"name"`
	Path      string    `json:"path"`
	Icon      string    `json:"icon"`
	Translate string    `json:"translate"`
	Status    bool      `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type MenuUpdateRequest struct {
	CompanyId uuid.UUID `json:"company_id"`
	ParentId  uuid.UUID `json:"parent_id"`
	Name      string    `json:"name"`
	Path      string    `json:"path"`
	Icon      string    `json:"icon"`
	Translate string    `json:"translate"`
	Status    bool      `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ResMenu struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Path  string    `json:"path"`
	Icon  string    `json:"icon"`
	Items ResMenus  `json:"items"`
}

type ResMenus []*ResMenu

func GetAll(c *fiber.Ctx) (u []*core.Menu, error error) {
	keyword := c.Query("q")

	qState := "SELECT * FROM menus "
	if keyword != "" {
		qState = qState + "WHERE LOWER(name) LIKE '%" + strings.ToLower(keyword) + "%'"
	}
	qState = qState + " ORDER BY created_at DESC"

	err := database.DB.Raw(qState).Scan(&u).Error
	if err != nil {
		return u, err
	}

	return u, nil
}

func GetDataByID(id string) (c *core.Menu, err error) {
	err = database.DB.Table("menus").Where("id = ?", id).Find(&c).Error
	if err != nil {
		return &core.Menu{}, err
	}
	return c, nil
}

func (c *MenuRequest) Create() (*MenuRequest, error) {
	var err = database.DB.Table("menus").Create(&c).Error
	if err != nil {
		return &MenuRequest{}, err
	}
	return c, nil
}

func (u *MenuUpdateRequest) Update(id string) (*MenuUpdateRequest, error) {
	var err = database.DB.Table("menus").Where("id = ?", id).Updates(&u).Error
	if err != nil {
		return u, err
	}
	return u, nil
}

func Delete(id string) error {
	var err = database.DB.Table("menus").Where("id = ?", id).Delete(core.Menu{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *MenuUpdateRequest) UpdateStatus(id string) (*MenuUpdateRequest, error) {
	var err = database.DB.Table("menus").Where("id = ?", id).Update("status", u.Status).Error
	if err != nil {
		return &MenuUpdateRequest{}, err
	}
	return u, nil
}

func GetAllMenu(c *fiber.Ctx, companyID, roleID string) (u ResMenus, error error) {
	var res []*core.Menu
	qState := "SELECT  m.id,m.name, m.path, m.icon,m.parent_id FROM user_role_permissions ur left join menus m on ur.menu_id = m.id "
	qState = qState + " where ur.company_id = '" + companyID + "'"
	qState = qState + " and ur.role_id  = '" + roleID + "'"
	qState = qState + " ORDER BY m.parent_id ASC, ur.order_id ASC"

	err := database.DB.Raw(qState).Scan(&res).Error
	if err != nil {
		return u, err
	}

	var result ResMenus
	for _, val := range res {
		res := &ResMenu{
			ID:   val.ID,
			Name: val.Name,
			Path: val.Path,
			Icon: val.Icon,
		}

		var found bool

		// iterate trough root nodes
		for _, root := range result {
			parent := findById(root, val.ParentId)
			if parent != nil {
				parent.Items = append(parent.Items, res)

				found = true
				break
			}
		}

		if !found {
			result = append(result, res)
		}
	}

	// out, err := json.Marshal(result)
	// if err != nil {
	// 	panic(err)
	// }

	return result, nil
}

func findById(root *ResMenu, id uuid.UUID) *ResMenu {
	queue := make([]*ResMenu, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		nextUp := queue[0]
		queue = queue[1:]
		if nextUp.ID == id {
			return nextUp
		}
		if len(nextUp.Items) > 0 {
			for _, child := range nextUp.Items {
				queue = append(queue, child)
			}
		}
	}
	return nil
}
