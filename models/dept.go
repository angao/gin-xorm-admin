package models

// Dept 部门
type Dept struct {
	// Id 主键
	Id int64

	// Num 排序
	Num int

	// Pid 父部门id
	Pid int

	// Pids 父级ids
	Pids string

	// SimpleName 简称
	SimpleName string `xorm:"simplename"`

	// fullname 全称
	Fullname string

	// Tips 提示
	Tips string
}
