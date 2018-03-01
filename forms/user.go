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

// UserAddForm 添加管理员
type UserAddForm struct {
	Account    string `form:"account"`
	Sex        int8   `form:"sex"`
	Password   string `form:"password"`
	RePassword string `form:"rePassword"`
	Email      string `form:"email"`
	Name       string `form:"name"`
	Birthday   string `form:"birthday"`
	Phone      string `form:"phone"`
}
