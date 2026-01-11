package user_repo

import (
	"context"
	"dxxproject/agreed/model"
	"dxxproject/internal/module/user/user_cache"
	"dxxproject/internal/module/user/user_dao"
	"dxxproject/my/my_logger"
	"fmt"
)

type User struct {
	myLogger  my_logger.MyLoggerIF
	userDao   *user_dao.User
	userCache *user_cache.User
}

func (r *User) GetUserByUsername(ctx context.Context, username string) (err error, user *model.User) {
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

func (r *User) GetUserById(ctx context.Context, id int64) (err error, user *model.User) {
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

func (r *User) Insert(ctx context.Context, user *model.User) (err error) {
	return r.userDao.UserInsert(ctx, user)
}

func NewUserRepo(
	myLogger my_logger.MyLoggerIF,
	userDao *user_dao.User,
	userCache *user_cache.User,
) (*User, error) {

	ur := &User{
		myLogger:  myLogger,
		userDao:   userDao,
		userCache: userCache,
	}
	return ur, nil
}
