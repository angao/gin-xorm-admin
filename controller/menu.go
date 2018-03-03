package controller

import (
	"net/http"
	"strconv"

	"github.com/angao/gin-xorm-admin/db"
	"github.com/angao/gin-xorm-admin/forms"
	"github.com/angao/gin-xorm-admin/models"
	"github.com/gin-gonic/gin"
)

// MenuController handle menu request
type MenuController struct {
}

// Index handle /menu
func (MenuController) Index(c *gin.Context) {
	r.HTML(c.Writer, http.StatusOK, "system/menu/menu.html", gin.H{})
}

// List query all menu
func (MenuController) List(c *gin.Context) {
	var menuDao db.MenuDao
	var menuForm forms.MenuForm
	if err := c.Bind(&menuForm); err != nil {
		r.JSON(c.Writer, http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	menus, err := menuDao.List(menuForm)
	if err != nil {
		r.JSON(c.Writer, http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	r.JSON(c.Writer, http.StatusOK, menus)
}

// Remove delete menu
func (MenuController) Remove(c *gin.Context) {
	menuID := c.PostForm("menuId")
	if menuID == "" {
		r.JSON(c.Writer, http.StatusBadRequest, gin.H{
			"error": "参数错误",
		})
		return
	}
	var menuDao db.MenuDao
	id, err := strconv.ParseInt(menuID, 10, 64)
	if err != nil {
		r.JSON(c.Writer, http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	menu := models.Menu{
		Id:     id,
		Status: 0,
	}
	err = menuDao.Update(menu)
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
