package db

import (
	"errors"

	"github.com/angao/gin-xorm-admin/models"
	"github.com/angao/gin-xorm-admin/utils"
)

// NoticeDao 通知操作
type NoticeDao struct{}

var noticeCols = []string{"id", "type", "content", "title", "createtime", "creater"}

// List query all notice
func (NoticeDao) List(page models.Page) ([]models.Notice, error) {
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
	var notice models.Notice
	has, err := x.Where("id = ?", id).Cols(noticeCols...).Get(&notice)
	if err != nil {
		return notice, err
	}
	if !has {
		return notice, errors.New("notice not found")
	}
	return notice, nil
}

// Save a notice
func (NoticeDao) Save(notice models.Notice) error {
	_, err := x.Insert(&notice)
	return err
}

// Update a notice
func (NoticeDao) Update(notice models.Notice) error {
	_, err := x.Id(notice.ID).Cols(noticeCols...).Update(&notice)
	return err
}

// Delete a notice
func (NoticeDao) Delete(id int64) error {
	notice := new(models.Notice)
	_, err := x.Id(id).Delete(notice)
	return err
}
