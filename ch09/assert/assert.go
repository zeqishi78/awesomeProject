package main

import "fmt"

func add(a, b interface{}) int {
	ai, ok := a.(int)
	if !ok {
		fmt.Println("not an int type...")
	}
	bi, _ := b.(int)
	return ai + bi
}

func main() {
	a := 1.3
	b := 2.3
	fmt.Println(add(a, b))

}
