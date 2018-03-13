package db

import (
	"errors"
	"strconv"
	"strings"

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
func (RoleDao) List(page models.Page) ([]models.Role, error) {
	var roles []models.Role
	param := utils.StructToMap(page)
	param["Id"] = 0
	err := x.SqlTemplateClient("role.list.sql", &param).Find(&roles)
	if err != nil {
		return nil, err
	}
	for _, role := range roles {
		if role.PName == "" {
			role.PName = "--"
		}
	}
	return roles, nil
}

// Get query one role
func (RoleDao) Get(id int64) (models.Role, error) {
	var role models.Role
	param := map[string]interface{}{
		"Id":     id,
		"Order":  "",
		"Offset": 0,
		"Limit":  0,
		"Name":   "",
	}
	has, err := x.SqlTemplateClient("role.list.sql", &param).Get(&role)
	if err != nil {
		return role, err
	}
	if !has {
		return role, errors.New("role not found")
	}
	if role.PName == "" {
		role.PName = "--"
	}
	return role, nil
}

// Save role
func (RoleDao) Save(role models.Role) error {
	_, err := x.Insert(&role)
	return err
}

// Update role
func (RoleDao) Update(role models.Role) error {
	_, err := x.Id(role.Id).Update(&role)
	return err
}

//Delete role
func (RoleDao) Delete(id int64) error {
	role := new(models.Role)
	_, err := x.Id(id).Delete(role)
	return err
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

// DeleteRolesByID delete roles by id
func (RoleDao) DeleteRolesByID(roleID int64) error {
	relation := new(models.Relation)
	_, err := x.Table("sys_relation").Where("roleid = ?", roleID).Delete(relation)
	return err
}

// SetAuthority set authority
func (RoleDao) SetAuthority(roleID int64, ids string) error {
	session := x.NewSession()
	defer session.Close()
	err := session.Begin()
	if err != nil {
		return err
	}
	relation := new(models.Relation)
	_, err = session.Where("roleid = ?", roleID).Delete(relation)
	if err != nil {
		session.Rollback()
		return err
	}
	menuIDs := strings.Split(ids, ",")
	for _, menuID := range menuIDs {
		id, err := strconv.ParseInt(menuID, 10, 64)
		if err != nil {
			session.Rollback()
			return err
		}
		relation := models.Relation{
			RoleID: roleID,
			MenuID: id,
		}
		_, err = session.Insert(&relation)
		if err != nil {
			session.Rollback()
			return err
		}
	}
	session.Commit()
	return nil
}

// GetURLByRoleID get url
func (RoleDao) GetURLByRoleID(id int64) ([]string, error) {
	permissions := make([]string, 0)
	err := x.SqlMapClient("queryURLByRoleID", id).Find(&permissions)
	return permissions, err
}
