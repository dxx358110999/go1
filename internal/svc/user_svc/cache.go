package user_svc

import (
	"context"
	"dxxproject/agreed/model"
	"dxxproject/my/my_err"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
	"github.com/samber/do/v2"
	"time"
)

type UserCache struct {
	redisClient *redis.Client
}

func keyById(id int64) (key string) {
	key = fmt.Sprintf("user:id:%d", id)
	return
}

func (r *UserCache) GetById(ctx context.Context, id int64) (err error, user *model.User) {
	/*
		数据不存在,包会返回redis.Nil,是一种错误
	*/

	userId := keyById(id)
	bytes, err := r.redisClient.Get(ctx, userId).Bytes()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			err = my_err.ErrorDataNotExist
		}
		return
	}

	user = &model.User{}
	err = json.Unmarshal(bytes, &user)
	if err != nil {
		return
	}
	return
}

func (r *UserCache) SetById(ctx context.Context, user *model.User) (err error) {
	userExpire := 10 * time.Minute
	value, err := json.Marshal(user)
	if err != nil {
		return err
	}
	userId := keyById(user.UserId)
	err = r.redisClient.Set(ctx, userId, value, userExpire).Err()
	if err != nil {
		return err
	}
	return
}

func NewUserCache(injector do.Injector) (*UserCache, error) {
	redisClient := do.MustInvoke[*redis.Client](injector)
	uc := &UserCache{
		redisClient: redisClient,
	}
	return uc, nil
}
