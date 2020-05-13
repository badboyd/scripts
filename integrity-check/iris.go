package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"strings"
)

func getImageFullName(id string) string {
	return fmt.Sprintf("%s/%s", prefix,
		strings.TrimRight(strings.TrimLeft(id, "0"), ".jpg"))
}

func genURL(id, options, format string) string {
	source := getImageFullName(id)
	path := fmt.Sprintf(irisURLFmt, options, source, format)

	mac := hmac.New(sha256.New, secretBin)
	mac.Write(saltBin)
	mac.Write([]byte(path))
	signature := base64.RawURLEncoding.EncodeToString(mac.Sum(nil))

	return fmt.Sprintf("%s/%s%s", irisBaseURL, signature, path)
}

// GetIrisThumbURL returns url for ad listing
func getIrisThumbURL(id string) string {
	return genURL(id, irisListingOption, irisFormat)
}

// GetIrisImageURL returns url for image in adview
func getIrisImageURL(id string) string {
	return genURL(id, irisViewOption, irisFormat)
}
