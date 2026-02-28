package main

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func TestAdd3(t *testing.T) {
	var dataSet = []struct {
		a   int
		b   int
		out int
	}{
		{1, 1, 2},
		{12, 12, 24},
		{-9, 10, 11},
	}

	for _, value := range dataSet {
		re := add(value.a, value.b)
		if re != value.out {
			t.Errorf("expect: %d,actual:%d", value.out, re)
		}

	}
}

/*
goos: windows
goarch: amd64
pkg: awesomeProject/ch11/test
cpu: 11th Gen Intel(R) Core(TM) i5-11300H @ 3.10GHz
BenchmarkAdd
BenchmarkAdd-8          1000000000               0.2615 ns/op
PASS
*/
func BenchmarkAdd(bb *testing.B) {
	var a, b, c int
	a = 123
	b = 456
	c = 579
	for i := 0; i < bb.N; i++ {
		if actual := add(a, b); actual != c {
			fmt.Printf("%d + %d,expect:%d,actual:%d", a, b, c, actual)
		}
	}
}

const numbers = 10000

func BenchmarkStringSprinft(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var str string
		for j := 0; j < numbers; j++ {
			fmt.Sprintf("%s%d", str, j)
		}
	}
	b.ResetTimer()
}

func BechmarkStringAdd(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var str string
		for j := 0; j < numbers; j++ {
			str = str + strconv.Itoa(j)
		}
	}
	b.ResetTimer()
}

func BenchmarkStringBulder(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var builder strings.Builder
		for j := 0; j < numbers; j++ {
			builder.WriteString(strconv.Itoa(j))
		}
		_ = builder.String()
	}
	b.ResetTimer()
}
