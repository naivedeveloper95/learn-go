package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	file, err := os.Open("test.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	byteSlice := make([]byte, 2)

	numberOfBytesRead, err := io.ReadFull(file, byteSlice)

	if err != nil && err != io.EOF {
		log.Fatal(err)
	}

	log.Printf("No of bytes read: %d\n", numberOfBytesRead)
	log.Printf("Data read: %s\n", byteSlice)

	file, err = os.Open("main.go")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	data, err := io.ReadAll(file)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Data as string: %s\n", string(data))
	log.Printf("No of bytes read: %d\n", len(data))

	data, err = ioutil.ReadFile("test.txt")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Data read: %s\n", data)
	fmt.Printf("No of bytes read: %d\n", len(data))

}
