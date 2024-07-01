package main

import (
	"fmt"
	"time"
)

func main() {
	// for循环
	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}

	var i int
	for i < 3 {
		time.Sleep(2 * time.Second)
		fmt.Println(i)
		i++
	}

	// for range，主要是对 字符串、数组、切片、map、channel
	name := "imooc go"
	for index, value := range name {
		fmt.Println(index, value)
	}

	/**
	for range key, value
	字符串， key：索引， value：对应字符值的拷贝，不写key，返回的是索引
	数组、切片类似
	map，key：map的key，value：对应值的拷贝，不写key，返回map的值
	channel，value：返回channel接受的数据
	*/
	chName := "imooc go体系课"
	chNameRune := []rune(chName)
	for index, value := range chNameRune {
		fmt.Printf("%d, %c\r\n", index, value)
	}
}
