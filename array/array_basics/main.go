package main

import "fmt"

func main() {
	var nums [4]int

	fmt.Printf("%v\n", nums)
	fmt.Printf("%#v\n", nums)

	var a1 = [4]float64{}
	fmt.Printf("%#v\n", a1)

	var a2 = [3]int{-10, 1, 100}
	fmt.Printf("%#v\n", a2)

	a3 := [4]string{"Dan", "Diana", "Paul", "John"}
	fmt.Printf("%#v\n", a3)

	a4 := [4]string{"x", "y"}
	fmt.Printf("%#v\n", a4)

	a5 := [...]int{1, 2, 3, 4, 10, 45, 67}
	fmt.Printf("The length of the array a5 is %d\n", len(a5))

	a6 := [...]int{1,
		2,
		3,
		4,
		10,
		45,
		67,
	}
	fmt.Printf("The length of the array a6 is %d\n", len(a6))
}
