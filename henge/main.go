package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	// read first line
	firstLine, err := reader.ReadString('\n')
	panicOnError(err)

	// convert to number of testcases
	numTestcase, err := strconv.Atoi(firstLine[:len(firstLine)-1])
	panicOnError(err)

	// read all testcases
	testcases := readTestcases(numTestcase, reader)

	// process all testcases
	processTestCases(testcases)
}

// read all testcases
func readTestcases(numTcs int, reader *bufio.Reader) [][]string {
	if numTcs == 0 {
		return [][]string{}
	}

	// read len of the this testcase array
	line, err := reader.ReadString('\n')
	panicOnError(err)

	// convert the len to int
	arrLen, err := strconv.Atoi(line[:len(line)-1])
	panicOnError(err)

	// read the array
	line, err = reader.ReadString('\n')
	panicOnError(err)

	// split by space
	arr := strings.Split(line[:len(line)-1], " ")
	if len(arr) != arrLen {
		panic("Input length is invalid")
	}

	numTcs--
	return append([][]string{arr}, readTestcases(numTcs, reader)...)
}

// process all testcases
func processTestCases(tcs [][]string) {
	if len(tcs) == 0 {
		return
	}

	// print the result here
	fmt.Println(calcSumSquares(tcs[0]))

	// higher bound always lower than cap
	// this condition is for safe
	if len(tcs) > 0 {
		processTestCases(tcs[1:])
	}
}

// calc sum of squares from integer array
func calcSumSquares(arr []string) int {
	if len(arr) == 0 {
		return 0
	}

	num, err := strconv.Atoi(arr[0])
	panicOnError(err)

	if num < 0 {
		num = 0
	}

	if len(arr) > 0 {
		return num*num + calcSumSquares(arr[1:])
	}

	return num * num
}

// panic if err is not nil
// it is a utility function
func panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}
