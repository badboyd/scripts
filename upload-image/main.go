package main

import (
	"fmt"
	"net/http"
	"os"
	// "bytes"
	// "encoding/json"
	"io/ioutil"
)

// type Token struct {
// 	Token string `json:"token"`
// }

func main() {
	args := os.Args[1:]
	if len(args) < 2 {
		fmt.Println("not enough params")
		return
	}
	imageUrl := args[0]
	baseUrl := args[1]
	fmt.Println(uploadImage(imageUrl, baseUrl))
}

func uploadImage(imageUrl string, baseUrl string) (url string, err error) {
	client := NewClient(nil, BaseURL(baseUrl))
	fmt.Println(imageUrl)
	var output ImgPutResponse
	resp, err := http.Get(imageUrl)
	fmt.Println(resp)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	client.ImgPut(&ImgPutRequest{
		Image:     body,
		Thumbnail: body,
	}, &output)
	fmt.Println(output)
	return output.ImageName, err
}
