package models

import (
	"time"
)

// Notice 通知公告
type Notice struct {
	// 主键
	Id int64
	// 标题
	Title string
	// 类型
	Type string
	// 内容
	Content string
	// 创建时间
	CreateTime time.Time `xorm:"created 'createtime'"`
	Creater int64 
}