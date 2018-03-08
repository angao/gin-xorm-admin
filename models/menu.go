package models

// Menu 菜单权限
type Menu struct {
	// Id 主键
	Id int64 `json:"id" form:"id"`
	// Code 菜单编号
	Code string `json:"code" form:"code"`
	// Pcode 菜单父编号
	Pcode string `json:"pcode" form:"pcode"`
	// PcodeName 父菜单名称
	PcodeName string `json:"pcodeName" xorm:"-"`
	Pid       int64  `json:"pid" xorm:"-"`
	// Pcodes 当前菜单的所有父菜单编号
	Pcodes string `json:"pcodes"`
	// Name 菜单名称
	Name string `json:"name" form:"name"`
	// Icon 菜单图标
	Icon string `json:"icon" form:"icon"`
	// URL 地址
	URL string `json:"url" xorm:"url" form:"url"`
	// Num 菜单排序号
	Num int `json:"num" form:"num"`
	// Levels 菜单层级
	Levels int `json:"levels"`
	// IsMenu 是否是菜单（1：是  0：不是）
	IsMenu int `json:"isMenu" xorm:"ismenu" form:"ismenu"`
	// Tips 备注
	Tips string `json:"tips"`
	// Status 菜单状态 :  1:启用   0:不启用
	Status int `json:"status"`
	// IsOpen 是否打开:    1:打开   0:不打开
	IsOpen int `json:"isopen" xorm:"isopen"`

	IsMenuName string `json:"isMenuName"`
	StatusName string `json:"statusName"`

	Children []Menu `xorm:"-"`
}

// TableName set table
func (Menu) TableName() string {
	return "sys_menu"
}
