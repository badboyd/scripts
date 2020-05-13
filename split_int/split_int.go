package main

import (
	"fmt"
)

func splitInt(i int) []int {
	fmt.Print(i)
	res := []int{}
	for tmp := i % 10; i > 0; tmp = i % 10 {
		res = append(res, tmp)
		i /= 10
	}
	return res
}

var numMap = map[int]int{
	'0': 0,
	'1': 1,
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
}

func atoi(input string) (int, error) {
	result := 0
	inputLen := len(input)

	for idx, ch := range input {
		num, ok := numMap[int(ch)]
		if !ok {
			return 0, fmt.Errorf("%s is not a valid number", input)
		}

		result += num * power(10, inputLen-idx)
	}
	return result, nil
}

func power(base, num int) int {
	result := 1
	for i := 1; i < num; i++ {
		result *= base
	}
	return result
}

func main() {
	for i := 1; i <= 100; i++ {
		switch {
		case i%15 == 0:
			fmt.Println("FizzBuzz")
		case i%5 == 0:
			fmt.Println("Buzz")
		case i%3 == 0:
			fmt.Println("Fizz")
		default:
			fmt.Println(i)
		}
	}
	// fmt.Println(splitInt(0123))
	fmt.Println(atoi("123"))
}
