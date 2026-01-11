package ioc

import (
	"dxxproject/internal/module/article/article_dao"
	"dxxproject/internal/module/article/article_hdl"
	"dxxproject/internal/module/article/article_repo"
	"dxxproject/internal/module/article/article_svc"
	"dxxproject/my/id_generator"
	"github.com/samber/do/v2"
	"gorm.io/gorm"
)

func Article(injector do.Injector) {
	//redisClient := do.MustInvoke[*redis.Client](injector)

	//myLogger := do.MustInvoke[my_logger.MyLoggerIF](injector)

	do.Provide(injector, func(injector do.Injector) (*article_dao.Article, error) {
		db := do.MustInvoke[*gorm.DB](injector)
		return article_dao.NewArticleDao(db)
	})

	do.Provide(injector, func(injector do.Injector) (*article_repo.ArticleRepo, error) {
		articleDao := do.MustInvoke[*article_dao.Article](injector)
		idGen := do.MustInvoke[id_generator.IdGenIface](injector)
		return article_repo.NewArticleRepo(articleDao, idGen)
	})

	do.Provide(injector, func(injector do.Injector) (*article_svc.Save, error) {
		articleRepo := do.MustInvoke[*article_repo.ArticleRepo](injector)
		return article_svc.NewSave(articleRepo)
	})

	do.Provide(injector, func(injector do.Injector) (*article_hdl.Save, error) {
		articleSvc := do.MustInvoke[*article_svc.Save](injector)
		return article_hdl.NewSave(articleSvc)
	})

}
