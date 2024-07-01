package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a int8 = 12
	var b = uint8(a)

	var f float32 = 3.14
	var c = int32(f)
	fmt.Println(b, c)

	type IT int
	var d = IT(a)
	fmt.Println(d)

	var istr = "12"
	myint, err := strconv.Atoi(istr)
	if err != nil {
		fmt.Println("convert error")
	}
	fmt.Println(myint)

	var myi = 32
	mstr := strconv.Itoa(myi)
	fmt.Println(mstr)

	// 字符串转换为float32，转换为bool
	float, err := strconv.ParseFloat("3.1415", 64)
	if err != nil {
		fmt.Println("convert error")
	}
	fmt.Println(float)

	parseInt, err := strconv.ParseInt("12", 8, 64)
	if err != nil {
		fmt.Println("convert error")
	}
	fmt.Println(parseInt)

	parseBool, err := strconv.ParseBool("true")
	if err != nil {
		fmt.Println("convert error")
	}
	fmt.Println(parseBool)

	// 基本类型转字符串
	formatBool := strconv.FormatBool(true)
	fmt.Println(formatBool)

	formatFloat := strconv.FormatFloat(3.1415926, 'E', -1, 64)
	fmt.Println(formatFloat)

	formatInt := strconv.FormatInt(42, 16)
	fmt.Println(formatInt)
}
