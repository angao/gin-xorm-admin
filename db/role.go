package db

import (
	"github.com/angao/gin-xorm-admin/forms"
	"github.com/angao/gin-xorm-admin/models"
	"github.com/angao/gin-xorm-admin/utils"
)

// RoleDao operate role
type RoleDao struct{}

// QueryAllRole query all role
func (RoleDao) QueryAllRole() ([]models.ZTreeNode, error) {
	var nodes []models.ZTreeNode
	err := x.SqlMapClient("queryAllRole").Find(&nodes)
	if err != nil {
		return nil, err
	}
	return nodes, nil
}

// List query all role containes dept
func (RoleDao) List(roleForm forms.RoleForm) ([]models.Role, error) {
	var roles []models.Role
	param := utils.StructToMap(roleForm)
	err := x.SqlTemplateClient("role.list.sql", &param).Find(&roles)
	if err != nil {
		return nil, err
	}
	return roles, nil
}

// TreeListByRoleID query role by roleID
func (RoleDao) TreeListByRoleID(roleIds []string) ([]models.ZTreeNode, error) {
	var roles []models.ZTreeNode
	param := make(map[string]interface{}, 2)
	param["roleIds"] = roleIds
	param["length"] = len(roleIds) - 1
	err := x.SqlTemplateClient("role.treeByRoleId.sql", &param).Find(&roles)
	if err != nil {
		return nil, err
	}
	return roles, nil
}
