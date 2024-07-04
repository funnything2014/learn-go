package main

import "fmt"

func mPrint(data ...interface{}) {
	for _, value := range data {
		fmt.Println(value)
	}
}

type myinfo struct {
}

func (m *myinfo) Error() string {
	return "myinfo"
}

func main() {
	//var data = []interface{}{
	//	"bobby", 18, 1.80,
	//}
	//mPrint(data...)

	var data = []string{
		"bobby", "bobby2", "bobby3",
	}

	var datai []interface{}
	for _, value := range data {
		datai = append(datai, value)
	}

	mPrint(datai...)

	var err error
	err = &myinfo{}
	fmt.Println(err.Error())
}
