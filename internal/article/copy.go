package article

import (
	"dxxproject/agreed/domain"
	"dxxproject/agreed/dto"
	"dxxproject/agreed/model"
	"time"
)

func copyDtoToDomain(dtoA *dto.Article) *domain.Article {
	return &domain.Article{
		Id:         0,
		AuthorID:   dtoA.AuthorID,
		Title:      dtoA.Title,
		Content:    dtoA.Content,
		CreateTime: time.Time{},
		UpdateTime: time.Time{},
	}
}

func copyDomainToModel(da *domain.Article) *model.Article {
	return &model.Article{
		Id:         da.Id,
		AuthorID:   da.AuthorID,
		Title:      da.Title,
		Content:    da.Content,
		CreateTime: time.Time{},
		UpdateTime: time.Time{},
	}
}
