package main

import (
	"fmt"

	"errors"
)

func main() {
	errA := errors.New("A")
	errB := errors.New("B")

	errC := fmt.Errorf("%s %w", errA.Error(), errB)

	if errors.Is(errC, errB) {
		fmt.Println("errB")
	}
}
