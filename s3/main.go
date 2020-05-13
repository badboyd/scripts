package main

import (
	"io/ioutil"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	// "net/http"
)

func main() {
	sess, err := session.NewSession(&aws.Config{
		Region:                        aws.String("eu-west-1"),
		CredentialsChainVerboseErrors: aws.Bool(true),
		EnableEndpointDiscovery:       aws.Bool(true)},
	)

	if err != nil {
		panic(err)
	}

	s3svc := s3.New(sess)
	inputparams := &s3.ListObjectsInput{
		Bucket:  aws.String("test-golang-recipes"),
		MaxKeys: aws.Int64(10),
	}
	pageNum := 0

	s3svc.ListObjectsPages(inputparams, func(page *s3.ListObjectsOutput, lastPage bool) bool {
		fmt.Println("Page", pageNum)
		pageNum++
		for _, value := range page.Contents {
			fmt.Println(*value.Key)
		}
		fmt.Println("pageNum", pageNum, "lastPage", lastPage)

		// return if we should continue with the next page
		return true
	})

	obj, err := s3svc.GetObject(&s3.GetObjectInput{
		Key:    aws.String("3"),
		Bucket: aws.String("test-golang-recipes"),
	})


	fmt.Println(err)
	if err == nil {
		b, _ := ioutil.ReadAll(obj.Body)
		fmt.Printf("%s", string(b))
	}

	// http.Get()
}
