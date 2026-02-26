package main

import "fmt"

func mprint(datas ...interface{}) {
	for _, value := range datas {
		fmt.Println(value)
	}
}

func mprint2(data interface{}) {
	fmt.Println(data)
}
func main() {
	var data = []interface{}{
		"bobby", 18, 178,
	}
	mprint(data...)
	//slice类型的值不能放进去
	var data1 = []string{
		"bobby1", "bobby2", "bobby3",
	}
	var datai []interface{}
	for _, value := range data1 {
		datai = append(datai, value)
	}
	mprint(datai...)
}
