package my_err

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
	"reflect"
	"testing"
)

func newErr() error {
	return ErrorDataExist
}

func TestErrorType(t *testing.T) {
	err := newErr()
	var err2 *MyError
	if errors.As(err, &err2) {
		fmt.Println(err2.msg, err2.appCode)
	} else {
		fmt.Println("类型转换失败")
	}

	if ok := errors.Is(err, err2); ok {
		fmt.Println("类型相同")

	} else {
		fmt.Println("类型不同")
	}

}

func TestReflectType(t *testing.T) {
	err := newErr()
	fmt.Println(reflect.TypeOf(err))

	var err2 *MyError
	fmt.Println(reflect.TypeOf(err2))

	var err3 *validator.ValidationErrors
	fmt.Println(reflect.TypeOf(err3))

	fmt.Println(reflect.TypeOf(err) == reflect.TypeOf(err2))
}
