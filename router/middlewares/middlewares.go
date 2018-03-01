package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions"
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
		_, ok := session.Get("user_id").(int64)
		if ok {
			c.Next()
			return
		}
		c.Redirect(http.StatusFound, "login")
		c.Abort()
		return
	}
}
