package controller

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/angao/gin-xorm-admin/db"
	"github.com/angao/gin-xorm-admin/forms"
	"github.com/angao/gin-xorm-admin/models"
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
	var page forms.Page
	if err := c.Bind(&page); err != nil {
		r.JSON(c.Writer, http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	roles, err := roleDao.List(page)
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
func (RoleController) Add(c *gin.Context) {
	var role models.Role
	if err := c.Bind(&role); err != nil {
		r.JSON(c.Writer, http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	var roleDao db.RoleDao
	err := roleDao.Save(role)
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
func (RoleController) ToEdit(c *gin.Context) {
	roleID := c.Param("roleId")
	if roleID == "" {
		r.JSON(c.Writer, http.StatusBadRequest, gin.H{
			"message": "参数错误",
		})
		return
	}
	id, err := strconv.ParseInt(roleID, 10, 64)
	if err != nil {
		r.JSON(c.Writer, http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	var roleDao db.RoleDao
	role, err := roleDao.Get(id)
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
func (RoleController) Edit(c *gin.Context) {
	var role models.Role
	if err := c.Bind(&role); err != nil {
		r.JSON(c.Writer, http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	var roleDao db.RoleDao
	err := roleDao.Update(role)
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
func (RoleController) Remove(c *gin.Context) {
	roleID := c.PostForm("roleId")
	id, err := strconv.ParseInt(roleID, 10, 64)
	if err != nil {
		r.JSON(c.Writer, http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	var roleDao db.RoleDao
	err = roleDao.Delete(id)
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
func (RoleController) ToAssign(c *gin.Context) {
	roleID := c.Param("roleId")
	if roleID == "" {
		r.JSON(c.Writer, http.StatusBadRequest, gin.H{
			"message": "参数错误",
		})
		return
	}
	id, err := strconv.ParseInt(roleID, 10, 64)
	if err != nil {
		r.JSON(c.Writer, http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	var roleDao db.RoleDao
	role, err := roleDao.Get(id)
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
			"error": err.Error(),
		})
		return
	}
	user, err := userDao.GetUserByID(id)
	if err != nil {
		r.JSON(c.Writer, http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	var roleDao db.RoleDao
	if user.RoleId == "" {
		roles, err := roleDao.QueryAllRole()
		if err != nil {
			r.JSON(c.Writer, http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		r.JSON(c.Writer, http.StatusOK, roles)
		return
	}
	roleIds := strings.Split(user.RoleId, ",")
	roles, err := roleDao.TreeListByRoleID(roleIds)
	if err != nil {
		r.JSON(c.Writer, http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	r.JSON(c.Writer, http.StatusOK, roles)
}

// TreeList query all role to tree
func (RoleController) TreeList(c *gin.Context) {
	var roleDao db.RoleDao
	nodes, err := roleDao.QueryAllRole()
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
func (RoleController) SetAuthority(c *gin.Context) {
	roleID := c.PostForm("roleId")
	ids := c.PostForm("ids")
	if roleID == "" {
		r.JSON(c.Writer, http.StatusBadRequest, gin.H{
			"message": "参数错误",
		})
		return
	}
	id, err := strconv.ParseInt(roleID, 10, 64)
	if err != nil {
		r.JSON(c.Writer, http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	var roleDao db.RoleDao
	err = roleDao.SetAuthority(id, ids)
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
