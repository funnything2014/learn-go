package main

import (
	"fmt"
	"strconv"
)

func printSlice(data []string) {
	data[0] = "java"

	for i := 0; i < 10; i++ {
		data = append(data, strconv.Itoa(i))
	}
}

func main() {
	// go的slice在函数参数传递的时候，是值传递还是引用传递：值传递，效果又呈现出了引用的效果（不完全是）
	courses := []string{"go", "grpc", "gin"}
	printSlice(courses)
	fmt.Println(courses)

	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	s1 := data[1:6]
	s2 := data[2:7]
	fmt.Println(len(s1), cap(s1))
	fmt.Println(len(s2), cap(s2))
	s2 = append(s2, 1, 2, 3, 4, 5, 6, 7, 8, 9) // s2有append，s2扩容，s2修改，s1不会影响
	s2[0] = 22                                 // s2没有append，s2修改，s1也被改到
	fmt.Println(s2)
	fmt.Println(s1)

	var dataSlice []int
	for i := 0; i < 2000; i++ {
		dataSlice = append(dataSlice, i)
		fmt.Printf("len: %d, cap: %d\n", len(dataSlice), cap(dataSlice))
	}
}
