package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println(os.Args)
	fmt.Println(os.Args[1])

	filePath := os.Args[1]

	fileReference, err := os.Open(filePath)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, fileReference)
}
