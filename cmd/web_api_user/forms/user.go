package forms

type PwdLoginForm struct {
	Mobile   string `form:"mobile" json:"mobile" binding:"required,mobile"`
	PassWord string `form:"password" json:"password" binding:"required"`
}

type RegisterForm struct {
	Mobile   string `form:"mobile" json:"mobile" binding:"mobile"`
	PassWord string `form:"password" json:"password" binding:"required"`
}
