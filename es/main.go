package main

import (
	"fmt"
)

func main() {
	esClient := NewESVault("http://esearch.chotot.vn/", "ad_history", "record")
	tmp, err := esClient.Get("7122426210")
	fmt.Println(tmp, "-----------", err)
}
