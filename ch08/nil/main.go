package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	/**
	不同类型的数据零值不一样
	bool false
	numbers 0
	string ""
	pointer nil
	slice nil
	channel、interface、function nil

	struct默认值不是nil、默认值是具体字段的默认值
	*/
	p1 := Person{
		name: "bobby",
		age:  23,
	}

	p2 := Person{
		name: "bobby",
		age:  23,
	}

	if p1 == p2 {
		fmt.Println("p1 and p2 are equal")
	}

	// slice的默认值，本质是struct {ptr *element, len int, cap int}
	var ps []Person // nil slice
	//var ps = make([]Person, 0) // empty slice
	if ps == nil {
		fmt.Println("ps is nil")
	}

	var m map[string]string             // nil map
	var m2 = make(map[string]string, 0) // empty map

	for key, value := range m {
		fmt.Println(key, value)
	}

	for key, value := range m2 {
		fmt.Println(key, value)
	}

	fmt.Println(m["name"])
	//m["name"] = "bobby"
	m2["name"] = "bobby"

	if m2 == nil {
		fmt.Println("m2 is nil")
	}
}
