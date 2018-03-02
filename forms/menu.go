package forms

// MenuForm 绑定Menu表单
type MenuForm struct {
	Order string `form:"order" json:"order"`
	Name  string `form:"menuName" json:"name"`
}
