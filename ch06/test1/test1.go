package main

import (
	"fmt"
)

func main() {
	//1
	//repeat(func() {
	//	fmt.Println("调用了")
	//}, 3)
	//2.
	s := []int{1, 2, 3, 4, 5, 6}
	//newSlice := forEach(s, func(i int) int {
	//	if i <= 3 {
	//		fmt.Println("========>当前值：", i)
	//		i += 3
	//	}
	//	fmt.Println("========>更正值：", i)
	//	return i
	//})
	//printSlice(newSlice)
	//3.
	filters := myFilter(s, func(i int) (int, bool) {
		if i < 3 {
			return i, true
		}
		return 0, false
	})
	printSlice(filters)
	//fmt.Printf("类型：%T\r\n", filters)
	//println(filters)
}

func printSlice(s []int) {
	for _, value := range s {
		fmt.Println(value)
	}
}

// 写一个高阶函数repeat，它接受一个函数f和数字n，调用f函数n次
func repeat(f func(), n int) {
	for i := 1; i <= n; i++ {
		f()
	}
}

// 写一个ForEach函数，接收一个[]int切片和一个func(int)函数，对每个元素执行函数
func forEach(slices []int, f func(i int) int) []int {
	newSlice := []int{}
	for m := 0; m < len(slices); m++ {
		newSlice = append(newSlice, f(slices[m]))
	}
	return newSlice
}

// 3.实现Filter函数，接收[]int和一个func(int) bool判断函数，
// 返回符合条件的元素切片
func myFilter(slicePara []int, filter func(i int) (int, bool)) []int {
	slice := []int{}
	for _, v := range slicePara {
		i, bool := filter(v)
		if bool {
			slice = append(slice, i)
		}
	}
	return slice
}

// 4.设计 SortBy函数，接收任意类型的切片和一个比较函数 func(a, b interface{}) bool，根据比较函数排序切片。
//func sortBy(slice []int, f func(a, b interface{}) bool) []int {
//	v := reflect.ValueOf(slice)
//	if !v.IsValid() {
//		return nil
//	}
//	if v.Kind() == reflect.Ptr {
//		v = v.Elem()
//	}
//	if v.Kind() != reflect.Slice {
//		return nil
//	}
//	sort.SliceStable(v.Interface(), func(i, j int) bool {
//		return less(v.Index(i).Interface(), v.Index(j).Interface())
//	})
//	return nil
//
//}
