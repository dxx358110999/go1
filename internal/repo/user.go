package repo

import (
	"context"
	"dxxproject/internal/cache"
	"dxxproject/internal/dao"
	"dxxproject/internal/model"
	"dxxproject/pkg/my_logger"
	"fmt"
	"github.com/samber/do/v2"
)

type Repo struct {
	myLogger  my_logger.MyLoggerIF
	userDao   *dao.Dao
	userCache *cache.CacheUser
}

func (r *Repo) GetUserByUsername(ctx context.Context, username string) (err error, user *model.User) {
	err, user = r.userDao.GetUserByUsername(ctx, username)
	if err != nil {
		return
	}

	//写缓存
	err = r.userCache.SetById(ctx, user)
	if err != nil { //缓存失败,记录日志,不返回
		msg := fmt.Sprintln("user 写入缓存失败", user.UserId, err.Error())
		r.myLogger.Info(msg)
	}

	return
}

func (r *Repo) GetUserById(ctx context.Context, id int64) (err error, user *model.User) {
	//查缓存
	err, user = r.userCache.GetById(ctx, id)

	if err != nil { //缓存失败,记录日志,不返回错误
		msg := fmt.Sprintln("user 查询缓存失败", id, err.Error())
		r.myLogger.Info(msg)
	} else {
		return
	}

	//查库
	err, user = r.userDao.GetUserById(ctx, id)
	if err != nil {
		return
	}

	//写缓存
	err = r.userCache.SetById(ctx, user)
	if err != nil { //缓存失败,记录日志,不返回
		msg := fmt.Sprintln("user 写入缓存失败", id, err.Error())
		r.myLogger.Info(msg)
	}
	return
}

func (r *Repo) Insert(ctx context.Context, user *model.User) (err error) {
	return r.userDao.UserInsert(ctx, user)
}

func NewUserRepo(injector do.Injector) (*Repo, error) {
	myLogger := do.MustInvoke[my_logger.MyLoggerIF](injector)
	userDao := do.MustInvoke[*dao.Dao](injector)
	userCache := do.MustInvoke[*cache.CacheUser](injector)
	ur := &Repo{
		myLogger:  myLogger,
		userDao:   userDao,
		userCache: userCache,
	}
	return ur, nil
}
