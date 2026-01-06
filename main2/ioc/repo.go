package ioc

import (
	"dxxproject/internal/repo"
	"github.com/samber/do/v2"
)

func repoIoc(injector do.Injector) (err error) {
	do.Provide(injector, repo.NewUserRepo)

	return err
}
