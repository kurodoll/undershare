package main

import (
	"fmt"
	"os"
)

func getDirectoryListing(directory string) {
	dir, err := os.Open(directory)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer dir.Close()

	listing, err := dir.Readdir(-1)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, file := range listing {
		fmt.Println(file.Name())
	}
}
