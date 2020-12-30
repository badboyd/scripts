package main

import "fmt"

// func main() {
// 	// first 3 lines
// 	lines := make([]string, 3, 3)
// 	for idx := range lines {
// 		in := bufio.NewReader(os.Stdin)
// 		line, err := in.ReadString('\n')
// 		if err != nil {
// 			panic(err)
// 		}
// 		fmt.Println(line)
// tmp := ""
// if _, err := fmt.Scan(&tmp); err != nil {
// 	panic(err)
// }
// lines[idx] = line
// fmt.Println(tmp)
// }

// K max items at one time
// N total orders in warehouse
// K, N, err := readFirstLine(lines[0])
// if err != nil {
// 	panic(err)
// }

// weights, orders, err := readOrders(N, lines[1], lines[2])
// if err != nil {
// 	panic(err)
// }

// }

// p r3 o1 g a mm e

func main() {
	a := map[string]int{
		"a": 1,
		"r": 3,
		"p": 1,
		"o": 1,
		"m": 2,
		"g": 1,
		"e": 1,
	}

	b := map[string]int{
		"a": 1,
		"r": 3,
		"p": 1,
		"o": 1,
		"m": 2,
		"g": 1,
		"e": 1,
	}

	hyphen := "_"

	tmp := "prog_rma_emrppprmmograe_r"
	j := len(tmp) - 1
	i := 0

	for j > i {
		if string(tmp[i]) == hyphen {
			i++
		} else {
			if m, ok := a[string(tmp[i])]; ok {
				fmt.Println(string(tmp[i]))

				m--
				a[string(tmp[i])] = m

				if m == 0 {
					delete(a, string(tmp[i]))
				}

				i++
			}

		}

		if string(tmp[j]) == hyphen {
			j--
		} else {
			if m, ok := b[string(tmp[j])]; ok {

				m--
				b[string(tmp[j])] = m

				if m == 0 {
					delete(b, string(tmp[j]))
				}

				j--
			}
		}

		if len(a) == 0 && len(b) == 0 {
			break
		}
	}

	fmt.Println(j - i + 1)
}

// func main() {
//     reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

//     nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
//     checkError(err)
//     n := int32(nTemp)

//     fizzBuzz(n)
// }

// func readLine(reader *bufio.Reader) string {
//     str, _, err := reader.ReadLine()
//     if err == io.EOF {
//         return ""
//     }

//     return strings.TrimRight(string(str), "\r\n")
// }

// func checkError(err error) {
//     if err != nil {
//         panic(err)
//     }
// }

// // read first line returns number of orders to pick and maximum items can pick at one time
// func readFirstLine(s string) (int, int, error) {
// 	// {K, N}
// 	// K: max items at one time
// 	// N: total number of orders
// 	var K, N int

// 	_, err := fmt.Scanf("%d %d", &K, &N)
// 	return K, N, err

// }

// // returns ordered max value orders
// // map for order value and order items
// func readOrders(n int, w string, v string) ([]int, map[int]int, error) {
// 	weights, err := convertStrArrToIntArr(strings.SplitN(w, " ", n))
// 	if err != nil {
// 		return nil, nil, err
// 	}

// 	values, err := convertStrArrToIntArr(strings.SplitN(v, " ", n))
// 	if err != nil {
// 		return nil, nil, err
// 	}

// 	m := make(map[int]int)
// 	for idx, v := range values {
// 		m[v] = weights[idx]
// 	}

// 	sort.Slice(values, func(i, j int) bool {
// 		return values[i] > values[j]
// 	})

// 	return values, m, nil
// }

// func convertStrArrToIntArr(a []string) ([]int, error) {
// 	res := make([]int, len(a))
// 	for idx, v := range a {
// 		n, err := strconv.Atoi(v)
// 		if err != nil {
// 			return nil, err
// 		}

// 		res[idx] = n
// 	}

// 	return res, nil
// }
