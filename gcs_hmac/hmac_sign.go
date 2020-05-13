package main

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"net/url"
	"text/template"
	"time"
)

var (
	urlFmt = "{{.BaseURL}}/" +
		"{{.BucketName}}/{{.ObjectName}}" +
		"?GoogleAccessId={{.AccessKey}}" +
		"&Expires={{.Expiration}}" +
		"&Signature={{.SignedSignature}}"

	signatureFmt = "GET\n" +
		"{{with .ContentMD5}}{{.}}{{end}}\n" +
		"{{with .ContentType}}{{.}}{{end}}\n" +
		"{{with .Expiration}}{{.}}{{end}}\n" +
		"{{with .Resource}}{{.}}{{end}}"
)

// URL ...
type URL struct {
	BaseURL        string
	SecretKey      string
	AccessKey      string
	ExpirationTime time.Duration
	template       *template.Template
	signature      *template.Template
}

// Sign creates a new url
func (u *URL) Sign(name, bucket string) (string, error) {
	data := map[string]string{
		"AccessKey":  u.AccessKey,
		"BucketName": bucket,
		"BaseURL":    u.BaseURL,
		"ObjectName": name,
		"Resource":   fmt.Sprintf("/%s/%s", bucket, name),
		"Expiration": fmt.Sprint(time.Now().Add(u.ExpirationTime).Unix()),
	}

	signature := bytes.Buffer{}
	if err := u.signature.Execute(&signature, data); err != nil {
		return "", err
	}

	mac := hmac.New(sha1.New, []byte(u.SecretKey))
	mac.Write(signature.Bytes())

	sig := base64.StdEncoding.EncodeToString(mac.Sum(nil))
	data["SignedSignature"] = url.QueryEscape(sig)

	url := bytes.Buffer{}
	if err := u.template.Execute(&url, data); err != nil {
		return "", err
	}

	return url.String(), nil
}

// Init the URL struct
func (u *URL) Init() error {
	tmp, err := template.New("").Parse(urlFmt)
	if err != nil {
		return err
	}
	u.template = tmp

	tmp, err = template.New("").Parse(signatureFmt)
	if err != nil {
		return err
	}
	u.signature = tmp

	return nil
}

func main() {
	ctURL := URL{
		BaseURL:        "https://cdn.badboyd.com",
		SecretKey:      "0CVQiJq4Om/47hqADKVZK08E0ggmXfHd2lq8x51z",
		AccessKey:      "GOOG1EZZDTPVEPPYI4CEXAXHMJ3YBP3Q23LQ2N3CI2HC5AJFYJYH6RLAG2BMQ",
		ExpirationTime: 3 * time.Minute,
	}
	if err := ctURL.Init(); err != nil {
		panic(err)
	}

	fmt.Println(ctURL.Sign("d22c8bc8caba71b28fce4200703cdec2-2633799274031022595.jpg", "chotot-photo-staging"))
}
