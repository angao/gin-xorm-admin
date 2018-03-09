package forms

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
