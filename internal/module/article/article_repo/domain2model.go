package article_repo

import (
	"dxxproject/agreed/domain"
	"dxxproject/agreed/model"
	"time"
)

func DomainToModel(da *domain.Article) *model.Article {
	return &model.Article{
		Id:         da.Id,
		AuthorID:   da.AuthorID,
		Title:      da.Title,
		Content:    da.Content,
		CreateTime: time.Time{},
		UpdateTime: time.Time{},
	}
}
