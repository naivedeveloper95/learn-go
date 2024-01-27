package main

import "fmt"

func main() {
	name := "Satyam"
	fmt.Println(&name)

	var x int = 2
	ptr := &x

	fmt.Printf("Ptr is of type %T with a value %v and address %p\n", ptr, ptr, &ptr)
	fmt.Printf("Address of x is %p\n", &x)

	var ptr1 *float64

	_ = ptr1

	p := new(int)
	x = 100
	p = &x

	fmt.Printf("p is of type %T with a value %v\n", p, p)
	fmt.Printf("Address of x is %p\n", &x)

	*p = x // equivalent to x = 90
	fmt.Println(x, *p)
	fmt.Println("*p==x: ", *p == x)

	*p = 10     // equivalent to x = 10
	*p = *p / 2 // equivalent to x = x / 2
	fmt.Println(x)
}
