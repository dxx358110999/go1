package jwt_user

import "github.com/golang-jwt/jwt/v5"

/*
自定义声明结构体并内嵌jwt.StandardClaims

jwt包自带的jwt.StandardClaims只包含了官方字段
我们这里需要额外记录一个username字段，所以要自定义结构体
如果想要保存更多信息，都可以添加到这个结构体中
*/

type UserInfo struct {
	UserId   int64  `json:"user_id"`
	Username string `json:"username"`
}

type UserClaims struct {
	UserId   int64  `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}
