package user_svc

import (
	"context"
	"dxxproject/agreed/domain"
	user_repo2 "dxxproject/internal/module/user/user_repo"
	"dxxproject/my/id_generator"
	"dxxproject/my/my_err"
	"dxxproject/my/my_util"
	"dxxproject/my/pwd_util"
	"github.com/pkg/errors"
)

type SignupSvc struct {
	passwordUtil pwd_util.PwdIface
	userRepo     *user_repo2.User
	snow         id_generator.IdGenIface
	//jwtUser       *jwt_user.UserImpl
	//verifyCodeSvc *vc_svc.VerifyCodeSvc

}

func (r *SignupSvc) Signup(ctx context.Context, userDomain *domain.User) (err error) {
	//拒绝简单密码
	ok := my_util.IsValidPassword(userDomain.Password, 3)
	if !ok {
		return my_err.ErrorPasswordTooSimple
	}

	// 对密码进行加密
	err, enPass := r.passwordUtil.Encrypt(userDomain.Password)
	if err != nil {
		return err
	}
	userDomain.Password = enPass

	//--添加用户
	userDomain.UserId = r.snow.GenSnowFlakeID() //生成 user id
	userModel := user_repo2.DomainToModel(userDomain)

	err = r.userRepo.Insert(ctx, userModel)
	if err != nil {
		err = errors.Wrap(err, "添加用户失败")
		return
	}
	return
}

func NewSignup(
	passwordUtil pwd_util.PwdIface,
	snow id_generator.IdGenIface,
	userRepo *user_repo2.User,
) (*SignupSvc, error) {

	user := &SignupSvc{
		passwordUtil: passwordUtil,
		snow:         snow,
		userRepo:     userRepo,
	}
	return user, nil
}
