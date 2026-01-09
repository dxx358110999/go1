package repo

import (
	"context"
	"dxxproject/agreed/model"
	"dxxproject/internal/user/cache"
	"dxxproject/internal/user/dao"
	"dxxproject/my/my_logger"
	"fmt"
)

type UserRepo struct {
	myLogger  my_logger.MyLoggerIF
	userDao   *dao.UserDao
	userCache *cache.User
}

func (r *UserRepo) GetUserByUsername(ctx context.Context, username string) (err error, user *model.User) {
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

func (r *UserRepo) GetUserById(ctx context.Context, id int64) (err error, user *model.User) {
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

func (r *UserRepo) Insert(ctx context.Context, user *model.User) (err error) {
	return r.userDao.UserInsert(ctx, user)
}

func NewUserRepo(
	myLogger my_logger.MyLoggerIF,
	userDao *dao.UserDao,
	userCache *cache.User,
) (*UserRepo, error) {

	ur := &UserRepo{
		myLogger:  myLogger,
		userDao:   userDao,
		userCache: userCache,
	}
	return ur, nil
}
