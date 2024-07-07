package ch11

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func TestAdd(t *testing.T) {
	//re := Add(1, 3)
	//if re != 4 {
	//	t.Errorf("Add(1, 3) => %d, want %d", re, 4)
	//}
	var dataset = []struct {
		a   int
		b   int
		out int
	}{
		{1, 2, 3},
		{12, 12, 24},
		{-9, 8, -1},
		{0, 0, 0},
	}

	for _, d := range dataset {
		result := Add(d.a, d.b)
		if result != d.out {
			t.Errorf("Add(%d, %d) = %d, want %d", d.a, d.b, result, d.out)
		}
	}
}

// go test -short
func TestAdd2(t *testing.T) {
	fmt.Println("yes1")
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
	fmt.Println("yes")
	re := Add(1, 5)
	if re != 6 {
		t.Errorf("Add(1, 5) => %d, want %d", re, 6)
	}
}

func BenchmarkAdd(bb *testing.B) {
	var a, b, c int
	a = 123
	b = 456
	c = 579

	for i := 0; i < bb.N; i++ {
		if actual := Add(a, b); actual != c {
			fmt.Printf("Add(%d, %d) = %d, want %d", a, b, actual, c)
		}
	}
}

// go test -bench=".*"
const numbers = 10000

func BenchmarkStringSprintf(bb *testing.B) {
	bb.ResetTimer()

	for i := 0; i < bb.N; i++ {
		var str string
		for j := 0; j < numbers; j++ {
			str = fmt.Sprintf("%s%d", str, j)
		}
	}

	bb.StopTimer()
}

func BenchmarkStringAdd(bb *testing.B) {
	bb.ResetTimer()

	for i := 0; i < bb.N; i++ {
		var str string
		for j := 0; j < numbers; j++ {
			str = str + strconv.Itoa(j)
		}
	}

	bb.StopTimer()
}

func BenchmarkStringBuilder(bb *testing.B) {
	bb.ResetTimer()

	for i := 0; i < bb.N; i++ {
		var builder strings.Builder
		for j := 0; j < numbers; j++ {
			builder.WriteString(strconv.Itoa(j))
		}
		_ = builder.String()
	}

	bb.StopTimer()
}
