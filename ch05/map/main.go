package main

import (
	"fmt"
)

func main() {
	// map是一个key 索引 和value 值 的无序集合，主要是查询方便
	var courseMap = map[string]string{
		"go":   "go工程师",
		"grpc": "grpc入门",
		"gin":  "gin深入理解",
	}

	fmt.Println(courseMap["grpc"])

	courseMap["mysql"] = "mysql的原理"

	fmt.Println(courseMap)

	//var courseMap2 map[string]string // nil， map类型想要设置值必须要先初始化
	var courseMap2 = make(map[string]string, 3) // make是内置函数，主要用于初始化slice、map、channel
	courseMap2["mysql"] = "mysql的原理"
	fmt.Println(courseMap2)

	// map必须初始化才能使用 1.map[string]string{} 2.make(map[string]string, 3)  但是slice可以不初始化
	var m []string
	if m == nil {
		fmt.Println("nil yes")
	}
	m = append(m, "a")

	// 遍历
	for key, value := range courseMap {
		fmt.Println(key, value)
	}

	for key := range courseMap {
		fmt.Println(courseMap[key])
	}

	// map是无序的，而且不保证每次打印都是相同的顺序

	//d, ok := courseMap2["java"]
	//if !ok {
	//	fmt.Println("not in")
	//} else {
	//	fmt.Println(d)
	//}

	if d, ok := courseMap2["java"]; !ok {
		fmt.Println("not in")
	} else {
		fmt.Println(d)
	}

	// 删除一个元素
	delete(courseMap, "mysql")
	// 很重要的提示，map不是线程安全的
	//var syncMap sync.Map
	fmt.Println(courseMap)
}
