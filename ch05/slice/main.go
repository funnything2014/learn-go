package main

import "fmt"

func main() {
	// slice 切片 - 数组
	var courses []string
	fmt.Printf("%T\r\n", courses)

	// 这个方法很特别
	courses = append(courses, "go")
	courses = append(courses, "grpc")
	courses = append(courses, "gin")
	fmt.Println(courses)

	// 切片的初始化3种 1. 从数组直接创建 2. 使用string{} 3. make
	//allCourses := [3]string{"go", "grpc", "http"}

	allCourses := []string{"go", "grpc", "http"}
	courseSlice := allCourses[0:1] // 左闭右开, python
	fmt.Println(courseSlice)

	allCourses2 := make([]string, 3)
	allCourses2[0] = "go"
	allCourses2[1] = "grpc"
	allCourses2[2] = "http"
	fmt.Println(allCourses2)

	var allCourses3 []string
	allCourses3 = append(allCourses3, "gin")
	fmt.Println(allCourses3)

	// 访问切片的元素，访问单个，访问多个
	fmt.Println(allCourses2[1])
	fmt.Println(allCourses2[0:3])

	courseSlice1 := []string{"go", "grpc", "http"}
	courseSlice2 := []string{"mysql", "es", "gin"}
	courseSlice = append(courseSlice1, courseSlice2[1:]...)
	fmt.Println(courseSlice)

	// 如何删除slice的元素，比较麻烦
	courseSlice3 := []string{"go", "grpc", "http", "mysql", "es", "gin"}
	mySlice := append(courseSlice3[:2], courseSlice3[3:]...)
	fmt.Println(mySlice)

	// 复制slice
	courseSliceCopy := courseSlice
	courseSliceCopy2 := courseSlice[:]
	fmt.Println(courseSliceCopy, courseSliceCopy2)

	var courseSliceCopy3 = make([]string, len(courseSlice))
	copy(courseSliceCopy3, courseSlice)
	fmt.Println(courseSliceCopy3)

	fmt.Println("----------------------------")
	courseSlice[0] = "java"
	fmt.Println(courseSliceCopy2)
	fmt.Println(courseSliceCopy3)
}
