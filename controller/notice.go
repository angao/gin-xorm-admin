package controller

import (
	"net/http"

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
