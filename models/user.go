package models

import "time"

// User 用户
type User struct {
	// Id 主键
	Id       int64
	// Avatar 头像
	Avatar string
	// Account 账号
	Account string
	// Password 密码
	Password string
	// Salt md5密码盐
	Salt string
	// Name 名称
	Name     string
	// Birthday 生日
	Birthday time.Time
	// Sex 性别
	Sex int8
	// Email 电子邮件
	Email string
	// Phone 电话
	Phone    string
	// RoleId 角色ID
	RoleId string `xorm:"roleid"`
	// DeptId 部门Id
	DeptId int `xorm:"deptid"`
	// Status 状态(1：启用  2：冻结  3：删除）
	Status   int8
	// CreateAt 创建时间
	CreateTime time.Time `xorm:"created 'createtime'"`
}

// UserRole 用户角色
type UserRole struct {
	User `xorm:"extends"`
	Role `xorm:"extends"`
}

func (UserRole) TableName() string {
	return "sys_user"
}