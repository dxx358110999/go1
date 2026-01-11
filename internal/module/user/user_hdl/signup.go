package user_hdl

import (
	"dxxproject/agreed/dto"
	"dxxproject/internal/module/user/user_svc"
	"github.com/gin-gonic/gin"
)

type SignupHdl struct {
	signupSvc *user_svc.SignupSvc
}

func (r *SignupHdl) Signup(ctx *gin.Context) {
	var err error

	// 获取参数和参数校验
	dtoSignup := new(UserSignup)
	err = ctx.ShouldBindJSON(dtoSignup)
	if err != nil {
		ctx.Error(err)
		return
	}

	//用户注册逻辑
	domainUser := SignupDtoToDomain(dtoSignup)
	err = r.signupSvc.Signup(ctx, domainUser)
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

func NewSignupHdl(signupSvc *user_svc.SignupSvc) (*SignupHdl, error) {
	u := &SignupHdl{
		signupSvc: signupSvc,
	}
	return u, nil
}
