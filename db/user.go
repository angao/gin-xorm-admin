package db

import (
	"errors"
	"log"
	"time"

	"github.com/angao/gin-xorm-admin/forms"
	"github.com/angao/gin-xorm-admin/models"
	"github.com/angao/gin-xorm-admin/utils"
)

// UserDao operate user
type UserDao struct {
}

type UserBean struct {
	Id         int64
	Avatar     string
	Account    string
	Name       string
	Birthday   time.Time
	Sex        string
	Email      string
	Phone      string
	RoleId     string `xorm:"roleid"`
	RoleName   string
	DeptId     int `xorm:"deptid"`
	DeptName   string
	Status     string
	CreateTime string `xorm:"'createtime'"`
}

// GetUser query user by account
func (UserDao) GetUser(account string) (*models.User, error) {
	user := new(models.User)
	has, err := x.Table("sys_user").Where("account = ?", account).Get(user)
	if err != nil {
		log.Printf("error: %#v\n", err)
		return nil, err
	}
	if !has {
		return nil, errors.New("user not found")
	}
	return user, nil
}

// GetUserRole query user by primary key
func (UserDao) GetUserRole(id int64) (*models.UserRole, error) {
	user := new(models.UserRole)
	has, err := x.Table("sys_user").Join("INNER", "sys_role", "sys_user.roleid = sys_role.id").Where("sys_user.id = ?", id).Get(user)
	if err != nil {
		log.Printf("error: %#v\n", err)
		return nil, err
	}
	if !has {
		return nil, errors.New("user not found")
	}
	return user, nil
}

// List query all user
func (UserDao) List(userForm forms.UserForm) ([]UserBean, error) {
	users := make([]UserBean, 0)
	param := utils.StructToMap(userForm)
	err := x.SqlTemplateClient("user.all.sql", &param).Find(&users)
	if err != nil {
		log.Printf("error: %#v\n", err)
		return nil, err
	}
	return users, nil
}
