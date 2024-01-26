package main

import "fmt"

func main() {
	var age int = 10
	fmt.Println("Age: ", age)

	var name = "Satyam"
	fmt.Println("Your name is : ", name)

	s := "Learning Golang"
	fmt.Println(s)

	car, cost := "Audi", 60000
	fmt.Println(car, cost)

	car, year := "BMW", 2024
	_ = year

	var opened = false
	opened, file := true, "file.txt"

	_, _ = opened, file

	var (
		salary    float64
		firstName string
		gender    bool
	)

	fmt.Println(salary, firstName, gender)

	var a, b, c int
	fmt.Println(a, b, c)

	var i, j int
	i, j = 5, 8

	j, i = i, j
	fmt.Println(i, j)

	sum := 5 + 2.3
	fmt.Println(sum)

}
