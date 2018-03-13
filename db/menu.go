package db

import (
	"errors"

	"github.com/angao/gin-xorm-admin/models"
	"github.com/angao/gin-xorm-admin/utils"
)

// menu column
var cols = []string{"id", "code", "pcode", "pcodes", "name", "icon", "url", "num", "levels", "ismenu", "tips", "status", "isopen"}

// MenuDao 菜单操作
type MenuDao struct{}

// Save a menu
func (MenuDao) Save(menu models.Menu) error {
	_, err := x.Table("sys_menu").Cols(cols...).Insert(&menu)
	return err
}

// GetMenuByRoleIds 根据角色查询所属菜单
func (MenuDao) GetMenuByRoleIds(roleIDs []string) ([]models.Menu, error) {
	var menus []models.Menu
	param := make(map[string]interface{})
	param["roleIds"] = roleIDs
	param["length"] = len(roleIDs) - 1
	err := x.SqlTemplateClient("menu.getByRoleIds.sql", &param).Find(&menus)
	if err != nil {
		return nil, err
	}
	return menus, nil
}

// List query menu
func (MenuDao) List(page models.Page) ([]models.Menu, error) {
	var menus []models.Menu
	param := utils.StructToMap(page)
	err := x.SqlTemplateClient("menu.list.sql", &param).Find(&menus)
	if err != nil {
		return nil, err
	}
	return menus, nil
}

// Update update menu
func (MenuDao) Update(menu models.Menu) error {
	_, err := x.Id(menu.Id).Cols(cols...).Update(&menu)
	return err
}

// Delete menu
func (MenuDao) Delete(id int64) error {
	menu := new(models.Menu)
	_, err := x.Id(id).Delete(menu)
	return err
}

// Get query one record by id
func (MenuDao) Get(id int64) (*models.Menu, error) {
	menu := new(models.Menu)
	menu.Id = id
	has, err := x.Cols(cols...).Get(menu)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, errors.New("menu has not found")
	}
	return menu, nil
}

// GetByPcode get menu by pcode
func (MenuDao) GetByPcode(pcode string) (*models.Menu, error) {
	menu := new(models.Menu)
	has, err := x.Cols(cols...).Where("code = ?", pcode).Get(menu)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, errors.New("menu has not found")
	}
	return menu, nil
}

// SelectMenuTreeList query menu
func (MenuDao) SelectMenuTreeList() ([]models.ZTreeNode, error) {
	var menus []models.ZTreeNode
	err := x.SqlTemplateClient("menu.tree.sql").Find(&menus)
	if err != nil {
		return nil, err
	}
	return menus, nil
}

// GetMenuIdsByRoleID query menuIDs by roleID
func (MenuDao) GetMenuIdsByRoleID(roleID int64) ([]int64, error) {
	menuIDs := make([]int64, 0)
	err := x.Table("sys_relation").Cols("menuid").Where("roleid = ?", roleID).Find(&menuIDs)
	if err != nil {
		return menuIDs, err
	}
	return menuIDs, nil
}

// GetMenusByMenuIDs query menu tree by menuIds
func (MenuDao) GetMenusByMenuIDs(menuIDs []int64) ([]models.ZTreeNode, error) {
	var menus []models.ZTreeNode
	param := make(map[string]interface{}, 2)
	param["menuIDs"] = menuIDs
	param["length"] = len(menuIDs) - 1
	err := x.SqlTemplateClient("menu.treeByMenuIds.sql", &param).Find(&menus)
	if err != nil {
		return nil, err
	}
	return menus, nil
}

// GetAllURL query all menu url
func (MenuDao) GetAllURL() ([]string, error) {
	url := make([]string, 0)
	err := x.Cols("url").Find(&url)
	return url, err
}
