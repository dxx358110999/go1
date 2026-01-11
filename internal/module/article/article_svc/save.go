package article_svc

import (
	"context"
	"dxxproject/agreed/domain"
	"dxxproject/internal/module/article/article_repo"
)

type Save struct {
	repo *article_repo.ArticleRepo
}

func (r *Save) Save(ctx context.Context, domainArticle *domain.Article) error {
	/*
		新增或更新,通过是否有帖子id判断
	*/
	if domainArticle.Id == 0 {
		err := r.repo.Insert(ctx, domainArticle)
		if err != nil {
			return err
		}
	} else {
		//更新帖子数据
		err := r.repo.Update(ctx, domainArticle)
		if err != nil {
			return err
		}
	}
	return nil
}

func NewSave(repo *article_repo.ArticleRepo) (*Save, error) {
	return &Save{
		repo,
	}, nil
}
