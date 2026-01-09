package article

import (
	"context"
	"dxxproject/agreed/domain"
)

type ArticleSvc struct {
	repo *ArticleRepo
}

func (r *ArticleSvc) Insert(ctx context.Context, dmA *domain.Article) error {
	return r.repo.Insert(ctx, dmA)
}

func NewArticleSvc(repo *ArticleRepo) (*ArticleSvc, error) {
	return &ArticleSvc{
		repo: repo,
	}, nil
}
