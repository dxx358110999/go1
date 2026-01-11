package user_hdl

/*
binding标签,是validator用来校验数据
*/

type UserSignup struct {
	/*
		注册
	*/
	Username     string  `json:"username" binding:"required"`
	Email        *string `json:"email"`
	Phone        *string `json:"phone"`
	Password     string  `json:"password" binding:"required"`
	RePassword   string  `json:"checkPass" binding:"required,eqfield=Password"`
	Code         string  `json:"code" binding:"required"`
	RegisterType string  `json:"registerType" binding:"required"`
}

type UserLogin struct {
	/*
		登录
		http请求
	*/
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
