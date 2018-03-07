package models

// Role 用户角色
type Role struct {
	// Id 主键
	Id int64 `json:"id"`
	// Num 序号
	Num int `json:"num"`
	// Pid 父角色id
	Pid   int    `json:"pId" xorm:"pId"`
	PName string `json:"pName" xorm:"<- p_name"`
	// Name 角色名称
	Name string `json:"name"`
	// DeptId 部门名称
	DeptID int `json:"deptid" xorm:"deptid"`
	// DeptName 部门名称
	DeptName string `json:"deptName" xorm:"<-"`
	// Tips 描述
	Tips string `json:"tips"`
}
