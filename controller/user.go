package controller

import (
	"net/http"
	"strconv"

	"github.com/angao/gin-xorm-admin/db"
	"github.com/angao/gin-xorm-admin/forms"
	"github.com/angao/gin-xorm-admin/models"
	"github.com/angao/gin-xorm-admin/utils"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// UserController handle user request
type UserController struct {
}

// Home user home page
func (UserController) Home(c *gin.Context) {
	r.HTML(c.Writer, http.StatusOK, "system/user/user.html", gin.H{})
}

// List query all user
func (UserController) List(c *gin.Context) {
	var userDao db.UserDao

	var page forms.Page
	if err := c.Bind(&page); err != nil {
		r.JSON(c.Writer, http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	users, err := userDao.List(page)
	if err != nil {
		r.JSON(c.Writer, http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	r.JSON(c.Writer, http.StatusOK, gin.H{
		"data": users,
	})
}

// Info is handle user info
func (UserController) Info(c *gin.Context) {
	var userDao db.UserDao
	var err error
	var user *models.UserRole
	session := sessions.Default(c)
	id, ok := session.Get("user_id").(int64)
	if ok {
		user, err = userDao.GetUserRole(id)
		if err != nil {
			r.JSON(c.Writer, http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		r.HTML(c.Writer, http.StatusOK, "system/user/user_view.html", gin.H{
			"user":     user.User,
			"roleName": user.Role.Name,
		})
		return
	}
	r.HTML(c.Writer, http.StatusInternalServerError, "system/user/user_view.html", gin.H{
		"error": err.Error(),
	})
}

// ToAdd handle add user page
func (UserController) ToAdd(c *gin.Context) {
	r.HTML(c.Writer, http.StatusOK, "system/user/user_add.html", gin.H{})
}

// ToEdit handle edit user paget
func (UserController) ToEdit(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		r.JSON(c.Writer, http.StatusBadRequest, gin.H{
			"error": "参数错误",
		})
		return
	}
	var userDao db.UserDao
	pid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		r.JSON(c.Writer, http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	user, err := userDao.GetUserRole(pid)
	if err != nil {
		r.JSON(c.Writer, http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	r.HTML(c.Writer, http.StatusOK, "system/user/user_edit.html", gin.H{
		"user":     user.User,
		"roleName": user.Role.Name,
	})
}

// ToRoleAssign handle user role
func (UserController) ToRoleAssign(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		r.JSON(c.Writer, http.StatusBadRequest, gin.H{
			"error": "参数错误",
		})
		return
	}
	var userDao db.UserDao
	pid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		r.JSON(c.Writer, http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	user, err := userDao.GetUserByID(pid)
	if err != nil {
		r.JSON(c.Writer, http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	r.HTML(c.Writer, http.StatusOK, "system/user/user_roleassign.html", gin.H{
		"user": user,
	})
}

// Add handle save user
func (UserController) Add(c *gin.Context) {
	var userDao db.UserDao
	var userAddForm forms.UserAddForm

	if err := c.Bind(&userAddForm); err != nil {
		r.JSON(c.Writer, http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if userAddForm.Password != userAddForm.RePassword {
		r.JSON(c.Writer, http.StatusBadRequest, gin.H{
			"error": "密码不一致",
		})
		return
	}

	salt := utils.RandomString(5)
	password, err := utils.Encrypt(userAddForm.Password, salt)
	if err != nil {
		r.JSON(c.Writer, http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	user := models.User{
		Name:     userAddForm.Name,
		Account:  userAddForm.Account,
		Email:    userAddForm.Email,
		Sex:      userAddForm.Sex,
		Password: password,
		Salt:     salt,
	}
	err = userDao.Save(user)
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

// Delete 删除用户
func (UserController) Delete(c *gin.Context) {
	id := c.PostForm("userId")
	if id == "" {
		r.JSON(c.Writer, http.StatusBadRequest, gin.H{
			"message": "参数错误",
		})
		return
	}
	var userDao db.UserDao
	pid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		r.JSON(c.Writer, http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	err = userDao.Delete(pid)
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

// Reset password
func (UserController) Reset(c *gin.Context) {
	id := c.PostForm("userId")
	if id == "" {
		r.JSON(c.Writer, http.StatusBadRequest, gin.H{
			"message": "参数错误",
		})
		return
	}
	var userDao db.UserDao
	pid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		r.JSON(c.Writer, http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	var user *models.User
	user, err = userDao.GetUserByID(pid)
	if err != nil {
		r.JSON(c.Writer, http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	user.Id = pid
	password, err := utils.Encrypt("111111", user.Salt)
	if err != nil {
		r.JSON(c.Writer, http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	user.Password = password
	err = userDao.Update(user)
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

// SetRole set user role
func (UserController) SetRole(c *gin.Context) {
	roleIDs := c.PostForm("roleIds")
	userID := c.PostForm("userId")

	if roleIDs == "" || userID == "" {
		r.JSON(c.Writer, http.StatusBadRequest, gin.H{
			"message": "参数错误",
		})
		return
	}
	var userDao db.UserDao
	id, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		r.JSON(c.Writer, http.StatusBadRequest, gin.H{
			"message": err.Error(),
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
	user.RoleId = roleIDs
	err = userDao.Update(user)
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

// Freeze user
func (UserController) Freeze(c *gin.Context) {
	userID := c.PostForm("userId")
	if userID == "" {
		r.JSON(c.Writer, http.StatusBadRequest, gin.H{
			"message": "参数错误",
		})
		return
	}
	var userDao db.UserDao
	id, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		r.JSON(c.Writer, http.StatusBadRequest, gin.H{
			"message": err.Error(),
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
	user.Status = 2
	err = userDao.Update(user)
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

// UnFreeze user
func (UserController) UnFreeze(c *gin.Context) {
	userID := c.PostForm("userId")
	if userID == "" {
		r.JSON(c.Writer, http.StatusBadRequest, gin.H{
			"message": "参数错误",
		})
		return
	}
	var userDao db.UserDao
	id, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		r.JSON(c.Writer, http.StatusBadRequest, gin.H{
			"message": err.Error(),
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
	user.Status = 1
	err = userDao.Update(user)
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

// ToChangePasswd to change password page
func (UserController) ToChangePasswd(c *gin.Context) {
	r.HTML(c.Writer, http.StatusOK, "system/user/user_chpwd.html", gin.H{})
}

// ChangePwd change password
func (UserController) ChangePwd(c *gin.Context) {
	oldPwd := c.PostForm("oldPwd")
	newPwd := c.PostForm("newPwd")
	rePwd := c.PostForm("rePwd")
	if oldPwd == "" || newPwd == "" || newPwd != rePwd {
		r.JSON(c.Writer, http.StatusBadRequest, gin.H{
			"message": "参数为空，或新密码和确认密码不一致",
		})
		return
	}
	session := sessions.Default(c)
	id, ok := session.Get("user_id").(int64)
	if ok {
		var userDao db.UserDao
		user, err := userDao.GetUserByID(id)
		if err != nil {
			r.JSON(c.Writer, http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		password, err := utils.Encrypt(newPwd, user.Salt)
		if err != nil {
			r.JSON(c.Writer, http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}
		user.Password = password
		err = userDao.Update(user)
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
}
