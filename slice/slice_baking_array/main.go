package main

import (
	"fmt"
	"unsafe"
)

func main() {
	s1 := []int{10, 20, 30, 40, 50}
	s3, s4 := s1[0:2], s1[1:3]

	s1[1] = 600
	_ = s3
	fmt.Println(s1)
	fmt.Println(s4)

	arr1 := [4]int{10, 20, 30, 40}
	slice1, slice2 := arr1[0:2], arr1[1:3]

	arr1[1] = 2
	fmt.Println(arr1, slice1, slice2)

	cars := []string{"Ford", "Honda", "Audi", "Range Rover"}
	newCars := []string{}

	newCars = append(newCars, cars[0:2]...)

	cars[0] = "Nissan"
	fmt.Println(cars, newCars)

	s1n := []int{10, 20, 30, 40, 50}
	newSliceN := s1n[0:3]
	fmt.Println(len(newSliceN), cap(newSliceN))

	newSliceN = s1n[2:5]
	fmt.Println(len(newSliceN), cap(newSliceN))

	fmt.Printf("%p\n", s1n)

	fmt.Printf("%p %p \n", &s1n, &newSliceN)

	newSliceN[0] = 1000
	fmt.Println("s1n: ", s1n)

	a := [5]int{1, 2, 3, 4, 5}
	s := []int{1, 2, 3, 4, 5}

	fmt.Printf("array's size in bytes: %d \n", unsafe.Sizeof(a))
	fmt.Printf("slice's size in bytes: %d \n", unsafe.Sizeof(s))
}
