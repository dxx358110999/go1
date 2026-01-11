package user_svc

import (
	"context"
	"dxxproject/agreed/domain"
	user_repo2 "dxxproject/internal/module/user/user_repo"
	"dxxproject/my/jwt_utils/jwt_user"
	"dxxproject/my/my_err"
	"dxxproject/my/my_util"
	"dxxproject/my/pwd_util"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type Login struct {
	pwdUtil pwd_util.PwdIface
	//snow          id_generator.IdGenIface
	//verifyCodeSvc *vc_svc.VerifyCodeSvc
	userRepo *user_repo2.User
	jwtUser  *jwt_user.UserImpl
}

func (r *Login) Login(ctx context.Context, domainUser *domain.User) (
	accessToken string,
	refreshToken string,
	err error) {
	//拒绝简单密码
	ok := my_util.IsValidPassword(domainUser.Password, 3)
	if !ok {
		err = my_err.ErrorPasswordTooSimple
		return
	}

	//--查询用户
	err, user := r.userRepo.GetUserByUsername(ctx, domainUser.Username) //查询用户是否存在
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = my_err.ErrorDataNotExist
			return
		}
		return
	}
	domainUser.UserId = user.UserId //web需要使用user id

	// 对密码进行核对
	err = r.pwdUtil.Compare(user.Password, domainUser.Password)
	if err != nil {
		err = my_err.ErrorUsernameOrPasswordWrong
		return
	}

	//校验通过,创建token
	accessToken, err = r.jwtUser.GenerateAccess(&jwt_user.UserInfo{
		UserId:   domainUser.UserId,
		Username: domainUser.Username,
	})
	if err != nil {
		return
	}

	refreshToken, err = r.jwtUser.GenerateRefresh(&jwt_user.UserInfo{
		UserId:   domainUser.UserId,
		Username: domainUser.Username,
	})
	return
}

func NewLoginSvc(
	passwordUtil pwd_util.PwdIface,
	userRepo *user_repo2.User,
	jwtUser *jwt_user.UserImpl,
) (*Login, error) {

	user := &Login{
		pwdUtil:  passwordUtil,
		userRepo: userRepo,
		jwtUser:  jwtUser,
	}
	return user, nil
}
