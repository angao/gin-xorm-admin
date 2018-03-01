package db

import "github.com/angao/gin-xorm-admin/models"

// DeptDao operate dept db
type DeptDao struct{}

// List query all dept
func (DeptDao) List() ([]models.Dept, error) {
	depts := make([]models.Dept, 0)
	err := x.Table("sys_dept").Find(&depts)
	if err != nil {
		return nil, err
	}
	return depts, nil
}
