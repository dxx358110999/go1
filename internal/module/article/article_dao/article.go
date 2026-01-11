package article_dao

import (
	"context"
	"dxxproject/agreed/model"
	"fmt"
	"gorm.io/gorm"
)

type Article struct {
	db *gorm.DB
}

func (r *Article) Insert(ctx context.Context, article *model.Article) error {
	result := gorm.WithResult()
	err := gorm.G[model.Article](r.db, result).Create(ctx, article)
	if err != nil {
		return err
	}
	return nil
}

func (r *Article) UpdateById(ctx context.Context, article *model.Article) error {
	/*
		tips:
		1,更新时检查author id,防止a用户的帖子被b用户修改
	*/
	rows, err := gorm.G[model.Article](r.db).
		Where("id = ? and author_id = ?", article.Id, article.AuthorID).
		Updates(ctx, *article)

	if err != nil {
		return err
	} else {
		if rows == 0 {
			return fmt.Errorf("更新失败,帖子不存在,或者与创作者不匹配,%d,%d", article.Id, article.AuthorID)
		}
	}
	return err
}

func NewArticleDao(db *gorm.DB) (*Article, error) {
	return &Article{db: db}, nil
}
