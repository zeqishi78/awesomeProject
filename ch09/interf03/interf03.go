package main

import "fmt"

// 可读接口
type Readable interface {
	Read() string
}

// 可写接口
type Writable interface {
	Write(data string)
}

// 文件结构体
type File struct {
	Content string
}

func (f *File) Read() string {
	return f.Content
}

func (f *File) Write(data string) {
	f.Content = data
	fmt.Printf("写入内容:%s\r\n", data)
}

func main() {
	file := &File{
		Content: "hello world",
	}
	file.Write("张三")
	fmt.Println(file.Read())
}
