package user

import (
	"dxxproject/internal/user/cache"
	"dxxproject/internal/user/dao"
	"dxxproject/internal/user/hdl"
	"dxxproject/internal/user/repo"
	"dxxproject/internal/user/svc"
	vcSvc "dxxproject/internal/verify_code/svc"
	"dxxproject/my/jwt_utils/jwt_user"
	"dxxproject/my/my_logger"
	"dxxproject/my/passwd_util"
	"dxxproject/pkg/snowflake_ok"
	"github.com/redis/go-redis/v9"
	"github.com/samber/do/v2"
	"gorm.io/gorm"
)

func Provide(injector do.Injector) {
	redisClient := do.MustInvoke[*redis.Client](injector)
	db := do.MustInvoke[*gorm.DB](injector)
	myLogger := do.MustInvoke[my_logger.MyLoggerIF](injector)
	userDao := do.MustInvoke[*dao.UserDao](injector)
	userCache := do.MustInvoke[*cache.User](injector)
	passwordUtil := do.MustInvoke[passwd_util.PasswordUtilIface](injector)
	snow := do.MustInvoke[snowflake_ok.SnowflakeIface](injector)
	userRepo := do.MustInvoke[*repo.UserRepo](injector)
	jwtUser := do.MustInvoke[*jwt_user.UserImpl](injector)

	verifyCodeSvc := do.MustInvoke[*vcSvc.VerifyCodeSvc](injector)
	userSvc := do.MustInvoke[*svc.UserSvc](injector)

	do.Provide(injector, func(injector do.Injector) (*cache.User, error) {
		return cache.NewUserCache(redisClient)
	})
	do.Provide(injector, func(injector do.Injector) (*dao.UserDao, error) {
		return dao.NewUserDao(db)
	})
	do.Provide(injector, func(injector do.Injector) (*repo.UserRepo, error) {
		return repo.NewUserRepo(myLogger, userDao, userCache)
	})
	do.Provide(injector, func(injector do.Injector) (*svc.UserSvc, error) {
		return svc.NewUserSvc(
			passwordUtil,
			snow,
			verifyCodeSvc,
			userRepo,
			jwtUser,
		)
	})
	do.Provide(injector, func(injector do.Injector) (*hdl.UserHandler, error) {
		return hdl.NewUserHandlers(userSvc)
	})

}
