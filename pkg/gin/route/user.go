package route

import (
	"dxxproject/internal/module/user/user_hdl"
	"dxxproject/my/jwt_utils/jwt_gin_middlerware"
	"github.com/gin-gonic/gin"
	"github.com/samber/do/v2"
)

func AddWebRoute(injector do.Injector, prefixGroup *gin.RouterGroup) {
	loginHdl := do.MustInvoke[*user_hdl.LoginHdl](injector)
	signupHdl := do.MustInvoke[*user_hdl.SignupHdl](injector)
	profileHdl := do.MustInvoke[*user_hdl.Profile](injector)

	userGroup := prefixGroup.Group("/user")
	userGroup.POST("/signup", signupHdl.Signup)
	userGroup.POST("/login", loginHdl.Login)

	//后续的接口需要认证
	userGroup.Use(jwt_gin_middlerware.JwtAuthMiddleware(injector))
	userGroup.POST("/profile", profileHdl.Profile)

}
