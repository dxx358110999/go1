package customized

import (
	"github.com/dlclark/regexp2"
)

const (
	patternEmail = `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*` //匹配电子邮箱

)

var (
	CompiledPatternEmail *regexp2.Regexp
)

func init() {
	CompiledPatternEmail = regexp2.MustCompile(patternEmail, 0)
}
