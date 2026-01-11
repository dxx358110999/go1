package user_hdl

import (
	"dxxproject/agreed/dto"
	"dxxproject/internal/module/user/user_svc"
	"github.com/gin-gonic/gin"
)

type Profile struct {
	userSvc *user_svc.Login
}

func (r *Profile) Profile(ctx *gin.Context) {
	ctx.JSON(200, dto.ResponseBody{
		Code:    dto.AppCodeSuccess,
		Message: "用户资料",
		Data:    nil,
	})
	return
}

func NewProfileHdl(userSvc *user_svc.Login) (*Profile, error) {
	u := &Profile{
		userSvc: userSvc,
	}
	return u, nil
}
