package user_svc

import (
	"context"
	"dxxproject/agreed/model"
	"github.com/samber/do/v2"
	"gorm.io/gorm"
)

type UserDao struct {
	db *gorm.DB
}

func (r *UserDao) GetUserById(ctx context.Context, id int64) (err error, user *model.User) {
	us, err := gorm.G[model.User](r.db).Where("id = ?", id).Take(ctx)

	if err != nil {
		return
	}

	user = &us
	return
}

func (r *UserDao) UserInsert(ctx context.Context, user *model.User) (err error) {
	/*
		添加用户
	*/
	//tx := db_init.Db.Create(user)
	//err = tx.Error
	err = gorm.G[model.User](r.db).Create(ctx, user)
	if err != nil {
		return err
	}
	return
}

func (r *UserDao) GetUserByUsername(ctx context.Context, username string) (err error, user *model.User) {
	/*
		检查指定用户名的用户是否存在
	*/

	us, err := gorm.G[model.User](r.db).Where("username = ?", username).Take(ctx)

	if err != nil {
		return
	}

	user = &us
	return

}

func NewUserDao(injector do.Injector) (*UserDao, error) {
	db := do.MustInvoke[*gorm.DB](injector)
	ud := &UserDao{
		db: db,
	}
	return ud, nil
}
