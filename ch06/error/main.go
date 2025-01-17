package main

import (
	"errors"
	"fmt"
)

func A() (int, error) {
	//panic("this is a panic") // panic会导致程序的退出，平时开发中不要随便用，一般我们在哪里用到：我们一个服务启动的过程中，必须有些依赖服务准备好，日志文件存在，配置文件没有问题，这个时候服务才能启动的时候
	// 我们的服务启动检查中发现了任何一个不满足，那就主动调用panic
	// 但是你的服务一旦启动了，这个时候你的某行代码不小心panic，你的程序挂了，这是重大事故
	// 但是架不住有些地方我们的代码写的不小心会导致被动触发panic
	// recover 这个函数能捕获到panic
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered if A:", r)
		}
	}()
	var names map[string]string
	names["go"] = "go工程师"
	//fmt.Println("this is a func")
	return 0, errors.New("this is an error")
}

func main() {
	if _, err := A(); err != nil {
		fmt.Println(err)
	}

	// panic 这个函数会导致你的程序退出，不推荐随便使用panic

	// 1. defer需要放在panic之前定义，recover只有在defer调用的函数中才会生效
	// 2. recover处理异常后，逻辑并不会恢复到panic的那个点去
	// 3. 多个defer会形成栈，后定义的defer会先执行
}
