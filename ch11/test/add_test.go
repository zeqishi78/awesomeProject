package main

import (
	"fmt"
	"testing"
)

// 功能测试以Test开头，后面再加上函数名
func TestAdd(t *testing.T) {
	res := add(1, 2)
	if res != 3 {
		t.Errorf("expect %d,actual %d\r\n", 3, res)
	}
}

func TestAdd2(t *testing.T) {
	fmt.Println("yes1")
	if testing.Short() {
		t.Skip("short 模式下跳过")
	}
	fmt.Println("yes")
	res := add(1, 5)
	if res != 6 {
		t.Errorf("expect %d,actual %d\r\n", 6, res)
	}
}
