// Package tips contains tips for writing Cloud Functions in Go.
package tips

import (
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// ListFiles lists the files in the current directory.
func ListFiles(w http.ResponseWriter, r *http.Request) {
	files, err := ioutil.ReadDir("./")
	if err != nil {
		http.Error(w, "Unable to read files", http.StatusInternalServerError)
		log.Printf("ioutil.ListFiles: %v", err)
		return
	}
	fmt.Fprintln(w, "Files:")
	for _, f := range files {
		fmt.Fprintf(w, "\t%v\n", f.Name())
	}
	loadWm()
}

func loadWm() {
	existingImageFile, err := os.Open("watermark.png")
	if err != nil {
		fmt.Printf("%v", err)
	}
	defer existingImageFile.Close()

	if _, _, err := image.Decode(existingImageFile); err != nil {
		fmt.Printf("%v", err)
	}
}
