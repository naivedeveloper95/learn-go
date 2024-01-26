package main

import "fmt"

func main() {
	friends := map[string]int{"Dan": 40, "Maria": 25}
	neighbours := friends

	friends["Dan"] = 50
	fmt.Println(neighbours)

	people := make(map[string]int)

	for key, value := range friends {
		people[key] = value
	}

	people["Anne"] = 19
	fmt.Printf("%#v\n", people)
	fmt.Printf("%#v\n", friends)

	delete(friends, "Dan")
	fmt.Println(friends)
	fmt.Println(people)
}
