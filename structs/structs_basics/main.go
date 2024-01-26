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

	diana := struct {
		firstName string
		lastName  string
		age       int
	}{
		firstName: "Diana",
		lastName:  "Aligheri",
		age:       25,
	}

	fmt.Printf("%#v\n", diana)
	fmt.Printf("%d\n", diana.age)
	fmt.Printf("%s\n", diana.firstName)
	fmt.Printf("%s\n", diana.lastName)

	type Book struct {
		string
		float64
		bool
	}

	b1 := Book{"1984 By Goerge Orwell", 10.2, false}
	fmt.Printf("%#v\n", b1)

	fmt.Println(b1.string)

	type Employee struct {
		name   string
		salary int64
		bool
	}

	e := Employee{"John", 50000, false}
	fmt.Printf("%#v\n", e)
	e.bool = true
	fmt.Printf("%#v\n", e)
}
