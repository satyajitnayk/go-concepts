package main

import (
	"errors"
	"fmt"
	"io"
	"os"
)

func runPanicRecoverySimpleFunction() {
	var file io.ReadCloser
	file, err := openCSV("data.csv")
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}
	defer file.Close()

	// Do something
}

// 1. defer is a powerful and useful way of dealing with panics,
// as well as reliably cleaning up
// 2. Best practices suggest closing files, network connections, and other similar
// resources inside a defer clause. This ensures that even when errors or panics
// occur, system resources will be freed.
func openCSV(filename string) (file *os.File, err error) {
	defer func() {
		if r := recover(); r != nil {
			file.Close()
			err = r.(error)
		}
	}()

	file, err = os.Open(filename)
	if err != nil {
		fmt.Printf("Failed to open file\n")
		return file, err
	}

	removeEmptyLines(file)
	return file, err
}

func removeEmptyLines(f *os.File) {
	panic(errors.New("Failed parse"))
}
