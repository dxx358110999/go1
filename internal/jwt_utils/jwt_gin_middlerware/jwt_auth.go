package jwt_gin_middlerware

import (
	"dxxproject/internal/agreed/my_err"
	"dxxproject/internal/dto"
	"dxxproject/internal/jwt_utils/jwt_user"
	"github.com/gin-gonic/gin"
	"github.com/samber/do/v2"
)

func JwtAuthMiddleware(injector do.Injector) func(c *gin.Context) {
	/*
		基于JWT的认证中间件

		步骤:
		如果有refresh token,返回新的access token
		检查access token的合法性

	*/
	jwtUser := do.MustInvoke[*jwt_user.UserImpl](injector)

	return func(c *gin.Context) {
		/*
			客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI

			可以把Token放在Header的Authorization中，并使用Bearer开头
			Authorization: Bearer xxxxxxx.xxx.xxx  / X-TOKEN: xxx.xxx.xx

			这里使用自定的头
		*/

		//有 refresh ,检查有效,返回新的 access
		refreshHeader := c.Request.Header.Get(dto.HeaderRefreshToken)
		if refreshHeader != "" {
			//fmt.Println("refresh", refreshHeader)

			info, err := jwtUser.RefreshValid(refreshHeader)
			if err != nil {
				c.Error(err)
				c.Abort()
				return
			} else {
				// 将当前请求的userID信息保存到请求的上下文c上
				c.Set(dto.CtxUserIdKey, info.UserId)

				//生成新的access token并返回
				accessToken, err := jwtUser.GenerateAccess(info)
				if err != nil {
					c.Error(err)
					return
				}
				c.Header(dto.HeaderAccessToken, accessToken)
				return
			}

		}

		//检查 access
		accessHeader := c.Request.Header.Get(dto.HeaderAccessToken)
		if accessHeader == "" {
			c.Error(my_err.ErrorUserNotLogin)
			c.Abort()
			return
		} else {
			//fmt.Println("access", accessHeader)
			info, err := jwtUser.AccessValid(accessHeader)
			if err != nil {
				c.Error(my_err.ErrorUserNotLogin)
				c.Abort()
				return
			}
			c.Set(dto.CtxUserIdKey, info.UserId) //把 user id 放进ctx
			return

			// 按空格分割
			//parts := strings.SplitN(accessHeader, " ", 2)
			//if len(parts) == 2 && parts[0] == "Bearer" {}
			// parts[1]是获取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它

		}

		c.Next() // 后续的处理请求的函数中 可以用过c.Get(ContextUserIDKey) 来获取当前请求的用户信息
	}
}
