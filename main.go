package main

import (
	"dxxproject/main2"
	"fmt"
	"os"
)

func main() {
	err := main2.Main2()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
