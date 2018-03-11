package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/angao/gin-xorm-admin/db"
	"github.com/angao/gin-xorm-admin/models"
	"github.com/gin-gonic/gin"
)

// DeptController operate dept
type DeptController struct {
	DeptDao db.DeptDao
}

// Tree query dept
func (dc DeptController) Tree(c *gin.Context) {
	depts, err := dc.DeptDao.List("")
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
func (dc DeptController) List(c *gin.Context) {
	name := c.PostForm("condition")
	depts, err := dc.DeptDao.List(name)
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

// Add add dept
func (dc DeptController) Add(c *gin.Context) {
	var dept models.Dept
	if err := c.Bind(&dept); err != nil {
		r.JSON(c.Writer, http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	if dept.SimpleName == "" {
		r.JSON(c.Writer, http.StatusBadRequest, gin.H{
			"message": "参数错误",
		})
		return
	}
	dept, err := deptSetPid(dept, dc.DeptDao)
	if err != nil {
		r.JSON(c.Writer, http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	err = dc.DeptDao.Save(dept)
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
func (dc DeptController) ToEdit(c *gin.Context) {
	deptID := c.Param("deptId")
	id, err := strconv.ParseInt(deptID, 10, 64)
	if err != nil {
		r.JSON(c.Writer, http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	dept, err := dc.DeptDao.Get(id)
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

// Edit update dept
func (dc DeptController) Edit(c *gin.Context) {
	var dept models.Dept
	if err := c.Bind(&dept); err != nil {
		r.JSON(c.Writer, http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	if dept.Id == 0 || dept.SimpleName == "" {
		r.JSON(c.Writer, http.StatusBadRequest, gin.H{
			"message": "参数错误",
		})
		return
	}
	dept, err := deptSetPid(dept, dc.DeptDao)
	if err != nil {
		r.JSON(c.Writer, http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = dc.DeptDao.Update(dept)
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

//Delete dept
func (dc DeptController) Delete(c *gin.Context) {
	deptID := c.PostForm("deptId")
	id, err := strconv.ParseInt(deptID, 10, 64)
	if err != nil {
		r.JSON(c.Writer, http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	err = dc.DeptDao.Delete(id)
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

func deptSetPid(dept models.Dept, deptDao db.DeptDao) (models.Dept, error) {
	if dept.Pid == 0 {
		dept.Pids = "[0],"
	} else {
		pDept, err := deptDao.Get(dept.Pid)
		if err != nil {
			return dept, err
		}
		dept.Pids = pDept.Pids + "[" + fmt.Sprintf("%d", dept.Pid) + "],"
	}
	return dept, nil
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
