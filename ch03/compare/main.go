package main

import (
	"fmt"
	"strings"
)

func main() {
	// 字符串的比较
	a := "hello"
	b := "bello"
	fmt.Println(a == b)

	// 字符串的大小比较
	fmt.Println(a > b)

	name := "imooc体系课-go工程师"
	fmt.Println(strings.Contains(name, "go"))

	fmt.Println(strings.Count(name, "o"))

	fmt.Println(strings.Split(name, "-"))

	fmt.Println(strings.HasPrefix(name, "imooc"))

	fmt.Println(strings.HasSuffix(name, "工程师"))

	fmt.Println(strings.Index(name, "go"))

	fmt.Println(strings.Replace(name, "go", "java", -1))

	fmt.Println(strings.ToLower("GO"))

	fmt.Println(strings.ToUpper("java"))

	fmt.Println(strings.Trim("#$hello #go#", "#$"))
}
