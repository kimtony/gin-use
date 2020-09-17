package main

import (
	"fmt"
	"gin-demo/helpers"
)

func test() {
	//加密密码
	hash, err := helpers.GeneratePassword("123456")
	fmt.Println("加密密码:", string(hash), err)
	//解析密码
	verify, err := helpers.ValidatePassword("123456", string(hash))
	fmt.Println("验证密码:", verify, err)
}
