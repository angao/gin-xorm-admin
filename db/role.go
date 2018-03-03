package db

import (
	"fmt"

	"github.com/angao/gin-xorm-admin/forms"
	"github.com/angao/gin-xorm-admin/utils"
)

// RoleDao operate role
type RoleDao struct{}

// RoleBean for ztree
type RoleBean struct {
	ID       int    `xorm:"id" json:"id"`
	Pid      int    `xorm:"pId" json:"pId"`
	PName    string `xorm:"p_name" json:"pName"`
	Deptid   int    `xorm:"deptid" json:"deptId"`
	Tips     string `xorm:"tips" json:"tips"`
	DeptName string `xorm:"dept_name" json:"dept_name"`
	Name     string `json:"name"`
	Open     bool   `xorm:"open" json:"open"`
	Checked  bool   `json:"checked"`
}

// QueryAllRole query all role
func (RoleDao) QueryAllRole() ([]RoleBean, error) {
	var roles []RoleBean
	err := x.SqlMapClient("queryAllRole").Find(&roles)
	if err != nil {
		fmt.Printf("%#v\n", err)
		return nil, err
	}
	return roles, nil
}

// List query all role containes dept
func (RoleDao) List(roleForm forms.RoleForm) ([]RoleBean, error) {
	var roles []RoleBean
	param := utils.StructToMap(roleForm)
	err := x.SqlTemplateClient("role.list.sql", &param).Find(&roles)
	if err != nil {
		fmt.Printf("%#v\n", err)
		return nil, err
	}
	return roles, nil
}

// TreeListByRoleID query role by roleID
func (RoleDao) TreeListByRoleID(roleIds []string) ([]RoleBean, error) {
	var roles []RoleBean
	param := make(map[string]interface{})
	param["roleIds"] = roleIds
	param["length"] = len(roleIds) - 1
	err := x.SqlTemplateClient("treeListByRoleID.sql", &param).Find(&roles)
	if err != nil {
		fmt.Printf("%#v\n", err)
		return nil, err
	}
	return roles, nil
}
