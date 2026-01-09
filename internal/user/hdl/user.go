package hdl

import (
	"dxxproject/agreed/dto"
	"dxxproject/internal/user/obj_copy"
	"dxxproject/internal/user/svc"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userSvc *svc.UserSvc
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
	domainUser := obj_copy.LoginDtoToDomain(loginDto)
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
	domainUser := obj_copy.SignupDtoToDomain(dtoSignup)
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

func NewUserHandlers(userSvc *svc.UserSvc) (*UserHandler, error) {
	u := &UserHandler{
		userSvc: userSvc,
	}
	return u, nil
}
