package controller

import (
	"net/http"
	"strconv"

	"github.com/angao/gin-xorm-admin/models"
	"github.com/gin-contrib/sessions"

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

// Add save a notice
func (NoticeController) Add(c *gin.Context) {
	notice := models.Notice{}
	if err := c.Bind(&notice); err != nil {
		r.JSON(c.Writer, http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	session := sessions.Default(c)
	id, ok := session.Get("user_id").(int64)
	if ok {
		notice.Creater = id
		var noticeDao db.NoticeDao
		err := noticeDao.Save(notice)
		if err != nil {
			r.JSON(c.Writer, http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}
		r.JSON(c.Writer, http.StatusOK, gin.H{
			"message": "success",
		})
		return
	}
	c.Redirect(http.StatusMovedPermanently, "/login")
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

// Edit update notice
func (NoticeController) Edit(c *gin.Context) {
	notice := models.Notice{}
	if err := c.Bind(&notice); err != nil {
		r.JSON(c.Writer, http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	session := sessions.Default(c)
	id, ok := session.Get("user_id").(int64)
	if ok {
		notice.Creater = id
		var noticeDao db.NoticeDao
		err := noticeDao.Update(notice)
		if err != nil {
			r.JSON(c.Writer, http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}
		r.JSON(c.Writer, http.StatusOK, gin.H{
			"message": "success",
		})
		return
	}
	c.Redirect(http.StatusMovedPermanently, "/login")
}

// Delete a notice
func (NoticeController) Delete(c *gin.Context) {
	noticeID := c.PostForm("noticeId")
	id, err := strconv.ParseInt(noticeID, 10, 64)
	if err != nil {
		r.JSON(c.Writer, http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	var noticeDao db.NoticeDao
	err = noticeDao.Delete(id)
	if err != nil {
		r.JSON(c.Writer, http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	r.JSON(c.Writer, http.StatusOK, gin.H{
		"message": "success",
	})
}
