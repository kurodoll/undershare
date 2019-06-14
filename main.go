package main

import (
	"fmt"
	"strconv"
)

func main() {
	dirListing := getDirectoryListing("F:/ASMR")
	dirListingLen := len(dirListing)

	if dirListingLen > 20 {
		fmt.Println(strconv.Itoa(dirListingLen) + " files found")
	} else {
		for k, v := range dirListing {
			fmt.Println(k + ": " + v)
		}
	}
}
