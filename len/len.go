package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Len of null: ", len(([]string)(nil)))
	fmt.Println("Len of null: ", strings.Join(([]string)(nil)[:0], ","))
	// fmt.Println("Len of null: ", []string(nil)[len(([]string)(nil))])
	arr := make([]string, 0, 0)
	fmt.Println("Len of null: ", len(arr))
	fmt.Println("Len of null: ", arr[:len(arr)])
}
