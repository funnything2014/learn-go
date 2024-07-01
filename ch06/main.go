package main

import "fmt"

// go语言中函数参数传递全部都是值传递
//func add(a int, b int) (int, error) {
//	return a + b, nil
//}

//func add(a int, b int) (sum int, err error) {
//	a = 3
//	sum = a + b
//	//return sum, err
//	return
//}

func add(items ...int) (sum int, err error) {
	for _, value := range items {
		sum += value
	}
	return
}

func cal(op string) func() {
	switch op {
	case "+":
		return func() {
			fmt.Println("这是加法")
		}
	case "-":
		return func() {
			fmt.Println("这是减法")
		}
	default:
		return func() {
			fmt.Println("不是加法也不是减法")
		}
	}
}

func callback(i int, f func(a int, b int)) {
	f(i, 2)
}

//func autoIncrement() int {
//	local := 0 // 一个函数中访问另一个函数的局部变量，闭包
//	return func() int {
//		local += 1
//		return local
//	}()
//}

func autoIncrement() func() int {
	local := 0 // 一个函数中访问另一个函数的局部变量，闭包
	return func() int {
		local += 1
		return local
	}
}

func main() {
	// go函数支持普通函数、匿名函数、闭包
	/**
	go中函数是 一等公民
	1. 函数本身可以当做变量
	2. 匿名函数 闭包
	3. 函数可以满足接口
	*/

	funcVar := add

	a := 1
	b := 2
	sum, _ := add(a, b, 3, 4)
	fmt.Println(sum)
	fmt.Println(a)
	sum1, _ := funcVar(a, b, 3, 4)
	fmt.Println(sum1)

	cal("+")()

	// 匿名函数
	localFunc := func(a, b int) {
		fmt.Printf("total is %d\r\n", a+b)
	}

	callback(1, localFunc)

	// 闭包
	//for i := 0; i < 5; i++ {
	//	fmt.Println(autoIncrement())
	//}

	nextFunc := autoIncrement()
	for i := 0; i < 10; i++ {
		fmt.Println(nextFunc())
	}

	nextFunc2 := autoIncrement() //归零
	for i := 0; i < 3; i++ {
		fmt.Println(nextFunc2())
	}
}
