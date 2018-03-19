package models

// Notice 通知公告
type Notice struct {
	// 主键
	ID int64 `json:"id" form:"id" xorm:"pk id"`
	// 标题
	Title string `json:"title" form:"title"`
	// 类型
	Type int `json:"type" form:"type"`
	// 内容
	Content string `json:"content" form:"content"`
	// 创建时间
	CreateTime Time   `json:"createtime" xorm:"created 'createtime'"`
	Creater    int64  `json:"creater" form:"creater"`
	CreateName string `json:"createName" xorm:"<- create_name"`
}

// TableName set table
func (Notice) TableName() string {
	return "sys_notice"
}
