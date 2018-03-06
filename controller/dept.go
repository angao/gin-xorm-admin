package controller

import (
	"net/http"

	"github.com/angao/gin-xorm-admin/db"
	"github.com/angao/gin-xorm-admin/models"
	"github.com/gin-gonic/gin"
)

// DeptController operate dept
type DeptController struct{}

// List query dept
func (DeptController) List(c *gin.Context) {
	var deptDao db.DeptDao

	depts, err := deptDao.List()
	if err != nil {
		r.JSON(c.Writer, http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	treeNodes := build(depts)
	r.JSON(c.Writer, http.StatusOK, treeNodes)
}

func build(depts []models.Dept) []models.ZTreeNode {
	treeNodes := make([]models.ZTreeNode, 0)
	treeNodes = append(treeNodes, models.ZTreeNode{
		ID:      0,
		Name:    "顶级",
		Pid:     0,
		Open:    true,
		Checked: true,
	})
	for _, dept := range depts {
		node := models.ZTreeNode{
			ID:   dept.Id,
			Name: dept.Fullname,
			Pid:  int64(dept.Pid),
		}
		if node.Pid == 0 {
			node.Open = true
		} else {
			node.Open = false
		}
		treeNodes = append(treeNodes, node)
	}
	return treeNodes
}
