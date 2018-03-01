package forms

// UserForm 绑定User查询
type UserForm struct {
	Order     string `form:"order" json:"order"`
	Offset    int    `form:"offset" json:"offset"`
	Limit     int    `form:"limit" json:"limit"`
	Name      string `form:"name" json:"name"`
	BeginTime string `form:"beginTime" json:"beginTime"`
	EndTime   string `form:"endTime" json:"endTime"`
	DeptID    int    `form:"deptid" json:"deptid"`
}
