package main

import _ "embed"

func main() {
	//go:embed "hello.txt"
	var s string
	print(s)
}
