package controller

import (
	"log"
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

// ToAdd to add menu page
func (MenuController) ToAdd(c *gin.Context) {
	r.HTML(c.Writer, http.StatusOK, "system/menu/menu_add.html", gin.H{})
}

// SelectMenuTreeList query menu
func (MenuController) SelectMenuTreeList(c *gin.Context) {
	var menuDao db.MenuDao

	menus, err := menuDao.SelectMenuTreeList()
	if err != nil {
		r.JSON(c.Writer, http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	treeNode := models.ZTreeNode{
		ID:   0,
		Name: "顶级",
		Open: true,
		Pid:  0,
	}
	menus = append(menus, treeNode)
	r.JSON(c.Writer, http.StatusOK, menus)
}

// Edit update menu
func (MenuController) Edit(c *gin.Context) {
	menuID := c.Param("menuId")
	if menuID == "" {
		r.JSON(c.Writer, http.StatusBadRequest, gin.H{
			"message": "参数错误",
		})
		return
	}
	id, err := strconv.ParseInt(menuID, 10, 64)
	if err != nil {
		r.JSON(c.Writer, http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	var menuDao db.MenuDao
	menu, err := menuDao.Get(id)
	if err != nil {
		r.JSON(c.Writer, http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	if menu.Pcode != "0" {
		pMenu, err := menuDao.GetByPcode(menu.Pcode)
		if err != nil {
			r.JSON(c.Writer, http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}
		log.Printf("%#v\n", pMenu)
		menu.PcodeName = pMenu.Name
	}

	r.HTML(c.Writer, http.StatusOK, "system/menu/menu_edit.html", gin.H{
		"menu": menu,
	})
}
