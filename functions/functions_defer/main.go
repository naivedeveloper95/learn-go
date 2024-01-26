package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	defer foo()
	bar()

	fmt.Println("Just a string after defering foo() and calling bar()")
	defer fooBar()

	file, err := os.Open("main.go")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
}

func foo() {
	fmt.Println("This is Foo()")
}

func bar() {
	fmt.Println("This is bar()")
}

func fooBar() {
	fmt.Println("This is fooBar()")
}
