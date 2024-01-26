package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	s1 := "Golang"
	fmt.Println(len(s1))

	name := "Codruta"
	fmt.Println(len(name))

	n := utf8.RuneCountInString(name)
	m := utf8.RuneCount([]byte(name))

	fmt.Println(n, m)
}
