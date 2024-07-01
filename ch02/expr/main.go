package main

import "fmt"

func main() {
	// 运算符 + - * / % ++ --
	var a, b = 1, 2
	var astr, bstr = "hello", "bobby"
	fmt.Println(a + b)
	fmt.Println(astr + bstr)

	// 逻辑运算符
	var abool, bbool = true, false
	if abool && bbool {
		fmt.Println("true")
	}

	// 位运算符，性能要求高的时候一般会考虑与运算
	var c = 60
	var d = 13
	fmt.Println(c & d)

	f := &c    // 取地址
	var g *int // 指针类型
	g = &d
	fmt.Println(*f)
	fmt.Println(*g)
}
