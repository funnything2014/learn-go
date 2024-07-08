package main

import (
	"fmt"
	"time"
)

// python, java, php 多线程编程、多进程编程，多线程和多进程存在的问题主要是耗费内存
// 内存、线程切换，web2.0，用户级线程 绿程、轻量级线程、协程，asyncio-python php-swoole java-netty
// 内存占用小(2k)、切换快，go语言的协程，go语言诞生之后就只有协程可用-goroutine

func asyncPrint() {
	fmt.Println("bobby")
}

// 主协程
func main() {
	// 主死随从
	//go asyncPrint()

	// 匿名函数启动goroutine
	//go func() {
	//	for {
	//		time.Sleep(1 * time.Second)
	//		fmt.Println("bobby")
	//	}
	//}()

	// 1. 闭包 2. for循环的问题，i变量会被重用
	for i := 0; i < 100; i++ {
		//tmp := i
		//go func() {
		//	fmt.Println(tmp)
		//}()
		go func(i int) {
			fmt.Println(i)
		}(i)
	}

	fmt.Println("main goroutine")
	time.Sleep(10 * time.Second)
}
