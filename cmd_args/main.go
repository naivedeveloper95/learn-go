package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("os.Args", os.Args)
	fmt.Println("Path: ", os.Args[0])
	fmt.Println("1st Args: ", os.Args[1])
	fmt.Println("2nd Args: ", os.Args[2])
	fmt.Println("Total number of args: ", len(os.Args))

	var result, err = strconv.ParseFloat(os.Args[1], 64)
	_ = err
	fmt.Printf("%T\n", os.Args[1])
	fmt.Printf("%T\n", result)
	fmt.Println(result)
}
