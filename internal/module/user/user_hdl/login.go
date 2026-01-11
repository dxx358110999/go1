package user_hdl

import (
	"dxxproject/agreed/dto"
	"dxxproject/internal/module/user/user_svc"
	"github.com/gin-gonic/gin"
)

type LoginHdl struct {
	loginSvc *user_svc.Login
}

func (r *LoginHdl) Login(ctx *gin.Context) {
	var err error

	// 获取参数和参数校验
	loginDto := new(UserLogin)
	err = ctx.ShouldBindJSON(loginDto)
	if err != nil {
		ctx.Error(err)
		return
	}

	//登录校验
	domainUser := LoginDtoToDomain(loginDto)
	accessToken, refreshToken, err := r.loginSvc.Login(ctx, domainUser)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.Header(dto.HeaderAccessToken, accessToken)   //使用 header 返回token,需要配置cors
	ctx.Header(dto.HeaderRefreshToken, refreshToken) //使用 header 返回token,需要配置cors

	ctx.JSON(200, dto.ResponseBody{
		Code:    dto.AppCodeSuccess,
		Message: "登录",
		Data:    nil,
	})
	return

}

func NewUserHandlers(loginSvc *user_svc.Login) (*LoginHdl, error) {
	u := &LoginHdl{
		loginSvc: loginSvc,
	}
	return u, nil
}
