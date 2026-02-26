package main

import (
	"errors"
	"fmt"
)

type MyError struct {
	Code    int
	Message string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("错误代码%d : %s", e.Code, e.Message)
}

func process() error {
	return &MyError{Code: 1001, Message: "处理失败"}
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("除数不能为0...")
	}
	return a / b, nil
}

func main() {
	i1, err := divide(6, 3)
	if err != nil {
		fmt.Println(err)
	} else if err == nil {
		fmt.Println(i1)
	}
	i2, err1 := divide(10, 0)
	if err1 != nil {
		fmt.Println(err1)
	}
	fmt.Println("结果:", i2)
	err2 := process()
	fmt.Println(err2)
}
