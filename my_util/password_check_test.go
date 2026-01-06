package my_util

import (
	"fmt"
	"testing"
)

func TestCheckPassword(t *testing.T) {
	target := 3
	fmt.Println(IsValidPassword("123", target))
	fmt.Println(IsValidPassword("abc123", target))
	fmt.Println(IsValidPassword("Aabc123", target))
	fmt.Println(IsValidPassword("!abc123", target))
}
