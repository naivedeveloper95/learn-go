package main

import (
	"fmt"
)

func main() {
	// INT TYPE
	var i1 int8 = -120
	fmt.Printf("%T\n", i1)

	var i2 uint16 = 65535
	fmt.Printf("%T\n", i2)

	// FLOAT TYPE
	var f1, f2, f3 float64 = 1.1, -0.2, 5.
	fmt.Printf("%T %T %T \n", f1, f2, f3)

	// RUNE TYPE
	var r rune = 'f'
	fmt.Printf("%T \n", r)
	fmt.Println(r)
	fmt.Printf("%x \n", r)

	// BOOL TYPE
	var b bool = true
	fmt.Printf("%T \n", b)

	// STRING TYPE
	var s string = "Hello Go!"
	fmt.Printf("%T \n", s)

	// ARRAY TYPE
	var numbers = [4]int{4, 5, -9, 100}
	fmt.Printf("%T\n", numbers)

	// SLICE TYPE
	var cities = []string{"London", "Tokyo", "New York"}
	fmt.Printf("%T\n", cities)

	// MAP TYPE
	balances := map[string]float64{
		"USD": 2332.2,
		"EUR": 511.11,
	}
	fmt.Printf("%T\n", balances)

	// STRUCT TYPE
	type Person struct {
		name string
		age  int
	}

	var you Person
	fmt.Printf("%T \n", you)

	// POINTER TYPE
	var x int = 2
	ptr := &x
	fmt.Printf("ptr is of type %T with the value of %v\n", ptr, ptr)

	// FUNCTION TYPE
	fmt.Printf("%T\n", f)
}

func f() {

}
