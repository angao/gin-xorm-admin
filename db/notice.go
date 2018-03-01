package db

import (
	"github.com/angao/gin-xorm-admin/models"
)

// NoticeDao 通知操作
type NoticeDao struct{}

// List query all notice
func (NoticeDao) List() ([]models.Notice, error) {
	var notices []models.Notice
	err := x.Table("sys_notice").Find(&notices)
	if err != nil {
		return nil, err
	}
	return notices, nil
}
