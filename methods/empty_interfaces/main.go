package main

import "fmt"

type emptyInterface struct {
}

type person struct {
	info interface{}
}

func main() {
	var empty interface{}

	empty = 5
	fmt.Println(empty)

	empty = "go"
	fmt.Println(empty)

	empty = []int{1, 2, 5}
	fmt.Println(empty)
	fmt.Println(len(empty.([]int)))

	you := person{}
	you.info = "satyam"
	fmt.Println(you.info)

	you.info = 40
	fmt.Println(you.info)

	you.info = []float64{0.4, 1., 2.4, 3., 4.1, 5.9}
	fmt.Println(you.info)
}
