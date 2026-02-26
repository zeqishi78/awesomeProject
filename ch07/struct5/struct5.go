package main

import "fmt"

type Employee struct {
	Id      int
	Name    string
	Age     int
	Address string
}

// 字段访问和修改
func main() {
	employee := Employee{
		1,
		"张三",
		23,
		"湖北省武汉市",
	}
	fmt.Println(employee.Name)
	employee.Name = "李四"
	fmt.Println(employee.Name)

	var e Employee
	fmt.Printf("%#v\r\n", e) //main.Employee{Id:0, Name:"", Age:0, Address:""}

	//匿名结构体
	s := struct {
		Username string
		Password string
	}{
		Username: "admin",
		Password: "root1234",
	}
	fmt.Println(s)

}
