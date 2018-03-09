package controller

import (
	"net/http"
	"strconv"

	"github.com/angao/gin-xorm-admin/db"
	"github.com/angao/gin-xorm-admin/forms"

	"github.com/gin-gonic/gin"
)

// NoticeController 通知
type NoticeController struct{}

// Index notice home
func (NoticeController) Index(c *gin.Context) {
	r.HTML(c.Writer, http.StatusOK, "system/notice/notice.html", gin.H{})
}

// List query all notice
func (NoticeController) List(c *gin.Context) {
	page := forms.Page{}
	if err := c.Bind(&page); err != nil {
		r.JSON(c.Writer, http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	var noticeDao db.NoticeDao
	notices, err := noticeDao.List(page)
	if err != nil {
		r.JSON(c.Writer, http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	r.JSON(c.Writer, http.StatusOK, notices)
}

// ToAdd to add page
func (NoticeController) ToAdd(c *gin.Context) {
	r.HTML(c.Writer, http.StatusOK, "system/notice/notice_add.html", gin.H{})
}

// ToEdit to edit page
func (NoticeController) ToEdit(c *gin.Context) {
	noticeID := c.Param("noticeId")
	id, err := strconv.ParseInt(noticeID, 10, 64)
	if err != nil {
		r.JSON(c.Writer, http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	var noticeDao db.NoticeDao
	notice, err := noticeDao.Get(id)
	if err != nil {
		r.JSON(c.Writer, http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	r.HTML(c.Writer, http.StatusOK, "system/notice/notice_edit.html", gin.H{
		"notice": notice,
	})
}
