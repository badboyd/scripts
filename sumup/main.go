package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {

	fmt.Println(validBraces2("([{}])"))
}

func OrderByWeight(strn string) string {
	nums := strings.Split(strn, " ")

	sort.Slice(nums, func(i, j int) bool {
		wi := calcWeight(nums[i])
		fmt.Printf("w %s n %d\n", nums[i], wi)
		wj := calcWeight(nums[j])
		fmt.Printf("w %s n %d\n", nums[j], wj)

		if wi == wj {
			return nums[i] < nums[j]
		}

		return wi < wj
	})

	return strings.Join(nums, " ")
}

func calcWeight(s string) int {
	n := 0
	for _, ch := range []byte(s) {
		ch -= '0'
		if ch > 9 {
			return 0
		}
		n = n*10 + int(ch)
	}

	return n
}

func validBraces2(name string) bool {
	stack := []byte{}

	for _, ch := range []byte(name) {
		switch ch {
		case '[', '{', '(':
			stack = append(stack, ch)
		case ']':
			if stack[len(stack)-1] != '[' {
				return false
			}
			stack = stack[:len(stack)-1]
		case '}':
			if stack[len(stack)-1] != '{' {
				return false
			}
			stack = stack[:len(stack)-1]
		case ')':
			if stack[len(stack)-1] != '(' {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return true
}

func validBraces(name string) bool {
	closedChs := map[byte]int{
		'}': 0,
		']': 0,
		')': 0,
	}

	braces := map[byte]byte{
		'{': '}',
		'[': ']',
		'(': ')',
	}

	for _, ch := range []byte(name) {
		switch ch {
		case '[', '{', '(':
			closedChs[braces[ch]] = closedChs[braces[ch]] - 1
		case ']', '}', ')':
			closedChs[ch] = closedChs[ch] + 1
		}
	}

	fmt.Println(closedChs)

	for _, c := range closedChs {
		if c != 0 {
			return false
		}
	}

	return true
}
