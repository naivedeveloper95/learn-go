package main

import "fmt"

func main() {
	const days int = 7

	var i int
	fmt.Println(i)

	const pi float64 = 3.14
	const secondsInHr = 3600

	duration := 234

	fmt.Printf("Duration in seconds is: %v\n", duration*secondsInHr)

	x, y := 5, 1
	fmt.Println(x / y)

	const a = 5
	const b = 1

	fmt.Println(a / b)

	const n, m int = 4, 5
	const n1, m1 = 6, 7

	const (
		min1 = -500
		min2
		min3
	)

	fmt.Println(min1, min2, min3)

	const l1 = len("hello")
	fmt.Println(l1)

}
