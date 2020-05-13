package main

import (
	"fmt"

	cconf "git.chotot.org/ad-service/cconf/client"
)

func main() {
	cClient := cconf.NewClient(
		// change the ip to your cconf ip
		cconf.BaseURL("http://10.60.3.47:9090"),
	)

	// Get config by key, you can set default value
	// you can get value in string, in or bool
	numberOfExtraImage := cClient.GetString("category_settings.extra_images_newad.default", "")
	fmt.Println(numberOfExtraImage)

	// Get config by prefix
	// response is in map[string]interface{}
	extraImage := cClient.GetMap("category_settings.extra_images_newad.*")
	fmt.Println(extraImage)

	// response is in map[string]interface{}
	// need to code :D
	allConfig := cClient.GetAll()
	fmt.Println(allConfig)
}
