package jwt_user

import (
	"dxxproject/config_prepare/app_config"
	"dxxproject/my/my_err"
	"github.com/pkg/errors"
	"github.com/samber/do/v2"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type UserImpl struct {
	accessSecret  []byte
	accessExpire  time.Duration
	refreshSecret []byte
	refreshExpire time.Duration
}

func (r *UserImpl) GenerateAccess(info *UserInfo) (tokenStr string, err error) {
	claims := UserClaims{
		UserId:   info.UserId,
		Username: info.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "",
			Subject:   "",
			Audience:  nil,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(r.accessExpire)),
			NotBefore: nil,
			IssuedAt:  nil,
			ID:        "",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) // 使用指定的签名方法创建签名对象
	tokenStr, err = token.SignedString(r.accessSecret)         // 使用指定的secret签名并获得完整的编码后的字符串token

	if err != nil {
		return "", err
	}
	return
}

func (r *UserImpl) GenerateRefresh(info *UserInfo) (tokenStr string, err error) {
	claims := UserClaims{
		UserId:   info.UserId,
		Username: info.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "",
			Subject:   "",
			Audience:  nil,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(r.refreshExpire)),
			NotBefore: nil,
			IssuedAt:  nil,
			ID:        "",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) // 使用指定的签名方法创建签名对象
	tokenStr, err = token.SignedString(r.refreshSecret)        // 使用指定的secret签名并获得完整的编码后的字符串token

	if err != nil {
		return "", err
	}
	return
}

func (r *UserImpl) AccessValid(tokenString string) (info *UserInfo, err error) {
	claims := new(UserClaims)

	//解析token
	keyFunc := func(token *jwt.Token) (i any, err error) {
		return r.accessSecret, nil
	}
	token, err := jwt.ParseWithClaims(tokenString, claims,
		keyFunc, jwt.WithLeeway(5*time.Second))
	if err != nil {
		err = errors.Wrap(err, "解析用户token失败")
		return
	}

	//校验
	if token.Valid { // 校验token
		info = &UserInfo{
			UserId:   claims.UserId,
			Username: claims.Username,
		}
		return
	} else {
		err = my_err.ErrorInvalidToken
	}

	return
}
func (r *UserImpl) RefreshValid(tokenString string) (info *UserInfo, err error) {
	claims := new(UserClaims)

	//解析token
	keyFunc := func(token *jwt.Token) (i any, err error) {
		return r.refreshSecret, nil
	}
	token, err := jwt.ParseWithClaims(tokenString, claims,
		keyFunc, jwt.WithLeeway(5*time.Second))
	if err != nil {
		err = errors.Wrap(err, "解析用户token失败")
		return
	}

	//校验
	if token.Valid { // 校验token
		info = &UserInfo{
			UserId:   claims.UserId,
			Username: claims.Username,
		}
		return
	} else {
		err = my_err.ErrorInvalidToken
	}

	return
}
func NewJwtUserImpl(injector do.Injector) (impl *UserImpl, err error) {
	cfg := do.MustInvoke[*app_config.AppConfig](injector).JwtUser

	impl = &UserImpl{
		accessExpire:  time.Duration(cfg.AccessExpire) * time.Minute,
		refreshExpire: time.Duration(cfg.RefreshExpire) * time.Minute,
		accessSecret:  []byte(cfg.AccessSecret),
		refreshSecret: []byte(cfg.RefreshSecret),
	}

	return
}

func Provide(injector do.Injector) {
	do.Provide(injector, NewJwtUserImpl)
}

//var _ UserIface = new(UserImpl) //检查实现
