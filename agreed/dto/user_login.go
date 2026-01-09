package dto

type UserLogin struct {
	/*
		登录
		http请求
	*/
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
