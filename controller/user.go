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
	UserDao db.UserDao
}

// Home user home page
func (UserController) Home(c *gin.Context) {
	r.HTML(c.Writer, http.StatusOK, "system/user/user.html", gin.H{})
}

// List query all user
func (uc UserController) List(c *gin.Context) {
	var page forms.Page
	if err := c.Bind(&page); err != nil {
		r.JSON(c.Writer, http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	users, err := uc.UserDao.List(page)
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
func (uc UserController) Info(c *gin.Context) {
	var err error
	var user *models.UserRole
	session := sessions.Default(c)
	id, ok := session.Get("user_id").(int64)
	if ok {
		user, err = uc.UserDao.GetUserRole(id)
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
func (uc UserController) ToEdit(c *gin.Context) {
	id := c.Param("id")
	pid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		r.JSON(c.Writer, http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	user, err := uc.UserDao.GetUserRole(pid)
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
func (uc UserController) ToRoleAssign(c *gin.Context) {
	id := c.Param("id")
	pid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		r.JSON(c.Writer, http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	user, err := uc.UserDao.GetUserByID(pid)
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
func (uc UserController) Add(c *gin.Context) {
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
	err = uc.UserDao.Save(user)
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
func (uc UserController) Delete(c *gin.Context) {
	id := c.PostForm("userId")
	pid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		r.JSON(c.Writer, http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	err = uc.UserDao.Delete(pid)
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
func (uc UserController) Reset(c *gin.Context) {
	id := c.PostForm("userId")
	pid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		r.JSON(c.Writer, http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	var user *models.User
	user, err = uc.UserDao.GetUserByID(pid)
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
	err = uc.UserDao.Update(user)
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
func (uc UserController) SetRole(c *gin.Context) {
	roleIDs := c.PostForm("roleIds")
	userID := c.PostForm("userId")

	if roleIDs == "" || userID == "" {
		r.JSON(c.Writer, http.StatusBadRequest, gin.H{
			"message": "参数错误",
		})
		return
	}
	id, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		r.JSON(c.Writer, http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	user, err := uc.UserDao.GetUserByID(id)
	if err != nil {
		r.JSON(c.Writer, http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	user.RoleId = roleIDs
	err = uc.UserDao.Update(user)
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
func (uc UserController) Freeze(c *gin.Context) {
	userID := c.PostForm("userId")
	id, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		r.JSON(c.Writer, http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	user, err := uc.UserDao.GetUserByID(id)
	if err != nil {
		r.JSON(c.Writer, http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	user.Status = 2
	err = uc.UserDao.Update(user)
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
func (uc UserController) UnFreeze(c *gin.Context) {
	userID := c.PostForm("userId")
	id, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		r.JSON(c.Writer, http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	user, err := uc.UserDao.GetUserByID(id)
	if err != nil {
		r.JSON(c.Writer, http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	user.Status = 1
	err = uc.UserDao.Update(user)
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
func (uc UserController) ChangePwd(c *gin.Context) {
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
		user, err := uc.UserDao.GetUserByID(id)
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
		err = uc.UserDao.Update(user)
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
