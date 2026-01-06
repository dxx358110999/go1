package biz

type Biz struct {
	Code string
	Name string
}

func newBiz(code string, name string) Biz {
	return Biz{
		Code: code,
		Name: name,
	}
}

var (
	UserRegister = newBiz("UserRegister", "用户注册")
)
