package controller

import (
	"fmt"
	"net/http"

	"github.com/angao/gin-xorm-admin/db"
	"github.com/angao/gin-xorm-admin/models"
	"github.com/gin-gonic/gin"
)

// DeptController operate dept
type DeptController struct{}

type DeptForm struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Pid    string `json:"pId"`
	IsOpen bool   `json:"isOpen"`
	Open   bool   `json:"open"`
	Check  bool   `json:"check"`
}

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
	deptRets := build(depts)
	r.JSON(c.Writer, http.StatusOK, deptRets)
}

func build(depts []models.Dept) []DeptForm {
	deptRets := make([]DeptForm, 0)
	deptRets = append(deptRets, DeptForm{
		Id:     "0",
		Name:   "顶级",
		Pid:    "0",
		IsOpen: true,
		Open:   true,
		Check:  true,
	})
	for _, dept := range depts {
		deptNew := DeptForm{}
		deptNew.Id = fmt.Sprintf("%d", dept.Id)
		deptNew.Check = false
		deptNew.Name = dept.Fullname
		deptNew.Pid = fmt.Sprintf("%d", dept.Pid)
		if deptNew.Pid == "0" {
			deptNew.IsOpen = true
			deptNew.Open = true
		} else {
			deptNew.IsOpen = false
			deptNew.Open = false
		}
		deptRets = append(deptRets, deptNew)
	}
	return deptRets
}
