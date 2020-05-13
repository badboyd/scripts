package main

import (
	"bytes"
	"fmt"

	"github.com/corona10/goimagehash"
)

func main() {
	firstPHash := "p:23c27aad96f1e107"
	secondPHash := "p:17c7075b87e60e15"

	firstImgHash, _ := goimagehash.ImageHashFromString(firstPHash)
	secondImgHash, _ := goimagehash.ImageHashFromString(secondPHash)
	distance, _ := firstImgHash.Distance(secondImgHash)
	fmt.Printf("Distance %d", distance)

	for i := 0; i < 33; i++ {
		fmt.Printf("\"word_%d\"", i)
		if i != 32 {
			fmt.Print(", ")
		}
	}
	test(nil)
}

func test(body *bytes.Buffer) {
	if body == nil {
		fmt.Println("body is nil")
	}
}
