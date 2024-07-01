package main

import "fmt"

func main() {
	// if条件判断
	age := 16
	country := "中国"
	if age < 18 && country == "中国" {
		fmt.Println("未成年")
	}

}
