package user_svc

import (
	"context"
	"dxxproject/internal/agreed/my_err"
	"dxxproject/internal/domain"
	"dxxproject/internal/jwt_utils/jwt_impl"
	"dxxproject/internal/model"
	"dxxproject/internal/obj_tranform"
	"dxxproject/internal/repo"
	"dxxproject/internal/svc/verify_code"
	"dxxproject/my_util"
	"dxxproject/pkg/password_utils"
	"dxxproject/pkg/sf_utils"
	"github.com/pkg/errors"
	"github.com/samber/do/v2"
	"gorm.io/gorm"
)

type UserSvc struct {
	passwordUtil  password_utils.PasswordUtilIF
	snow          sf_utils.SnowflakeIF
	verifyCodeSvc *verify_code.VerifyCodeSvc
	userRepo      *repo.Repo
	jwtUser       *jwt_impl.UserImpl
}

func (r *UserSvc) RefreshToken(ctx context.Context, refToken string) (token string, err error) {
	//校验refresh token
	info, err := r.jwtUser.RefreshValid(refToken)
	if err != nil {
		return
	}

	//生成新access token
	token, err = r.jwtUser.GenerateAccess(info)
	if err != nil {
		return
	}
	return
}

func (r *UserSvc) Login(ctx context.Context, domainUser *model.User) (
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
	err = r.passwordUtil.Compare(user.Password, domainUser.Password)
	if err != nil {
		err = my_err.ErrorUsernameOrPasswordWrong
		return
	}

	//校验通过,创建token
	accessToken, err = r.jwtUser.GenerateAccess(&jwt_impl.UserInfo{
		UserId:   domainUser.UserId,
		Username: domainUser.Username,
	})
	if err != nil {
		return
	}

	refreshToken, err = r.jwtUser.GenerateRefresh(&jwt_impl.UserInfo{
		UserId:   domainUser.UserId,
		Username: domainUser.Username,
	})
	return
}

func (r *UserSvc) Signup(ctx context.Context, userDomain *domain.User) (err error) {
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
	userModel := obj_tranform.UserDomainToModel(userDomain)

	err = r.userRepo.Insert(ctx, userModel)
	if err != nil {
		err = errors.Wrap(err, "添加用户失败")
		return
	}
	return
}

func NewUserSvc(injector do.Injector) (*UserSvc, error) {
	passwordUtil := do.MustInvoke[password_utils.PasswordUtilIF](injector)
	snow := do.MustInvoke[sf_utils.SnowflakeIF](injector)
	verifyCodeSvc := do.MustInvoke[*verify_code.VerifyCodeSvc](injector)
	userRepo := do.MustInvoke[*repo.Repo](injector)
	jwtUser := do.MustInvoke[*jwt_impl.UserImpl](injector)

	user := &UserSvc{
		passwordUtil:  passwordUtil,
		snow:          snow,
		verifyCodeSvc: verifyCodeSvc,
		userRepo:      userRepo,
		jwtUser:       jwtUser,
	}
	return user, nil
}
