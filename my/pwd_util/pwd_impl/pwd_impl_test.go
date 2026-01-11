package pwd_impl

import (
	"fmt"
	"testing"
	"time"
)

func TestEncrypt(t *testing.T) {
	util, err := NewPasswordUtil()
	if err != nil {
		return
	}

	for i := 0; i < 3; i++ {
		err, s := util.Encrypt("abc123")
		if err != nil {
			return
		}
		fmt.Println(s)
		err = util.Compare(s, "abc123")
		if err != nil {
			fmt.Println(err)
			return
		} else {
			fmt.Println("相同")
		}
		time.Sleep(1 * time.Second)
	}
}
