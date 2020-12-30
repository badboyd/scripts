package main

import (
	"fmt"
	"strings"
)

func solution(T []int) string {
	monthIdx := 0
	largestAmpl := 0

	for i := 0; i < 4; i++ {
		startIdx := i * len(T) / 4

		tempAmpl := calcAmplitude(T[startIdx : startIdx+len(T)/4])
		fmt.Println(tempAmpl)

		if tempAmpl > largestAmpl {
			largestAmpl = tempAmpl
			monthIdx = i
		}
	}

	switch monthIdx {
	case 0:
		return "WINTER"
	case 1:
		return "SPRING"
	case 2:
		return "SUMMER"
	case 3:
		return "AUTUMN"
	}

	return ""
}

func calcAmplitude(T []int) int {
	fmt.Println(T)

	if len(T) < 1 {
		return 0
	}

	minNum := T[0]
	maxNum := T[0]

	for i := 1; i < len(T); i++ {
		if T[i] < minNum {
			minNum = T[i]
		}

		if T[i] > maxNum {
			maxNum = T[i]
		}
	}

	return maxNum - minNum
}

func formatPhone(s string) string {
	fmt.Println(len(s))
	return formatter(strings.NewReplacer("-", "", " ", "").Replace(s))
}

func formatter(s string) string {
	if len(s) <= 3 {
		return s
	}

	if len(s) == 4 {
		return s[0:2] + "-" + s[2:]
	}

	return s[0:3] + "-" + formatter(s[3:])
}

func main() {

	fmt.Println(formatPhone("00-44 1321 13123213213  14 12 3 123 12 3 12 3 123  1 321 3123 1 23 1231 3 1 12 312 3 12 3123--123123"))

	fmt.Println(solution([]int{2, 3, 3, 1, 10, 8, 2, 5, 13, -5, 3, -18}))

	// fmt.Println(calcAmplitude([]int{4, 2, -1, 5}))
}
