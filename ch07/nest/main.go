package main

import "fmt"

type Person struct {
	name string
	age  int
}

// 接收器有两种形态
func (p Person) print() {
	p.age = 19
	fmt.Printf("%s is %d years old\r\n", p.name, p.age)
}

func (p *Person) print2() {
	// 有可能该函数中想要改变p的值，有可能person对象很大，数据较大的时候考虑使用指针方式
	p.age = 19
	fmt.Printf("%s is %d years old\r\n", p.name, p.age)
}

type Student struct {
	//// 第一种嵌套方式
	//p Person

	// 第二种嵌套方式，匿名嵌套
	Person
	score float32
	name  string
}

func main() {
	//s := Student{p: Person{"James", 20}, score: 25}
	//
	//fmt.Println(s.p.age)

	s := Student{
		Person{
			name: "James",
			age:  32,
		},
		42,
		"bobby",
	}

	fmt.Println(s.name)

	p := Person{
		name: "bobby",
		age:  18,
	}
	p.print()
	fmt.Println(p.age)

	p.print2()
	fmt.Println(p.age)
}
