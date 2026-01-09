package article

import (
	"context"
	"dxxproject/agreed/model"
	"gorm.io/gorm"
)

type ArticleDao struct {
	db *gorm.DB
}

func (r *ArticleDao) Insert(ctx context.Context, article *model.Article) error {
	result := gorm.WithResult()
	err := gorm.G[model.Article](r.db, result).Create(ctx, article)
	if err != nil {
		return err
	}
	return nil
}

func NewArticleDao(db *gorm.DB) (*ArticleDao, error) {
	return &ArticleDao{db: db}, nil
}
