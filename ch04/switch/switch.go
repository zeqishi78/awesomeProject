package main

import "fmt"

func main() {
	/*
		switch var1{
			case val1:
				...
			case val2:
				...
			case val3:
				...
			default:
				...
		}
	*/
	day := "星期三"
	switch day {
	case "星期一":
		fmt.Println("Monday")
	case "星期二":
		fmt.Println("Tuesday")
	case "星期三":
		fmt.Println("Wednesday")
	default:
		fmt.Println("unknow")
	}
}
