package models

// Menu 菜单权限
type Menu struct {
	// Id 主键
	Id int64
	// Code 菜单编号
	Code string
	// Pcode 菜单父编号
	Pcode string
	// Pcodes 当前菜单的所有父菜单编号
	Pcodes string
	// Name 菜单名称
	Name string
	// Icon 菜单图标
	Icon string
	// URL 地址
	URL string `xorm:"url"`
	// Num 菜单排序号
	Num int
	// Levels 菜单层级
	Levels int
	// IsMenu 是否是菜单（1：是  0：不是）
	IsMenu int `xorm:"ismenu"`
	// Tips 备注
	Tips string
	// Status 菜单状态 :  1:启用   0:不启用
	Status int
	// IsOpen 是否打开:    1:打开   0:不打开
	IsOpen int `xorm:"isopen"`

	Children []Menu `xorm:"-"`
}
