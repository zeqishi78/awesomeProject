package main

import "fmt"

/*
一个接口实现多个结构体  多接口实现
*/
type MyWriter interface {
	Write(string) error
}

type MyCloser interface {
	Close() error
}

type WriterAndCloser struct {
}

type WriterCloser1 struct {
	MyWriter //interface也是一种类型，想要放入一个写文件的实现，我想放入一个写数据库的实现
}

func (wc *WriterAndCloser) Write(string) error {
	fmt.Println("write string")
	return nil
}

func (wc *WriterAndCloser) Close() error {
	fmt.Println("close string")
	return nil
}

func main() {
	var w1 MyWriter = &WriterAndCloser{}
	w1.Write("张三")
	var w2 MyCloser = &WriterAndCloser{}
	w2.Close()
}
