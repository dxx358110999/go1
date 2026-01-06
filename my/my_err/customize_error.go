package my_err

type MyError struct {
	appCode int
	msg     string
}

func (e *MyError) Error() string {
	return e.msg
}
func (e *MyError) Code() int {
	return e.appCode
}

func NewMyError(appCode int, msg string) *MyError {
	return &MyError{
		appCode: appCode,
		msg:     msg,
	}

}

var (
	ErrorDataExist               = NewMyError(4000, "数据已存在")
	ErrorDataNotExist            = NewMyError(4001, "数据不存在")
	ErrorInvalidPassword         = NewMyError(4002, "密码校验错误")
	ErrorInvalidID               = NewMyError(4003, "无效id")
	ErrorInvalidToken            = NewMyError(4005, "token失效")
	ErrorInvalidParameters       = NewMyError(4006, "参数错误")
	ErrorUsernameOrPasswordWrong = NewMyError(4002, "用户名或密码错误")
	ErrorPasswordTooSimple       = NewMyError(4007, "密码需要数字,大写字母,小写字母,特殊字符中的3类")
	ErrRateLimited               = NewMyError(4008, "操作频繁")
	ErrVerifyCodeNotMatch        = NewMyError(4009, "验证码不匹配")

	ErrorServerBusy = NewMyError(4999, "服务器忙")

	ErrorUserNotLogin = NewMyError(10401, "用户未登录")
)
