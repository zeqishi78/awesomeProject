package main

import "fmt"

/*
定义了一个User结构体类型，包含两个字段:Name(字符串类型)和Age(整型)
*/
type User struct {
	Name string
	Age  int
}

func main() {
	var user *User  //声明一个指向User类型的指针变量，初始值为nil
	printUser(user) //输出:用户不存在
	//为user指针赋值，指向新创建的User实例
	user = &User{
		Name: "张三",
		Age:  23,
	}
	printUser(user)
}
func printUser(user *User) {
	if user == nil {
		fmt.Println("用户不存在...")
		return
	}
	fmt.Printf("用户:%s,年龄:%d\r\n", user.Name, user.Age)
}
