package handlers

import (
	"dxxproject/agreed/dto"
	"dxxproject/internal/obj_transform/user_trf"
	"dxxproject/internal/svc/user_svc"
	"github.com/gin-gonic/gin"
	"github.com/samber/do/v2"
)

type UserHandler struct {
	userSvc *user_svc.UserSvc
}

func (r *UserHandler) Profile(ctx *gin.Context) {
	ctx.JSON(200, dto.ResponseBody{
		Code:    dto.AppCodeSuccess,
		Message: "用户资料",
		Data:    nil,
	})
	return
}

func (r *UserHandler) Login(ctx *gin.Context) {
	var err error

	// 获取参数和参数校验
	loginDto := new(dto.UserLogin)
	err = ctx.ShouldBindJSON(loginDto)
	if err != nil {
		ctx.Error(err)
		return
	}

	//登录校验
	domainUser := user_trf.LoginDtoToDomain(loginDto)
	accessToken, refreshToken, err := r.userSvc.Login(ctx, domainUser)
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

func (r *UserHandler) Signup(ctx *gin.Context) {
	var err error

	// 获取参数和参数校验
	dtoSignup := new(dto.UserSignup)
	err = ctx.ShouldBindJSON(dtoSignup)
	if err != nil {
		ctx.Error(err)
		return
	}

	//用户注册逻辑
	domainUser := user_trf.SignupDtoToDomain(dtoSignup)
	err = r.userSvc.Signup(ctx, domainUser)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(200, dto.ResponseBody{
		Code:    dto.AppCodeSuccess,
		Message: "注册成功",
		Data:    nil,
	})
	return

}

func NewUserHandlers(injector do.Injector) (*UserHandler, error) {
	userSvc := do.MustInvoke[*user_svc.UserSvc](injector)
	u := &UserHandler{
		userSvc: userSvc,
	}
	return u, nil
}
