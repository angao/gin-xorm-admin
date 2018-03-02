package controller

import (
	"net/http"

	"github.com/angao/gin-xorm-admin/db"
	"github.com/angao/gin-xorm-admin/forms"
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
			"error": err,
		})
		return
	}
	menus, err := menuDao.List(menuForm)
	if err != nil {
		r.JSON(c.Writer, http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}
	r.JSON(c.Writer, http.StatusOK, menus)
}
