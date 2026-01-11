package article_hdl

import (
	"dxxproject/agreed/dto"
	"dxxproject/internal/module/article/article_svc"
	"github.com/gin-gonic/gin"
)

type Save struct {
	svc *article_svc.Save
}

func (r *Save) Save(ctx *gin.Context) {
	var err error

	// 获取参数和参数校验
	dtoA := new(dto.Article)
	err = ctx.ShouldBindJSON(dtoA)
	if err != nil {
		ctx.Error(err)
		return
	}

	//存放
	dmA, err := DtoToDomain(dtoA) //转domain
	if err != nil {
		ctx.Error(err)
		return
	}

	err = r.svc.Save(ctx, dmA)
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

func NewSave(svc *article_svc.Save) (*Save, error) {
	return &Save{
		svc: svc,
	}, nil
}
