package main

import "fmt"

func main() {
	var employee map[string]string
	fmt.Printf("%#v\n", employee)

	fmt.Printf("No of pairs: %d\n", len(employee))

	fmt.Printf("The value of key %q is %q", "Dan", employee["Dan"])

	var accounts map[string]float64
	fmt.Printf("%#v\n", accounts["0x323"])

	var m1 map[[5]int]string
	_ = m1

	// employee["Dan"] = "Programmer"

	people := map[string]float64{}
	people["John"] = 21.4
	people["Mary"] = 10

	fmt.Println(people)

	map1 := make(map[int]int)
	map1[4] = 8

	balances := map[string]float64{
		"USD": 323.11,
		"EUR": 267.34,
		// 50:    322.15,
		"CHF": 3745.23,
	}
	fmt.Println(balances)

	m := map[int]int{1: 10, 2: 20, 3: 30}
	_ = m

	balances["USD"] = 500.2
	balances["INR"] = 100

	fmt.Println(balances)

	balances["GBP"] = 0

	fmt.Println(balances["GBP"])

	v, ok := balances["GBP"]

	if ok {
		fmt.Println("The GBP value is: ", v)
	} else {
		fmt.Println("The GBP key doesn't exists in the map")
	}

	for k, v := range balances {
		fmt.Printf("Key: %#v, Value: %#v\n", k, v)
	}

	delete(balances, "GBP")
	fmt.Println(balances)
}
