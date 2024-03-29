package main

import (
	"fmt"
	"math"
)

func f1() {
	fmt.Println("This is f1() function")
}

func f3(a, b, c int, d, e float64, s string) {
	fmt.Println(a, b, c, d, e, s)
}

func f2(a int, b int) {
	fmt.Println("Sum: ", a+b)
}

func f4(a float64) float64 {
	return math.Pow(a, a)
}

func f5(a, b int) (int, int) {
	return a + b, a * b
}

func sum(a, b int) (s int) {
	fmt.Println(s)
	s = a + b
	return
}

func main() {
	f1()
	f2(7, 9)
	f3(4, 5, 6, 5.4, 6.4, "hello")
	fmt.Println(f4(4))
	a, b := f5(4, 9)
	fmt.Println(a, b)
	fmt.Println(sum(4, 8))
}
