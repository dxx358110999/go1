package article

import (
	"context"
	"dxxproject/agreed/domain"
)

type ArticleRepo struct {
	dao *ArticleDao
}

func (r *ArticleRepo) Insert(ctx context.Context, article *domain.Article) error {
	model := copyDomainToModel(article)
	err := r.dao.Insert(ctx, model)
	if err != nil {
		return err
	}
	return err
}

func NewArticleRepo(dao *ArticleDao) (*ArticleRepo, error) {
	return &ArticleRepo{dao}, nil
}
