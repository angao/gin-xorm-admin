package controller

import (
	"log"
	"net/http"

	"github.com/angao/gin-xorm-admin/db"
	"github.com/angao/gin-xorm-admin/forms"
	"github.com/angao/gin-xorm-admin/models"
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

	var userForm forms.UserForm
	if err := c.Bind(&userForm); err != nil {
		r.JSON(c.Writer, http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	users, err := userDao.List(userForm)
	if err != nil {
		r.JSON(c.Writer, http.StatusBadRequest, gin.H{
			"error": err,
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
			log.Printf("%#v\n", err)
			return
		}
		r.HTML(c.Writer, http.StatusOK, "container.html", gin.H{
			"user":     user.User,
			"roleName": user.Role.Name,
		})
		return
	}
	r.HTML(c.Writer, http.StatusInternalServerError, "container.html", gin.H{
		"error": err,
		"user":  user,
	})
}

// Add handle add user page
func (UserController) Add(c *gin.Context) {
	r.HTML(c.Writer, http.StatusOK, "system/user/user_add.html", gin.H{})
}
