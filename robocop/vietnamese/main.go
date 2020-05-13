package main

import (
	"fmt"
	"regexp"
	"unicode"
	"unicode/utf8"
)

func main() {
	re := regexp.MustCompile(`([\s\t\n]+)`)

	s := `Bán nhà mặt tiền đường trần quý phường 4 quận 11 vị trí đẹp cách đường nguyễn thị nhỏ 50m , kinh doanh thuận lợi.
	- Diện tích: 3.6 x 12,75m
                                                                                                         .
																										 - Diện tích xây dựng : 50m2
                                                                                       - Diện tích sàn : 177m2

                                                                         - Nhà gồm : 1 trệt + 2 lầu + sân thượng.

                                                - Gía : 15 tỷ . pháp lý sổ hồng .`

	s = re.ReplaceAllString(s, " ")
	fmt.Println("Body: ", s)

	var r rune
	var size int
	wordsWithAccents := 0
	words := 0
	n := 0
	for len(s) > 0 {
		r, size = utf8.DecodeRuneInString(s)
		if unicode.IsLetter(r) && size > 1 {
			n++
		} else {
			if unicode.IsSpace(r) {
				words++
				if n > 0 {
					fmt.Println("word: ", s[:size], " ", wordsWithAccents)
					wordsWithAccents++
				}
				n = 0
			}
		}
		s = s[size:]
	}

	if n > 0 {
		wordsWithAccents++
	}

	if !unicode.IsSpace(r) {
		fmt.Printf("IsSpace(%v)\n", r)
		words++
	}

	fmt.Println("wordsWithAccents: ", wordsWithAccents)
	fmt.Println("words: ", words)

	percent := 0
	if words != 0 {
		percent = int(100 * (float32(wordsWithAccents) / float32(words)))
		fmt.Printf("[RR5] Percent of words with accents in body of  : %d\n", percent)
	}

}
