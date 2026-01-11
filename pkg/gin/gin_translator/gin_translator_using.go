package gin_translator

import (
	"dxxproject/internal/module/user/user_hdl"
	"github.com/go-playground/validator/v10"
	"reflect"
	"strings"
)

type TranslatorMyFunctions struct {
}

func (r *TranslatorMyFunctions) getJsonTag(fld reflect.StructField) string {
	name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
	if name == "-" {
		return ""
	}
	return name
}

type TranslatorMyValidator struct {
}

func (r *TranslatorMyValidator) SignUpParametersValidation(level validator.StructLevel) {
	/*
		自定义SignUpParam结构体校验函数
	*/
	userSignup := level.Current().Interface().(user_hdl.UserSignup)

	if userSignup.Password != userSignup.RePassword {
		// 输出错误提示信息，最后一个参数就是传递的param
		level.ReportError(userSignup.RePassword, "re_password", "RePassword", "eqfield", "password")
	}
}
