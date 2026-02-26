package main

import (
	"fmt"
	"strings"
)

// 做到通用
func add(a, b interface{}) interface{} {
	switch a.(type) {
	case int:
		ai, _ := a.(int)
		bi, _ := b.(int)
		return ai + bi
	case int32:
		ai, _ := a.(int32)
		bi, _ := b.(int32)
		return ai + bi
	case int64:
		ai, _ := a.(int64)
		bi, _ := b.(int64)
		return ai + bi
	case string:
		as, _ := a.(string)
		bs, _ := b.(string)
		return as + bs
	default:
		panic("not supported type...")
	}
}
func main() {
	fmt.Println(add("hello", "bobby"))
	str := add("hello", "tom")
	s := str.(string)
	split := strings.Split(s, "o")
	fmt.Println(split)
}
