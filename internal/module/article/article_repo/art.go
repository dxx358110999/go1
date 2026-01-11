package article_repo

import (
	"context"
	"dxxproject/agreed/domain"
	"dxxproject/internal/module/article/article_dao"
	"dxxproject/my/id_generator"
)

type ArticleRepo struct {
	dao   *article_dao.Article
	idGen id_generator.IdGenIface
}

func (r *ArticleRepo) Insert(ctx context.Context, article *domain.Article) error {

	article.Id = r.idGen.GenSnowFlakeID() // article没有id,新增
	model := DomainToModel(article)

	//写库
	err := r.dao.Insert(ctx, model)
	if err != nil {
		return err
	}
	return err
}

func (r *ArticleRepo) Update(ctx context.Context, article *domain.Article) error {
	model := DomainToModel(article)
	err := r.dao.UpdateById(ctx, model)
	if err != nil {
		return err
	}
	return err
}

func NewArticleRepo(dao *article_dao.Article,
	idGen id_generator.IdGenIface,
) (*ArticleRepo, error) {
	return &ArticleRepo{
		dao:   dao,
		idGen: idGen,
	}, nil
}
