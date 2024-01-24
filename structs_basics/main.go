package main

import "fmt"

func main() {
	title1, author1, year1 := "The Devine Comedy", "Dante Aligheri", 1320

	title2, author2, year2 := "Macbeth", "William Shekesphere", 1606

	fmt.Println("Book1: ", title1, author1, year1)
	fmt.Println("Book2: ", title2, author2, year2)

	type book struct {
		title  string
		author string
		year   int
	}

	type book1 struct {
		title  string
		author string
		year   int
	}

	myBook := book{"The Devine Comedy", "Dante Aligheri", 1320}
	fmt.Println(myBook)

	myBook1 := book1{"Macbeth", "William Shekesphere", 1606}
	fmt.Println(myBook1)

	bestBook := book{title: "Animal Farm", author: "Goerge Orwell", year: 1320}
	_ = bestBook

	aBook := book{title: "Just a Random Book"}
	fmt.Printf("%#v\n", aBook)
}
