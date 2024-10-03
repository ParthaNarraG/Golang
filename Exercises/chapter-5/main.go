package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Functions Exercises")
	var length, _ = fileLen("D:/Golang/Control Statements/main.go")
	fmt.Println("length", length)
}

// Write a function called fileLen that has an input parameter of type string and returns an int and an error.
// The function takes in a filename and returns the number of bytes in the file.
// If there is an error reading the file, return the error. Use defer to make sure the file is closed properly.

func fileLen(fileName string) (int, error) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("error", err)
		return 0, err
	}
	defer file.Close()
	fileInfo, err := file.Stat()
	if err != nil {
		return 0, err
	}
	return int(fileInfo.Size()), nil
}
