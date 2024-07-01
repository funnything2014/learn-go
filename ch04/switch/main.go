package main

import "fmt"

func main() {
	score := 60
	switch {
	case score < 60:
		fmt.Println("E")
	case score >= 60 && score < 70:
		fmt.Println("D")
	default:
		fmt.Println("You are good")
	}

}
