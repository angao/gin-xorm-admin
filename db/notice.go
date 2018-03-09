package db

import (
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
