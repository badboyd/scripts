package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

type baz struct {
	foo int
	fb  []foobar
}

type foobar struct {
	bar int
}

func main() {
	bazChan := make(chan *baz)

	go func() {
		for b := range bazChan {
			// fmt.Printf("Inside the loop %p\n", &b)
			go func(b *baz) {
				// fmt.Printf("Inside the routine %p\n", &b)
				time.Sleep(time.Duration(rand.Intn(2)) * time.Second)
				log.Printf("%d", b.foo)
			}(b)
		}
	}()

	bazArr := []baz{
		{
			foo: 1,
			fb: []foobar{
				{
					bar: 1,
				},
			},
		},
		{
			foo: 2,
			fb: []foobar{
				{
					bar: 1,
				},
			},
		},
		{
			foo: 3,
			fb: []foobar{
				{
					bar: 1,
				},
			},
		},
	}

	// fmt.Printf("Original %p\n", &bazArr)
	// for _, b := range bazArr {
	// 	fmt.Printf("Original %p %+v\n", &b, b)
	// 	bazChan <- &b
	// }

	for i := 0; i < len(bazArr); i++ {
		fmt.Printf("idx %p %p\n", &bazArr, &bazArr[i])
		bazChan <- &bazArr[i]
	}

	close(bazChan)
	time.Sleep(5 * time.Second)

	// a := []foobar{
	// 	{
	// 		bar: 1,
	// 	},
	// 	{
	// 		bar: 2,
	// 	},
	// 	{
	// 		bar: 3,
	// 	},
	// }
	// b := make([]foobar, 3)
	// copy(b, a)
	// fmt.Println(a)
	// changeArr(a)
	// fmt.Println(a)
	// fmt.Println(b)

	// tmp := []baz{}
	// for idx, v := range bazArr {
	// 	tmp = append(tmp, bazArr[idx])
	// 	tmp[idx].fb = append([]foobar{}, bazArr[idx].fb...)

	// 	if idx == 1 {
	// 		v.foo = 10000
	// 		v.fb[0].bar = 10
	// 	}
	// }

	// // tmp = append(tmp, bazArr...)

	// for idx := range tmp {
	// 	fmt.Printf("Clone Address %p\n", &tmp[idx].fb[0])
	// 	fmt.Printf("Origin Address %p\n", &bazArr[idx].fb[0])
	// }

	// fmt.Printf("Clone  %+v\n", tmp)
	// fmt.Printf("Origin %+v\n", bazArr)

	// for idx := range bazArr {
	// 	c := bazArr[idx]
	// 	fmt.Printf("c's address %p\n", &c)
	// 	fmt.Printf("bazArr[%d] address %p\n", idx, &c)
	// }

	// b := baz{
	// 	foo: 1,
	// 	fb:  append([]*foobar{}, &foobar{bar: 10}),
	// }

	// bclone := clone(&b)
	// // bclone.fb.bar = 11
	// fmt.Printf("bclone address %p : %+v\n", &bclone, bclone)

	// fmt.Printf("b address %p : %+v\n ", &b, b)
}

func clone(b *baz) baz {
	if b == nil {
		return baz{}
	}

	bb := *b

	tmp := make([]foobar, len(b.fb))
	fmt.Println(copy(tmp, b.fb))

	bb.fb = tmp

	// tmp.fb = clonefoobar(tmp.fb)

	return bb
}

func clonefoobar(fb *foobar) foobar {
	if fb != nil {
		return *fb
	}

	return foobar{}
}

func changeArr(a []foobar) {
	if len(a) > 1 {
		a[1].bar = -1
	}
}
