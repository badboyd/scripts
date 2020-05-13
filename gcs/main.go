package main

import (
	"context"

	"cloud.google.com/go/storage"

	"git.chotot.org/go-kafka-consumer/logger"
)

var (
	log = logger.GetLogger("gcs-rename")

	file         string
	csvFile      string
	gcsClient    *storage.Client
	bucketHandle *storage.BucketHandle
)

func init() {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatalf("Cannot init GCS client: %v", err)
	}
	gcsClient = client
	bucketHandle = gcsClient.Bucket("st-chotot-org")
	// bucketHandle = gcsClient.Bucket("chotot-photo-staging")
}

func main() {
	ctx := context.Background()

	id := "1037859198.jpg"
	attrs, err := bucketHandle.Object(id).Attrs(ctx)
	if err != nil {
		return
	}

	log.Infof("MD5 %02x", attrs.MD5)
}
