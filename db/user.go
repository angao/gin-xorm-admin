package db

import (
	"errors"

	"github.com/angao/gin-xorm-admin/models"
	"github.com/angao/gin-xorm-admin/utils"
)

// UserDao operate user
type UserDao struct {
}

var usercols = []string{"id", "avatar", "account", "password", "salt", "name", "birthday", "sex", "email", "phone", "roleid", "deptid", "status", "createtime"}

// GetUser query user by account
func (UserDao) GetUser(account string) (*models.User, error) {
	user := new(models.User)
	has, err := x.Cols(usercols...).Where("account = ?", account).Get(user)
	if err != nil {
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
	has, err := x.Cols(usercols...).Where("id = ?", id).Get(user)
	if err != nil {
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
		return nil, err
	}
	if !has {
		return nil, errors.New("user not found")
	}
	return user, nil
}

// List query all user
func (UserDao) List(page models.Page) ([]models.User, error) {
	users := make([]models.User, 0)
	param := utils.StructToMap(page)
	err := x.SqlTemplateClient("user.all.sql", &param).Find(&users)
	if err != nil {
		return nil, err
	}
	return users, nil
}

// Save is save one user
func (UserDao) Save(user models.User) error {
	_, err := x.Insert(&user)
	return err
}

// Delete is delete a user
func (UserDao) Delete(id int64) error {
	user := new(models.User)
	_, err := x.Id(id).Delete(user)
	return err
}

// Update user
func (UserDao) Update(user *models.User) error {
	_, err := x.Id(user.ID).Cols(usercols...).Update(user)
	return err
}
