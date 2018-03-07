package models

// Dept 部门
type Dept struct {
	// Id 主键
	Id int64 `json:"id"`
	// Num 排序
	Num int `json:"num"`
	// Pid 父部门id
	Pid   int    `json:"pid"`
	PName string `json:"pName" xorm:"<- pname"`
	// Pids 父级ids
	Pids string `json:"pids"`
	// SimpleName 简称
	SimpleName string `json:"simplename" xorm:"simplename"`
	// fullname 全称
	Fullname string `json:"fullname"`
	// Tips 提示
	Tips string `json:"tips"`
}
