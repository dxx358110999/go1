package article

import (
	"dxxproject/agreed/dto"
	"github.com/gin-gonic/gin"
)

type ArticleHdl struct {
	svc *ArticleSvc
}

func (r *ArticleHdl) Save(ctx *gin.Context) {
	var err error

	// 获取参数和参数校验
	dtoA := new(dto.Article)
	err = ctx.ShouldBindJSON(dtoA)
	if err != nil {
		ctx.Error(err)
		return
	}

	//存放
	dmA := copyDtoToDomain(dtoA)
	err = r.svc.Insert(ctx, dmA)
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

func NewArticleHdl(svc *ArticleSvc) (*ArticleHdl, error) {
	return &ArticleHdl{
		svc: svc,
	}, nil
}
