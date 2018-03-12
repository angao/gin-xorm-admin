package controller

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/angao/gin-xorm-admin/db"
	"github.com/angao/gin-xorm-admin/models"
	"github.com/gin-gonic/gin"
)

// RoleController handle role request
type RoleController struct {
	RoleDao db.RoleDao
	UserDao db.UserDao
}

// Index handle /role
func (RoleController) Index(c *gin.Context) {
	r.HTML(c.Writer, http.StatusOK, "system/role/role.html", gin.H{})
}

// List query all role
func (rc RoleController) List(c *gin.Context) {
	var page models.Page
	if err := c.Bind(&page); err != nil {
		r.JSON(c.Writer, http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	roles, err := rc.RoleDao.List(page)
	if err != nil {
		r.JSON(c.Writer, http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	r.JSON(c.Writer, http.StatusOK, roles)
}

// ToAdd to add page
func (RoleController) ToAdd(c *gin.Context) {
	r.HTML(c.Writer, http.StatusOK, "system/role/role_add.html", "")
}

// Add save role
func (rc RoleController) Add(c *gin.Context) {
	var role models.Role
	if err := c.Bind(&role); err != nil {
		r.JSON(c.Writer, http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	err := rc.RoleDao.Save(role)
	if err != nil {
		r.JSON(c.Writer, http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	r.JSON(c.Writer, http.StatusOK, gin.H{
		"message": "success",
	})
}

// ToEdit to edit page
func (rc RoleController) ToEdit(c *gin.Context) {
	roleID := c.Param("roleId")
	id, err := strconv.ParseInt(roleID, 10, 64)
	if err != nil {
		r.JSON(c.Writer, http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	role, err := rc.RoleDao.Get(id)
	if err != nil {
		r.JSON(c.Writer, http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	r.HTML(c.Writer, http.StatusOK, "system/role/role_edit.html", gin.H{
		"role": role,
	})
}

// Edit save role
func (rc RoleController) Edit(c *gin.Context) {
	var role models.Role
	if err := c.Bind(&role); err != nil {
		r.JSON(c.Writer, http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	err := rc.RoleDao.Update(role)
	if err != nil {
		r.JSON(c.Writer, http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	r.JSON(c.Writer, http.StatusOK, gin.H{
		"message": "success",
	})
}

// Remove delete role
func (rc RoleController) Remove(c *gin.Context) {
	roleID := c.PostForm("roleId")
	id, err := strconv.ParseInt(roleID, 10, 64)
	if err != nil {
		r.JSON(c.Writer, http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	err = rc.RoleDao.Delete(id)
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

// ToAssign to assign page
func (rc RoleController) ToAssign(c *gin.Context) {
	roleID := c.Param("roleId")
	id, err := strconv.ParseInt(roleID, 10, 64)
	if err != nil {
		r.JSON(c.Writer, http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	role, err := rc.RoleDao.Get(id)
	if err != nil {
		r.JSON(c.Writer, http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	r.HTML(c.Writer, http.StatusOK, "system/role/role_assign.html", gin.H{
		"role": role,
	})
}

// RoleTreeListByUserID handle user role by userId
func (rc RoleController) RoleTreeListByUserID(c *gin.Context) {
	userID := c.Param("userId")
	id, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		r.JSON(c.Writer, http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	user, err := rc.UserDao.GetUserByID(id)
	if err != nil {
		r.JSON(c.Writer, http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	if user.RoleID == "" {
		roles, err := rc.RoleDao.QueryAllRole()
		if err != nil {
			r.JSON(c.Writer, http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		r.JSON(c.Writer, http.StatusOK, roles)
		return
	}
	roleIds := strings.Split(user.RoleID, ",")
	roles, err := rc.RoleDao.TreeListByRoleID(roleIds)
	if err != nil {
		r.JSON(c.Writer, http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	r.JSON(c.Writer, http.StatusOK, roles)
}

// TreeList query all role to tree
func (rc RoleController) TreeList(c *gin.Context) {
	nodes, err := rc.RoleDao.QueryAllRole()
	if err != nil {
		r.JSON(c.Writer, http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	nodes = append(nodes, models.ZTreeNode{
		ID:      0,
		Name:    "顶级",
		Pid:     0,
		Open:    true,
		Checked: true,
	})
	r.JSON(c.Writer, http.StatusOK, nodes)
}

// SetAuthority set role authority
func (rc RoleController) SetAuthority(c *gin.Context) {
	roleID := c.PostForm("roleId")
	ids := c.PostForm("ids")
	id, err := strconv.ParseInt(roleID, 10, 64)
	if err != nil {
		r.JSON(c.Writer, http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	err = rc.RoleDao.SetAuthority(id, ids)
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
