package models

// Page 分页
type Page struct {
	Order  string `form:"order" json:"order"`
	Offset int    `form:"offset" json:"offset"`
	Limit  int    `form:"limit" json:"limit"`

	// 查询条件
	Name      string `form:"name" json:"name"`
	BeginTime string `form:"beginTime" json:"beginTime"`
	EndTime   string `form:"endTime" json:"endTime"`
	DeptID    int    `form:"deptid" json:"deptid"`
}
