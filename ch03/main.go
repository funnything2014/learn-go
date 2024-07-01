package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// 转义符
	//courseName := "go\"体系课\""
	courseName := `go"体系课"`
	fmt.Println(courseName)

	// 格式化输出
	username := "bobby"
	age := 18
	fmt.Printf("用户名：%s, 年龄%d\n", username, age) // 这个非常的常用，但是性能较差
	userMsg := fmt.Sprintf("用户名：%s, 年龄%d\n", username, age)
	fmt.Println(userMsg)

	// 通过string的builder进行字符串拼接，高性能
	var builder strings.Builder
	builder.WriteString("用户名：")
	builder.WriteString(username)
	builder.WriteString(",年龄：")
	builder.WriteString(strconv.Itoa(age))
	fmt.Println(builder.String())

}
