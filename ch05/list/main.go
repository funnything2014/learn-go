package main

import (
	"container/list"
	"fmt"
)

func main() {
	//var myList list.List
	myList := list.New()
	// 放在尾部
	myList.PushBack("go")
	myList.PushBack("grpc")
	myList.PushBack("mysql")

	// 头部放数据
	myList.PushFront("gin")
	e := myList.Front()
	for ; e != nil; e = e.Next() {
		if e.Value == "grpc" {
			break
		}
	}
	myList.InsertBefore("bobby", e)

	myList.Remove(e)

	fmt.Println(myList)

	// 遍历打印值，正序
	for e := myList.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

	// 反向遍历
	for e := myList.Back(); e != nil; e = e.Prev() {
		fmt.Println(e.Value)
	}

	// 集合类型4种 1.数组-不同长度的数组类型不一样 2.切片-动态数组，方便，性能高 3.map 4.list-用的少
}
