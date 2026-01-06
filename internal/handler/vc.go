package handler

import (
	"dxxproject/internal/agreed/biz"
	"dxxproject/internal/dto"
	"dxxproject/internal/jwt_utils/jwt_impl"
	"dxxproject/internal/svc/verify_code"
	"github.com/gin-gonic/gin"
	"github.com/samber/do/v2"
)

type VerifyCodeHandlers struct {
	jwtUser *jwt_impl.UserImpl
	vcSvc   *verify_code.VerifyCodeSvc
}

func (r *VerifyCodeHandlers) SendUserRegisterCodeByEmail(ctx *gin.Context) {
	var err error

	// 获取参数和参数校验
	params := new(dto.Email)
	err = ctx.ShouldBindJSON(params)
	if err != nil {
		ctx.Error(err)
		return
	}

	//发送验证码
	err = r.vcSvc.SendCodeByEmail(ctx, biz.UserRegister, params.Email)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(200, dto.ResponseBody{
		Code:    dto.AppCodeSuccess,
		Message: "登录",
		Data:    nil,
	})
	return

}

func (r *VerifyCodeHandlers) SendUserRegisterCodeBySms(ctx *gin.Context) {
	type Phone struct {
		Phone string `json:"phone"`
	}
	var err error

	// 获取参数和参数校验
	params := new(Phone)
	err = ctx.ShouldBindJSON(params)
	if err != nil {
		ctx.Error(err)
		return
	}

	//发送验证码
	codeParams := map[string]string{}
	err = r.vcSvc.SendCodeBySms(ctx, biz.UserRegister, params.Phone,
		codeParams,
	)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(200, dto.ResponseBody{
		Code:    dto.AppCodeSuccess,
		Message: "登录",
		Data:    nil,
	})
	return

}

func NewVerifyCodeHandlers(injector do.Injector) (*VerifyCodeHandlers, error) {
	jwtUser := do.MustInvoke[*jwt_impl.UserImpl](injector)
	vcSvc := do.MustInvoke[*verify_code.VerifyCodeSvc](injector)
	vc := &VerifyCodeHandlers{
		jwtUser: jwtUser,
		vcSvc:   vcSvc,
	}
	return vc, nil
}
