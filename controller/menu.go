package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// MenuController handle menu request
type MenuController struct {
}

// Index handle /menu
func (MenuController) Index(c *gin.Context) {
	r.HTML(c.Writer, http.StatusOK, "system/menu/menu.html", gin.H{})
}
