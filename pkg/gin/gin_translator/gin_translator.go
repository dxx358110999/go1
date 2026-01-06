package gin_translator

import (
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
	"github.com/pkg/errors"
)

type MyTranslator struct {
	Translator ut.Translator
}

func (r *MyTranslator) Init(locale string) (err error) {
	/*
		初始化翻译器
	*/
	// 修改gin框架中的Validator引擎属性，实现自定制
	v, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		zhT := zh.New() // 中文翻译器
		enT := en.New() // 英文翻译器

		// 第一个参数是备用（fallback）的语言环境
		// 后面的参数是应该支持的语言环境（支持多个）
		// uni := ut.NewLogger(zhT, zhT) 也是可以的
		uni := ut.New(enT, zhT, enT)

		// locale 通常取决于 http 请求头的 'Accept-Language'
		// 也可以使用 uni.FindTranslator(...) 传入多个locale进行查找
		translator, ok := uni.GetTranslator(locale)
		if !ok {
			err = fmt.Errorf("uni.GetTranslator(%s) failed,不支持的语言", locale)
			return
		}

		// 注册翻译器
		switch locale {
		case "en":
			err = enTranslations.RegisterDefaultTranslations(v, translator)
		case "zh":
			err = zhTranslations.RegisterDefaultTranslations(v, translator)
		default:
			err = enTranslations.RegisterDefaultTranslations(v, translator)
		}

		if err != nil {
			return
		}

		tf := &TranslatorMyFunctions{}
		v.RegisterTagNameFunc(tf.getJsonTag) // 注册一个获取json tag的自定义方法

		// 为SignUpParam注册自定义校验方法
		//v.RegisterStructValidation(SignUpParametersValidation, domain.SignUp{})

		r.Translator = translator

	} else {
		err = errors.New("翻译器类型转换失败")
		return
	}

	return
}
