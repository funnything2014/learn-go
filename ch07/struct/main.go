package main

import "fmt"

type Person struct {
	name    string
	age     int
	address string
	height  float32
}

func main() {
	// 类型集合

	// 如何初始化结构体
	p1 := Person{"bobby1", 18, "慕课网", 1.80}

	p2 := Person{
		name:    "bobby2",
		age:     18,
		address: "北京市",
		height:  1.78,
	}

	var persons []Person
	persons = append(persons, p1, p2)
	persons = append(persons, Person{
		name: "bobby3",
	})

	persons = []Person{
		{"bobby1", 18, "慕课网", 1.80},
		{
			name:    "bobby2",
			age:     18,
			address: "北京市",
			height:  1.78,
		},
	}

	var p Person
	p.age = 20

	fmt.Println(p.height)

	// 匿名结构体
	address := struct {
		province string
		city     string
		address  string
	}{
		"北京市",
		"通州区",
		"xxx",
	}

	fmt.Println(address.city)
}
