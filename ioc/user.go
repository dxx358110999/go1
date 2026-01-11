package ioc

import (
	"dxxproject/internal/module/user/user_cache"
	"dxxproject/internal/module/user/user_dao"
	"dxxproject/internal/module/user/user_hdl"
	"dxxproject/internal/module/user/user_repo"
	"dxxproject/internal/module/user/user_svc"
	"dxxproject/my/id_generator/snow_impl"
	"dxxproject/my/jwt_utils/jwt_user"
	"dxxproject/my/my_logger"
	"dxxproject/my/pwd_util/pwd_impl"
	"github.com/redis/go-redis/v9"
	"github.com/samber/do/v2"
	"gorm.io/gorm"
)

func User(injector do.Injector) {

	do.Provide(injector, func(injector do.Injector) (*user_cache.User, error) {
		redisClient := do.MustInvoke[*redis.Client](injector)
		return user_cache.NewUserCache(redisClient)
	})
	do.Provide(injector, func(injector do.Injector) (*user_dao.User, error) {
		db := do.MustInvoke[*gorm.DB](injector)
		return user_dao.NewUserDao(db)
	})
	do.Provide(injector, func(injector do.Injector) (*user_repo.User, error) {
		myLogger := do.MustInvoke[*my_logger.MyLoggerZapImpl](injector)
		userDao := do.MustInvoke[*user_dao.User](injector)
		userCache := do.MustInvoke[*user_cache.User](injector)
		return user_repo.NewUserRepo(myLogger, userDao, userCache)
	})
	/*
		svc
	*/
	do.Provide(injector, func(injector do.Injector) (*user_svc.Login, error) {
		passwordUtil := do.MustInvoke[*pwd_impl.PwdImpl](injector)
		userRepo := do.MustInvoke[*user_repo.User](injector)
		jwtUser := do.MustInvoke[*jwt_user.UserImpl](injector)

		return user_svc.NewLoginSvc(
			passwordUtil,
			userRepo,
			jwtUser,
		)
	})
	do.Provide(injector, func(injector do.Injector) (*user_svc.SignupSvc, error) {
		passwordUtil := do.MustInvoke[*pwd_impl.PwdImpl](injector)
		userRepo := do.MustInvoke[*user_repo.User](injector)
		idGen := do.MustInvoke[*snow_impl.SnowflakeIMPL](injector)
		//jwtUser := do.MustInvoke[*jwt_user.UserImpl](injector)

		return user_svc.NewSignup(
			passwordUtil,
			idGen,
			userRepo,
		)
	})

	/*
		hdl
	*/
	do.Provide(injector, func(injector do.Injector) (*user_hdl.LoginHdl, error) {
		userSvc := do.MustInvoke[*user_svc.Login](injector)
		return user_hdl.NewUserHandlers(userSvc)
	})

	do.Provide(injector, func(injector do.Injector) (*user_hdl.SignupHdl, error) {
		userSvc := do.MustInvoke[*user_svc.SignupSvc](injector)
		return user_hdl.NewSignupHdl(userSvc)
	})

	do.Provide(injector, func(injector do.Injector) (*user_hdl.Profile, error) {
		userSvc := do.MustInvoke[*user_svc.Login](injector)
		return user_hdl.NewProfileHdl(userSvc)
	})
}
