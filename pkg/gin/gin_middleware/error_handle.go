package gin_middleware

import (
	"dxxproject/my/my_err"
	"fmt"
	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"reflect"
	"strings"
)

type ErrorHandler struct {
	Translator ut.Translator
}

func NewErrorHandler(
	translator ut.Translator) *ErrorHandler {
	mid := &ErrorHandler{
		Translator: translator,
	}
	return mid
}

func (rec ErrorHandler) ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if len(c.Errors) > 0 {
			err := c.Errors.Last().Err
			fmt.Println("ErrorHandler", err.Error())
			zap.L().Error(err.Error())
			c.JSON(200, rec.errorToResponseBody(err))
		}
	}
}

func (rec ErrorHandler) removeTopStruct(fields map[string]string) map[string]string {
	/*
		去除提示信息中的结构体名称
	*/

	result := map[string]string{}
	for field, errString := range fields {
		newField := field[strings.Index(field, ".")+1:] //去除结构体名称
		result[newField] = errString
	}
	return result
}

func (rec ErrorHandler) validateErrorToMsg(err error) (msg any) {
	/*
		请求参数校验失败,返回错误响应
	*/
	var errs validator.ValidationErrors

	// 判断err是不是validator.ValidationErrors
	if ok := errors.As(err, &errs); ok {
		fields := errs.Translate(rec.Translator) //做一次翻译
		msg = rec.removeTopStruct(fields)
		//msg = err.Error()
	} else {
		msg = err.Error()
	}

	return msg
}

func (rec ErrorHandler) errorToResponseBody(err error) (body *ResponseBody) {
	//不要把完整业务逻辑错误给外部
	var myErr *my_err.MyError
	var validateErr validator.ValidationErrors
	body = &ResponseBody{
		Code:    0,
		Message: nil,
		Data:    nil,
	}

	if reflect.TypeOf(err) == reflect.TypeOf(validateErr) {
		body.Message = rec.validateErrorToMsg(err)
		body.Code = my_err.ErrorInvalidParameters.Code()
		return
	}

	if reflect.TypeOf(err) == reflect.TypeOf(myErr) {
		if ok := errors.As(err, &myErr); ok {
			body.Message = myErr.Error()
			body.Code = myErr.Code()

		} else {
			body.Message = err.Error()
		}
		return
	}

	body.Message = my_err.ErrorServerBusy.Error()
	body.Code = my_err.ErrorServerBusy.Code()
	return
}
