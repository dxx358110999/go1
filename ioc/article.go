package ioc

import (
	"dxxproject/internal/article"
	"github.com/samber/do/v2"
	"gorm.io/gorm"
)

func Article(injector do.Injector) {
	//redisClient := do.MustInvoke[*redis.Client](injector)

	//myLogger := do.MustInvoke[my_logger.MyLoggerIF](injector)

	do.Provide(injector, func(injector do.Injector) (*article.ArticleDao, error) {
		db := do.MustInvoke[*gorm.DB](injector)
		return article.NewArticleDao(db)
	})

	do.Provide(injector, func(injector do.Injector) (*article.ArticleRepo, error) {
		articleDao := do.MustInvoke[*article.ArticleDao](injector)
		return article.NewArticleRepo(articleDao)
	})

	do.Provide(injector, func(injector do.Injector) (*article.ArticleSvc, error) {
		articleRepo := do.MustInvoke[*article.ArticleRepo](injector)

		return article.NewArticleSvc(articleRepo)
	})

	do.Provide(injector, func(injector do.Injector) (*article.ArticleHdl, error) {
		articleSvc := do.MustInvoke[*article.ArticleSvc](injector)
		return article.NewArticleHdl(articleSvc)
	})

}
