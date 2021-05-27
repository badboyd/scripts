package loop

import (
	"fmt"
	"sync"

	"github.com/corona10/goimagehash"
)

const (
	num = 100000
)

func loop(n int) {
	wg := sync.WaitGroup{}
	for n != 0 {
		l := n
		if n > num {
			n = n - num
			l = num
		} else {
			n = 0
		}

		go func(n, l int) {
			wg.Add(1)
			defer wg.Done()

			fmt.Printf("remaining %d times\n", l)

			for i := 0; i < l; i++ {
				hash()
			}
		}(n, l)
	}
	wg.Wait()
}

func hash() {
	firstPHash := "p:23c27aad96f1e107"
	secondPHash := "p:17c7075b87e60e15"

	firstImgHash, _ := goimagehash.ImageHashFromString(firstPHash)
	secondImgHash, _ := goimagehash.ImageHashFromString(secondPHash)
	firstImgHash.Distance(secondImgHash)
}
