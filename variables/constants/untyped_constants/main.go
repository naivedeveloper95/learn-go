package main

import "fmt"

func main() {

	const a float64 = 5.1 //typed constants

	const b = 6.7 // untyped constants

	const c float64 = a * b

	fmt.Println(c)

	const str = "Hello" + " Go!"

	fmt.Println(str)

	const d = 5 > 10

	fmt.Println(d)

	const x int = 5
	const y float64 = 2.2 * float64(x)

	fmt.Println(y)

	const k = 5
	const t = 2.2 * 5

	fmt.Printf("%T\n", t)

	var i int = k     // k changes to int
	var j float64 = k // var j float64 = float64(k)

	var p byte = k

	fmt.Println(i, j, p)

}
