package article_hdl

import (
	"dxxproject/agreed/domain"
	"dxxproject/agreed/dto"
	"github.com/pkg/errors"
	"strconv"
	"time"
)

func DtoToDomain(dtoA *dto.Article) (*domain.Article, error) {
	authorId, err := strconv.ParseInt(dtoA.AuthorID, 10, 64) // 第二个参数是基数（这里为10，表示十进制），第三个参数是位大小，这里为64位
	if err != nil {
		err = errors.Wrap(err, "author id 转换失败")
		return nil, err
	}

	var articleId int64
	if dtoA.Id == "" {
		articleId = 0
	} else {
		articleId, err = strconv.ParseInt(dtoA.Id, 10, 64)
		if err != nil {
			err = errors.Wrap(err, "author id 转换失败")
			return nil, err
		}
	}

	return &domain.Article{
		Id:         articleId,
		AuthorID:   authorId,
		Title:      dtoA.Title,
		Content:    dtoA.Content,
		CreateTime: time.Time{},
		UpdateTime: time.Time{},
	}, nil
}
