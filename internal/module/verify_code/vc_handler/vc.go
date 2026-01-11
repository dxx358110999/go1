package vc_handler

import (
	"dxxproject/agreed/biz"
	"dxxproject/agreed/dto"
	"dxxproject/internal/module/verify_code/vc_svc"
	"dxxproject/my/jwt_utils/jwt_user"
	"github.com/gin-gonic/gin"
	"github.com/samber/do/v2"
)

type VerifyCodeHdl struct {
	jwtUser *jwt_user.UserImpl
	vcSvc   *vc_svc.VerifyCodeSvc
}

func (r *VerifyCodeHdl) SendUserRegisterCodeByEmail(ctx *gin.Context) {
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

func (r *VerifyCodeHdl) SendUserRegisterCodeBySms(ctx *gin.Context) {
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

func NewVerifyCodeHdl(injector do.Injector) (*VerifyCodeHdl, error) {
	jwtUser := do.MustInvoke[*jwt_user.UserImpl](injector)
	vcSvc := do.MustInvoke[*vc_svc.VerifyCodeSvc](injector)
	vc := &VerifyCodeHdl{
		jwtUser: jwtUser,
		vcSvc:   vcSvc,
	}
	return vc, nil
}
