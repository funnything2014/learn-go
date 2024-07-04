package main

import (
	"fmt"
	"strings"
)

func add(a, b interface{}) interface{} {
	switch a.(type) {
	case int:
		return a.(int) + b.(int)
	case int32:
		return a.(int32) + b.(int32)
	case float64:
		return a.(float64) + b.(float64)
	case string:
		return a.(string) + b.(string)
	default:
		panic("not supported type")
	}
	//ai, ok := a.(int)
	//if !ok {
	//	panic("not int type")
	//}
	//bi, _ := a.(int)
	//return ai + bi
}

func main() {
	a := 1.0
	b := 2.0
	fmt.Println(add(a, b))
	re := add("hello ", "bobby")
	res, _ := re.(string)
	fmt.Println(strings.Split(res, ""))
	fmt.Println(res)
}
