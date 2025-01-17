package main

import "fmt"

// 接口的定义
type Duck interface {
	// 方法的声明
	Gaga()
	Walk()
	Swimming()
}

type pskDuck struct {
	legs int
}

func (p *pskDuck) Gaga() {
	fmt.Println("pskDuck GAGA")
}

func (p *pskDuck) Walk() {
	fmt.Println("pskDuck Walk")
}

func (p *pskDuck) Swimming() {
	fmt.Println("pskDuck Swimming")
}

func main() {
	// go语言的接口，鸭子类型，php、python
	// go语言中处处都是interface，到处都是鸭子类型 duck typing
	// 鸭子类型强调的是事物的外部行为，而不是内部的结构

	var d Duck = &pskDuck{}
	d.Walk()
}
