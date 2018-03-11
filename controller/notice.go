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
type NoticeController struct {
	NoticeDao db.NoticeDao
}

// Index notice home
func (NoticeController) Index(c *gin.Context) {
	r.HTML(c.Writer, http.StatusOK, "system/notice/notice.html", gin.H{})
}

// List query all notice
func (nc NoticeController) List(c *gin.Context) {
	page := forms.Page{}
	if err := c.Bind(&page); err != nil {
		r.JSON(c.Writer, http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	notices, err := nc.NoticeDao.List(page)
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
func (nc NoticeController) Add(c *gin.Context) {
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
		err := nc.NoticeDao.Save(notice)
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
func (nc NoticeController) ToEdit(c *gin.Context) {
	noticeID := c.Param("noticeId")
	id, err := strconv.ParseInt(noticeID, 10, 64)
	if err != nil {
		r.JSON(c.Writer, http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	notice, err := nc.NoticeDao.Get(id)
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
func (nc NoticeController) Edit(c *gin.Context) {
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
		err := nc.NoticeDao.Update(notice)
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
func (nc NoticeController) Delete(c *gin.Context) {
	noticeID := c.PostForm("noticeId")
	id, err := strconv.ParseInt(noticeID, 10, 64)
	if err != nil {
		r.JSON(c.Writer, http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	err = nc.NoticeDao.Delete(id)
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
