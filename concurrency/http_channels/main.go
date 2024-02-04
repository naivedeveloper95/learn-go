package main

import (
	"fmt"
	"net/http"
	"runtime"
	"time"
)

func checkUrl(url string, c chan string) {
	resp, err := http.Get(url)

	if err != nil {

		s := fmt.Sprintf("%s is DOWN!\n", url)
		s += fmt.Sprintf("Error: %v  \n", err)

		// Sending the string over the channel.
		// This is a blocking call, so this goroutine will
		// wait for the main goroutine to receive it on the other part of the channel.
		fmt.Println(s)
		c <- url
	} else {

		s := fmt.Sprintf("Status Code: %d  \n", resp.StatusCode)
		s += fmt.Sprintf("%s is UP\n", url)

		// sending s over the channel
		fmt.Println(s)
		c <- url
	}
}

func main() {
	urls := []string{"https://www.golang.org", "https://www.google.com", "https://www.medium.com"}

	// Declaring a new channel
	c := make(chan string)

	for _, url := range urls {
		// Launching the goroutines
		go checkUrl(url, c)
	}

	fmt.Println("No. of Goroutines:", runtime.NumGoroutine()) // => 4

	// Receiving the messages from the channel
	// for i := 0; i < len(urls); i++ {
	// 	fmt.Println(<-c)
	// }

	// for {
	// 	go checkUrl(<-c, c)
	// 	fmt.Println(strings.Repeat("#", 30))
	// 	time.Sleep(time.Second * 2)
	// }

	// for url := range c {
	// 	time.Sleep(time.Second * 2)
	// 	go checkUrl(url, c)
	// }

	for url := range c {
		go func(u string) {
			time.Sleep(time.Second * 2)
			go checkUrl(u, c)
		}(url)
	}
}
