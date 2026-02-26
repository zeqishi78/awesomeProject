package main

import (
	"fmt"
	"strconv"
)

func main() {
	//int类型转换
	var a int8 = 12
	var b = uint8(a)

	var f float32 = 3.14
	var c = int32(f)
	fmt.Println(b, c)

	var f64 = float64(a)
	fmt.Println(f64)

	type IT int
	var e IT = 13

	var d IT = IT(a) //类型要求很严格
	fmt.Println(e, d)

	//字符串转数字
	var istr = "a"
	myInt, err := strconv.Atoi(istr)
	if err != nil {
		fmt.Println("conver error")
	}
	fmt.Println(myInt)

	var myi = 32
	mystr := strconv.Itoa(myi)
	fmt.Println(mystr)

	//字符串转换为float32，转换为bool
	myf, err := strconv.ParseFloat("3.141592654", 64)
	if err != nil {
		fmt.Println("转换出错")
	}
	fmt.Println(myf)
	parseInt, err := strconv.ParseInt("-42", 10, 64)
	if err != nil {
		return
	}
	fmt.Println(parseInt)

	i, err := strconv.ParseInt("12", 8, 64)
	if err != nil {
		return
	}
	fmt.Println(i)
	//bool类型的0值是false，所以下面会打印false
	parseBool, err := strconv.ParseBool("true")
	if err != nil {
		fmt.Println("parseBool error")
		return
	}
	fmt.Println(parseBool)

	//基本数据类型，转换为字符串类型
	//parse，是将字符串转换为int,反过来，将int转换为format
	formatBool := strconv.FormatBool(true)
	fmt.Println(formatBool)

	floatStr := strconv.FormatFloat(3.141592654, 'E', -1, 64) //float类型转换为string类型
	fmt.Println(floatStr)

	formatInt := strconv.FormatInt(-42, 16)
	fmt.Println(formatInt)
}
