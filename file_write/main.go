package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.OpenFile("my_file.txt", os.O_WRONLY|os.O_CREATE, 0644)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	bufferedWriter := bufio.NewWriter(file)
	bs := []byte{97, 98, 99}
	bytesWritten, err := bufferedWriter.Write(bs)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Wrote %d bytes\n", bytesWritten)

	bytesAvailable := bufferedWriter.Available()
	fmt.Printf("Bytes available: %d\n", bytesAvailable)

	bytesWritten, err = bufferedWriter.WriteString("\nJust a random string")

	if err != nil {
		log.Fatal(err)
	}
	unFlushedBufferSize := bufferedWriter.Buffered()
	log.Printf("Bytes available: %d\n", unFlushedBufferSize)

	bufferedWriter.Flush()

	bufferedWriter.Reset(bufferedWriter)
}
