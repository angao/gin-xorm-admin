package models

// Role 用户角色
type Role struct {
	// Id 主键
	Id int64
	// Num 序号
	Num int
	// Pid 父角色id
	Pid int
	// Name 角色名称
	Name string
	// DeptId 部门名称
	DeptId int `xorm:"deptid"`
	// Tips 描述
	Tips string
}
