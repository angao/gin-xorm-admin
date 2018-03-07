package db

import (
	"errors"

	"github.com/angao/gin-xorm-admin/models"
)

// DeptDao operate dept db
type DeptDao struct{}

// List query all dept
func (DeptDao) List(name string) ([]models.Dept, error) {
	var depts []models.Dept
	param := make(map[string]interface{})
	param["name"] = name
	err := x.SqlTemplateClient("dept.all.sql", &param).Find(&depts)
	if err != nil {
		return nil, err
	}
	return depts, nil
}

// Get query one by primary key
func (DeptDao) Get(id int64) (models.Dept, error) {
	var dept models.Dept
	has, err := x.SqlMapClient("queryOneDept", id).Get(&dept)
	if err != nil {
		return dept, err
	}
	if !has {
		return dept, errors.New("dept not found")
	}
	return dept, nil
}
