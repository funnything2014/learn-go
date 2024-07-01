package main

import "fmt"

func main() {
	//var a int8
	//var b int16
	//var c int32
	//var d int64
	//
	//var ua uint8
	//var ub uint16
	//var uc uint32
	//var ud uint64
	//
	//var e int // 动态类型，用的时候就会知道
	//
	//a = int8(b)
	//
	//var f1 float32
	//var f2 float64
	//
	//f1 = 3
	//f2 = 3.14

	var c byte  // 主要适用于存放字符
	var c2 rune // 也是字符
	c2 = '慕'
	c = 'a'
	fmt.Println(c)
	fmt.Printf("c=%c\n", c)
	fmt.Printf("c2=%c\n", c2)

	var name string
	name = "i am bobby"
	fmt.Println(name)
}
