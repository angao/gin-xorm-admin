package forms

// RoleForm 绑定role查询
type RoleForm struct {
	Order  string `form:"order" json:"order"`
	Offset int    `form:"offset" json:"offset"`
	Limit  int    `form:"limit" json:"limit"`
	Name   string `form:"roleName" json:"name"`
}
