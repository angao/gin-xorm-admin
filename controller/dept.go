package controller

import (
	"net/http"
	"strconv"

	"github.com/angao/gin-xorm-admin/db"
	"github.com/angao/gin-xorm-admin/models"
	"github.com/gin-gonic/gin"
)

// DeptController operate dept
type DeptController struct{}

// Tree query dept
func (DeptController) Tree(c *gin.Context) {
	var deptDao db.DeptDao

	depts, err := deptDao.List("")
	if err != nil {
		r.JSON(c.Writer, http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	treeNodes := build(depts)
	r.JSON(c.Writer, http.StatusOK, treeNodes)
}

// Index dept index
func (DeptController) Index(c *gin.Context) {
	r.HTML(c.Writer, http.StatusOK, "system/dept/dept.html", gin.H{})
}

// List query all dept
func (DeptController) List(c *gin.Context) {
	name := c.PostForm("condition")
	var deptDao db.DeptDao
	depts, err := deptDao.List(name)
	if err != nil {
		r.JSON(c.Writer, http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	r.JSON(c.Writer, http.StatusOK, depts)
}

// ToAdd to add page
func (DeptController) ToAdd(c *gin.Context) {
	r.HTML(c.Writer, http.StatusOK, "system/dept/dept_add.html", gin.H{})
}

// ToEdit to edit page
func (DeptController) ToEdit(c *gin.Context) {
	deptID := c.Param("deptId")
	id, err := strconv.ParseInt(deptID, 10, 64)
	if err != nil {
		r.JSON(c.Writer, http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	var deptDao db.DeptDao
	dept, err := deptDao.Get(id)
	if err != nil {
		r.JSON(c.Writer, http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	r.HTML(c.Writer, http.StatusOK, "system/dept/dept_edit.html", gin.H{
		"dept": dept,
	})
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
