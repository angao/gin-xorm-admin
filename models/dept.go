package models

// Dept 部门
type Dept struct {
	// Id 主键
	Id int64 `json:"id" form:"id"`
	// Num 排序
	Num int `json:"num" form:"num"`
	// Pid 父部门id
	Pid   int64  `json:"pid" form:"pid"`
	PName string `json:"pName" xorm:"<- pname"`
	// Pids 父级ids
	Pids string `json:"pids"`
	// SimpleName 简称
	SimpleName string `json:"simplename" form:"simplename" xorm:"simplename"`
	// fullname 全称
	Fullname string `json:"fullname" form:"fullname"`
	// Tips 提示
	Tips string `json:"tips" form:"tips"`
}

// TableName set table
func (Dept) TableName() string {
	return "sys_dept"
}
