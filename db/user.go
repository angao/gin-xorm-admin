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

// UserBean for form
type UserBean struct {
	ID         int64 `json:"Id"`
	Avatar     string
	Account    string
	Name       string
	Birthday   time.Time
	Sex        string
	Email      string
	Phone      string
	RoleID     string `json:"RoleId" xorm:"roleid"`
	RoleName   string
	DeptID     int `json:"DeptId" xorm:"deptid"`
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

// GetUserByID query user by id
func (UserDao) GetUserByID(id int64) (*models.User, error) {
	user := new(models.User)
	has, err := x.Table("sys_user").Where("id = ?", id).Get(user)
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

// Save is save one user
func (UserDao) Save(user models.User) error {
	_, err := x.Table("sys_user").Insert(&user)
	return err
}

// Delete is delete a user
func (UserDao) Delete(id int64) error {
	user := new(models.User)
	_, err := x.Table("sys_user").Id(id).Delete(user)
	return err
}

// Update user
func (UserDao) Update(user *models.User) error {
	_, err := x.Table("sys_user").Id(user.Id).Cols("status").Update(user)
	return err
}
