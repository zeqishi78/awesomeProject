package main

import "fmt"

/*
	if  布尔表达式{
		逻辑
	}
*/
func main() {
	//if条件判断
	country := "中国"
	age := 16
	if age < 18 {
		if country == "中国" {
			fmt.Println("未成年")
		}
	} else if age == 18 {
		fmt.Println("刚好成年")
	} else {
		fmt.Println("成年")
	}
	/*
		实际应用
	*/
	if age < 18 {
		fmt.Println("未成年")
	}
	if age == 18 {
		fmt.Println("刚好成年")
	}

	if age > 18 {
		fmt.Println("成年")
	}
}
