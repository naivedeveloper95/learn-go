package main

import (
	"fmt"
	"runtime"
	"time"
)

func f1() {
	fmt.Println("f1 (goroutine) execution started")

	for i := 0; i < 3; i++ {
		fmt.Println("f1: ", i)
	}

	fmt.Println("f1 execution finished")
}

func f2() {
	fmt.Println("f2 (goroutine) execution started")

	for i := 0; i < 8; i++ {
		fmt.Println("f2: ", i)
	}

	fmt.Println("f2 execution finished")
}

func main() {
	fmt.Println("Main execution started")
	fmt.Println("No. of CPU's: ", runtime.NumCPU())
	fmt.Println("No. of goroutines: ", runtime.NumGoroutine())

	fmt.Println("OS: ", runtime.GOOS)
	fmt.Println("Arch: ", runtime.GOARCH)

	fmt.Println("GOMAXPROCS: ", runtime.GOMAXPROCS(0))

	go f1()
	fmt.Println("No. of goroutines after go f1(): ", runtime.NumGoroutine())

	f2()

	time.Sleep(time.Second * 2)
	fmt.Println("main execution stopped")

}
