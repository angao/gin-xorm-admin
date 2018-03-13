package middlewares

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/angao/gin-xorm-admin/db"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// ErrorHandler is a middleware to handle errors encountered during requests
func ErrorHandler(c *gin.Context) {
	c.Next()
	if len(c.Errors) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": c.Errors,
		})
	}
}

// NoRoute is a middleware to handle page not found during requests
func NoRoute(c *gin.Context) {
	c.HTML(http.StatusNotFound, "404.html", gin.H{})
}

// Auth is a middleware to handle the authenticate
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		userID, ok := session.Get("user_id").(int64)
		if ok {
			var roleDao db.RoleDao
			var userDao db.UserDao
			var menuDao db.MenuDao
			var permissions []string
			url := c.Request.URL.String()
			user, _ := userDao.GetUserByID(userID)
			roleIDs := strings.Split(user.RoleID, ",")
			for i := range roleIDs {
				id, _ := strconv.ParseInt(roleIDs[i], 10, 64)
				permission, _ := roleDao.GetURLByRoleID(id)
				permissions = append(permissions, permission...)
			}
			allPermissions, _ := menuDao.GetAllURL()
			if !Contains(allPermissions, url) {
				c.Next()
				return
			}
			if !Contains(permissions, url) {
				c.Redirect(http.StatusUnauthorized, "login")
				c.Abort()
				return
			}
			c.Next()
			return
		}
		c.Redirect(http.StatusFound, "login")
		c.Abort()
		return
	}
}

// Contains slice contain sub
func Contains(strs []string, s string) bool {
	for _, str := range strs {
		if str == s {
			return true
		}
	}
	return false
}
