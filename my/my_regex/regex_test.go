package my_regex

import (
	"fmt"
	"testing"
)

func TestVerifyEmailFormat(t *testing.T) {
	fmt.Println(CompiledPatternEmail.MatchString("12345@126.com"))
	fmt.Println(CompiledPatternEmail.MatchString("12345126.com"))

}
