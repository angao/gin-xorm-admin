package db

import (
	"github.com/angao/gin-xorm-admin/forms"
	"github.com/angao/gin-xorm-admin/models"
	"github.com/angao/gin-xorm-admin/utils"
)

// MenuDao 菜单操作
type MenuDao struct{}

// GetMenuByRoleIds 根据角色查询所属菜单
func (MenuDao) GetMenuByRoleIds(roleIDs []string) ([]models.Menu, error) {
	var menus []models.Menu
	param := make(map[string]interface{})
	param["roleIds"] = roleIDs
	param["length"] = len(roleIDs) - 1
	err := x.SqlTemplateClient("getMenuByRoleIds.sql", &param).Find(&menus)
	if err != nil {
		return nil, err
	}
	return menus, nil
}

// List query menu
func (MenuDao) List(menuForm forms.MenuForm) ([]models.Menu, error) {
	var menus []models.Menu
	param := utils.StructToMap(menuForm)
	err := x.SqlTemplateClient("menu.list.sql", &param).Find(&menus)
	if err != nil {
		return nil, err
	}
	return menus, nil
}
