package main

import (
	"fmt"
	"regexp"
)

func main() {
	// regexp.CompilePOSIX(expr)
	re := regexp.MustCompilePOSIX(`httperrorstatus: (.*); reason: (.*); code: (.*); details: (.*)`)
	var str = `http error status: 404; reason: app instance has been unregistered; code: registration-token-not-registered; details: Requested entity was not found.`

	for i, match := range re.FindAllSubmatch([]byte(str), -1) {
		fmt.Println(match, "found at index", i)
	}
}
