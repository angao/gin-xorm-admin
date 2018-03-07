package models

// Role 用户角色
type Role struct {
	// Id 主键
	Id int64 `json:"id" form:"id"`
	// Num 序号
	Num int `json:"num" form:"num"`
	// Pid 父角色id
	Pid   int    `json:"pId" form:"pid" xorm:"pId"`
	PName string `json:"pName" xorm:"<- p_name"`
	// Name 角色名称
	Name string `json:"name" form:"name"`
	// DeptId 部门名称
	DeptID int `json:"deptid" form:"deptid" xorm:"deptid"`
	// DeptName 部门名称
	DeptName string `json:"deptName" xorm:"<-"`
	// Tips 别名
	Tips string `json:"tips" form:"tips"`
}

// TableName set table
func (Role) TableName() string {
	return "sys_role"
}
