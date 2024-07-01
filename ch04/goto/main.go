package main

import "fmt"

func main() {
	// goto语句可以让我们的代码跳到指定的代码块中运行
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				goto over
			}
			fmt.Println(i, j)
		}
	}

over:
	fmt.Println("over")

	// go语言的goto语句可以实现程序的跳转，goto语句使用场景最多的是程序的错误处理
}
