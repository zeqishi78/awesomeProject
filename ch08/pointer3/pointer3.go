package main

import "fmt"

func main() {
	var x, y int = 100, 500
	fmt.Printf("x = %d,y = %d\r\n", x, y)
	swap(&x, &y)
	fmt.Printf("x = %d,y = %d\r\n", x, y)
}

func swap(a, b *int) {
	t := *a
	*a = *b
	*b = t
}
