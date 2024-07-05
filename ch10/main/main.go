package main

import (
	"fmt"
	u "learngo/ch10/user" // 导入的别名
	//. "learngo/ch10/user" 尽量少用
	//_ "learngo/ch09/user" 初始化，自动调用init
) // 包的路径

func main() {
	c := u.Course{
		Name: "go",
	}
	fmt.Println(u.GetCourse(c))
}
