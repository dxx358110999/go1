package route

import (
	"dxxproject/internal/verify_code/hdl"
	"github.com/gin-gonic/gin"
	"github.com/samber/do/v2"
)

func VerifyCodeRoute(injector do.Injector, prefixGroup *gin.RouterGroup) {
	vcHandlers := do.MustInvoke[*hdl.VerifyCodeHdl](injector)

	group := prefixGroup.Group("/code")
	group.POST("/registerCodeSendEmail", vcHandlers.SendUserRegisterCodeByEmail)
	group.POST("/registerCodeSendSms", vcHandlers.SendUserRegisterCodeBySms)

	//后续的接口需要认证
	//group.Use(gin_middleware.JwtAuthMiddleware(injector))
}
