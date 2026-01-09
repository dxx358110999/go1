package article

import (
	"github.com/samber/do/v2"
	"gorm.io/gorm"
)

func Provide(injector do.Injector) {
	//redisClient := do.MustInvoke[*redis.Client](injector)

	//myLogger := do.MustInvoke[my_logger.MyLoggerIF](injector)

	do.Provide(injector, func(injector do.Injector) (*ArticleDao, error) {
		db := do.MustInvoke[*gorm.DB](injector)
		return NewArticleDao(db)
	})

	do.Provide(injector, func(injector do.Injector) (*ArticleRepo, error) {
		articleDao := do.MustInvoke[*ArticleDao](injector)
		return NewArticleRepo(articleDao)
	})

	do.Provide(injector, func(injector do.Injector) (*ArticleSvc, error) {
		articleRepo := do.MustInvoke[*ArticleRepo](injector)

		return NewArticleSvc(articleRepo)
	})

	do.Provide(injector, func(injector do.Injector) (*ArticleHdl, error) {
		articleSvc := do.MustInvoke[*ArticleSvc](injector)
		return NewArticleHdl(articleSvc)
	})

}
