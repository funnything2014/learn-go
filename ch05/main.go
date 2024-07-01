package main

import "fmt"

func main() {
	// 集合类型的数据结构：数组、切片、map、list

	// 数组
	var courses1 [3]string // course类型，只有三个元素的数组
	var courses2 [4]string

	// []string 和 [3]string 是两种不同的类型，切片和数组
	fmt.Printf("%T\r\n", courses1)
	fmt.Printf("%T\r\n", courses2)

	// 数组的初始化 1
	var courses3 [3]string = [3]string{"go", "grpc", "gin"}
	for _, course := range courses3 {
		fmt.Println(course)
	}

	// 数组的初始化 2
	courses4 := [3]string{2: "gin"}
	for _, course := range courses4 {
		fmt.Println(course)
	}

	// 数组的初始化 3
	courses5 := [...]string{"go", "grpc"}
	for _, course := range courses5 {
		fmt.Println(course)
	}

	// 多维数组
	var courseInfo [3][4]string
	courseInfo[0] = [4]string{"go", "1h", "hobby", "go体系课"}
	courseInfo[1] = [4]string{"grpc", "2h", "hobby2", "grpc入门"}
	courseInfo[2] = [4]string{"gin", "1.5h", "hobby3", "gin高级开发"}

	for _, row := range courseInfo {
		fmt.Println(row)
	}
}
