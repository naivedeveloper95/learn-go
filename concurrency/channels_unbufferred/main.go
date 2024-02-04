package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan int) // unbufferred channel

	c2 := make(chan int, 3) // bufferred channel
	_ = c2

	go func(c chan int) {
		fmt.Println("func goroutine starts sending data into the channel")
		c <- 10
		fmt.Println("func goroutine ends sending data into the channel")
	}(c1)

	fmt.Println("main goroutine sleeps for 2 seconds")
	time.Sleep(time.Second * 2)

	fmt.Println("main starts receiving the data")
	d := <-c1
	fmt.Println("main goroutine received data: ", d)
	time.Sleep(time.Second * 2)
}
