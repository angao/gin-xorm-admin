package db

import (
	"errors"

	"github.com/angao/gin-xorm-admin/forms"
	"github.com/angao/gin-xorm-admin/models"
	"github.com/angao/gin-xorm-admin/utils"
)

// NoticeDao 通知操作
type NoticeDao struct{}

// List query all notice
func (NoticeDao) List(page forms.Page) ([]models.Notice, error) {
	notices := make([]models.Notice, 0)
	param := utils.StructToMap(page)
	err := x.SqlTemplateClient("notice.list.sql", &param).Find(&notices)
	if err != nil {
		return nil, err
	}
	return notices, nil
}

// Get get one notice
func (NoticeDao) Get(id int64) (models.Notice, error) {
	var cols = []string{"id", "type", "content", "title", "createtime", "creater"}
	var notice models.Notice
	has, err := x.Where("id = ?", id).Cols(cols...).Get(&notice)
	if err != nil {
		return notice, err
	}
	if !has {
		return notice, errors.New("notice not found")
	}
	return notice, nil
}
