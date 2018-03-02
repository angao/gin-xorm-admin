package controller

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/angao/gin-xorm-admin/db"
	"github.com/angao/gin-xorm-admin/forms"
	"github.com/gin-gonic/gin"
)

// RoleController handle role request
type RoleController struct {
}

// Index handle /role
func (RoleController) Index(c *gin.Context) {
	r.HTML(c.Writer, http.StatusOK, "system/role/role.html", gin.H{})
}

// List query all role
func (RoleController) List(c *gin.Context) {
	var roleDao db.RoleDao
	var roleForm forms.RoleForm
	if err := c.Bind(&roleForm); err != nil {
		r.JSON(c.Writer, http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	roles, err := roleDao.List(roleForm)
	if err != nil {
		r.JSON(c.Writer, http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}
	r.JSON(c.Writer, http.StatusOK, gin.H{
		"data": roles,
	})
}

// ToAdd to add page
func (RoleController) ToAdd(c *gin.Context) {
	r.HTML(c.Writer, http.StatusOK, "system/role/role_add.html", "")
}

// RoleTreeListByUserID handle user role by userId
func (RoleController) RoleTreeListByUserID(c *gin.Context) {
	userID := c.Param("userId")
	if userID == "" {
		r.JSON(c.Writer, http.StatusBadRequest, gin.H{
			"error": "参数错误",
		})
		return
	}
	var userDao db.UserDao
	id, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		r.JSON(c.Writer, http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	user, err := userDao.GetUserByID(id)
	if err != nil {
		r.JSON(c.Writer, http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	var roleDao db.RoleDao
	if user.RoleId == "" {
		roles, err := roleDao.QueryAllRole()
		if err != nil {
			r.JSON(c.Writer, http.StatusBadRequest, gin.H{
				"error": err,
			})
			return
		}
		r.JSON(c.Writer, http.StatusOK, roles)
		return
	}
	roleIds := strings.Split(user.RoleId, ",")
	roles, err := roleDao.TreeListByRoleID(roleIds)
	if err != nil {
		r.JSON(c.Writer, http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	r.JSON(c.Writer, http.StatusOK, roles)
}
