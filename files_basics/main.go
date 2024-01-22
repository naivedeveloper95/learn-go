package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	var newFile *os.File
	fmt.Printf("%T\n", newFile)

	var err error

	newFile, err = os.Create("a.txt")
	if err != nil {
		// return nil, err
		// os.Exit(1)
		log.Fatal(err)
	}

	err = os.Truncate("a.txt", 0)
	if err != nil {
		// return nil, err
		// os.Exit(1)
		log.Fatal(err)
	}

	newFile.Close()

	// file, err := os.Open("a.txt")
	file, err := os.OpenFile("a.txt", os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}

	file.Close()

	var fileInfo os.FileInfo
	fileInfo, err = os.Stat("a.txt")
	p := fmt.Println

	p("File Name: ", fileInfo.Name())
	p("File Size: ", fileInfo.Size())
	p("Last Modified: ", fileInfo.ModTime())
	p("IsDir? : ", fileInfo.IsDir())
	p("Permissions? : ", fileInfo.Mode())

	fileInfo, err = os.Stat("b.txt")

	if err != nil {
		if os.IsNotExist(err) {
			p("The file does not exist")
		}
	}

	oldPath := "a.txt"
	newPath := "aaa.txt"

	err = os.Rename(oldPath, newPath)
	if err != nil {
		log.Fatal(err)
	}

	err = os.Remove("aaa.txt")

	if err != nil {
		log.Fatal(err)
	}
}
