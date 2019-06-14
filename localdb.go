package main

import (
	"fmt"
	"os"
)

func getDirectoryListing(directory string) map[string]string {
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

	m := make(map[string]string)

	for _, file := range listing {
		fullPath := directory + "/" + file.Name()
		m[file.Name()] = fullPath

		if file.IsDir() {
			m2 := getDirectoryListing(fullPath)
			for k, v := range m2 {
				m[k] = v
			}
		}
	}

	return m
}
