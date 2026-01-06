package verify_code_svc

import (
	"math/rand"
	"strconv"
	"time"
)

func generateCode() string {
	/*
		生成6位数字验证码
	*/
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)
	code := random.Intn(900000) + 100000 // 生成100000-999999的随机数
	return strconv.Itoa(code)
}
