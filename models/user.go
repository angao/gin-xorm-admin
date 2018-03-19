package models

import "time"

// Time marshaljson
type Time time.Time

// TimeFormat for format time
const TimeFormat = "2006-01-02 15:04:05"

// UnmarshalJSON parse byte to time
func (t *Time) UnmarshalJSON(data []byte) error {
	now, err := time.ParseInLocation(`"`+TimeFormat+`"`, string(data), time.Local)
	*t = Time(now)
	return err
}

// MarshalJSON parse json to byte
func (t Time) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(TimeFormat)+2)
	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, TimeFormat)
	b = append(b, '"')
	return b, nil
}

func (t Time) String() string {
	return time.Time(t).Format(TimeFormat)
}

// User 用户
type User struct {
	// Id 主键
	ID int64 `json:"id" form:"id" xorm:"pk id"`
	// Avatar 头像
	Avatar string `json:"avatar" form:"avatar"`
	// Account 账号
	Account string `json:"account" form:"account"`
	// Password 密码
	Password string `json:"password" form:"password"`
	// RePassword
	RePassword string `form:"rePassword" xorm:"-"`
	// Salt md5密码盐
	Salt string `json:"salt"`
	// Name 名称
	Name string `json:"name" form:"name"`
	// Birthday 生日
	Birthday Time `json:"birthday"`
	// Sex 性别
	Sex     int8   `json:"sex" form:"sex"`
	SexName string `json:"sexname" xorm:"<- sexname"`
	// Email 电子邮件
	Email string `json:"email" form:"email"`
	// Phone 电话
	Phone string `json:"phone" form:"phone"`
	// RoleId 角色ID
	RoleID   string `json:"roleid" xorm:"roleid"`
	RoleName string `json:"roleName" xorm:"<- role_name"`
	// DeptId 部门Id
	DeptID   int    `json:"deptid" form:"deptid" xorm:"deptid"`
	DeptName string `json:"deptName" xorm:"<- dept_name"`
	// Status 状态(1：启用  2：冻结  3：删除）
	Status     int8   `json:"status"`
	StatusName string `json:"statusname" xorm:"<- statusname"`
	// CreateAt 创建时间
	CreateTime Time `json:"createTime" xorm:"created 'createtime'"`
}

// UserRole 用户角色
type UserRole struct {
	User `xorm:"extends"`
	Role `xorm:"extends"`
}

// TableName set table
func (User) TableName() string {
	return "sys_user"
}
