package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	// fmt.Printf("%T\n", scanner)

	scanner.Scan()

	text := scanner.Text()
	bytes := scanner.Bytes()

	fmt.Println("Input text: ", text)
	fmt.Println("Input bytes: ", bytes)

	for scanner.Scan() {
		text = scanner.Text()
		fmt.Println("You entered: ", text)

		if text == "exit" {
			fmt.Println("Exiting the scanning process..")
			break
		}
	}

	fmt.Println("Just a mmessage after the loop")

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
