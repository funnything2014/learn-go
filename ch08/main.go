package main

import "fmt"

type Person struct {
	name string
}

func (p *Person) SayHello() {
}

func swap(a, b *int) {
	//a, b = b, a
	t := *a
	*a = *b
	*b = t
}

func changeName(p *Person) {
	p.name = "imooc"
}

func main() {
	// 指针，结构体传值的时候，在函数中修改的值可以反映到变量中
	p := Person{name: "bobby"}
	var pi *Person = &p
	changeName(&p)
	fmt.Println(p.name)
	fmt.Printf("%p\n", pi)

	// 指针的定义
	//var po *Person
	//po = &p
	po := &Person{name: "bobby2"}
	(*po).name = "bobby3"
	po.name = "bobby4"
	// go语言限制了指针的运算，不能参加运算
	// unsafe包里面，一般不会使用，但是当用到的时候可以使用
	fmt.Println(po)

	//var a int = 10
	//b := &a

	// 指针需要初始化
	//var b *int
	//fmt.Println(b)

	//ps := &Person{} // 第一种初始化方式
	//var emptyPerson Person
	//pi := &emptyPerson // 第二种初始化方式

	// map、channel、slice初始化推荐使用make方法
	// 指针初始化推荐使用new函数，否则会出现nil pointer
	// map必须初始化
	var pp = new(Person) // 第三种初始化方式
	fmt.Println(pp.name)

	a, b := 1, 2
	swap(&a, &b)
	fmt.Println(a, b)
}
