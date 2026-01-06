package toolkit

import (
	"encoding/base64"
	"strings"
)

func StringToBase64RemoveEqual(str string) string {
	data := []byte(str)
	encodedStr := base64.StdEncoding.EncodeToString(data)
	removeEqual := strings.Replace(encodedStr, "=", "", -1)
	return removeEqual
}
