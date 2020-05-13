package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

var (
	inputLinkPattern    = "https://drive.google.com/open?id=%s"
	downloadLinkPattern = "https://drive.google.com/uc?export=download&id=%s"
)

func main() {
	in := "https://drive.google.com/open?id=1tkjDOODIh-qMKIro7ACUyCe5U1efi0Jm"
	expectedOut := "https://drive.google.com/uc?export=download&id=1tkjDOODIh-qMKIro7ACUyCe5U1efi0Jm"

	dwLink := getDownloadLink(in)
	if dwLink != expectedOut {
		fmt.Printf("Expect %s but got %s\n", expectedOut, dwLink)
	}

	saveImage(dwLink)
}

func getDownloadLink(link string) string {
	id := ""
	if _, err := fmt.Sscanf(link, inputLinkPattern, &id); err != nil {
		fmt.Println(err)
	}
	return fmt.Sprintf(downloadLinkPattern, id)
}

func saveImage(link string) {
	res, err := http.Get(link)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	file, err := os.Create("test.jpg")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	if _, err := io.Copy(file, res.Body); err != nil {
		fmt.Println(err)
		return
	}
}
