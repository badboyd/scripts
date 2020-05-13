package main

import (
	"bufio"
	"context"
	// "encoding/csv"
	"flag"
	"fmt"
	"hash/crc32"
	// "io"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"sync"
	"time"

	"cloud.google.com/go/storage"

	"git.chotot.org/go-kafka-consumer/logger"
)

var (
	myIP = getMyIP()
	log  = logger.GetLogger("ceph2gcs")

	file         string
	csvFile      string
	gcsClient    *storage.Client
	bucketHandle *storage.BucketHandle
)

func getMyIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}

	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}

func init() {
	flag.StringVar(&file, "file", "ad_media_id.txt", "File of ad_media")
	flag.StringVar(&csvFile, "csv-file", "ad_media_id.csv", "File of ad_media")
	flag.Parse()

	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatalf("Cannot init GCS client: %v", err)
	}
	gcsClient = client
	bucketHandle = gcsClient.Bucket("chotot-photo-production")
	// bucketHandle = gcsClient.Bucket("chotot-photo-staging")
}

func formatImageID(i interface{}) string {
	tmp := fmt.Sprint(i)
	if len(tmp) < 10 {
		return fmt.Sprintf("%010v", i)
	}
	return tmp
}

var (
	cTable = crc32.MakeTable(crc32.Castagnoli)
)

func tryPutObject(ctx context.Context, imageID interface{}) error {
	var err error
	for i := 0; i < 4; i++ {
		if err = putObject(ctx, imageID); err == nil || err != context.DeadlineExceeded {
			break
		}
		log.Infof("time [%d] put object %v err: %v", i, imageID, err)
	}
	return err
}

func putObject(ctx context.Context, imageID interface{}) error {
	id := fmt.Sprint(imageID) + ".jpg"
	oh := bucketHandle.Object(id)
	if _, err := oh.Attrs(ctx); err == nil {
		log.Infof("Image %s already exists", id)
		return nil
	}

	imageURL := "https://static.chotot.com.vn/raw/images/" + formatImageID(imageID) + ".jpg"
	resp, err := http.Get(imageURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return fmt.Errorf("Cannot get image from %s", imageURL)
	}
	//
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	timeoutCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	ow := oh.NewWriter(timeoutCtx)
	//
	ow.CRC32C = crc32.Checksum(data, cTable)
	ow.SendCRC32C = true

	log.Infof("Put image %v", id)

	writeTime := time.Now()
	if _, err = ow.Write(data); err != nil {
		return err
	}
	// io.Copy(ow, resp.Body)
	log.Infof("GCS [%s] Write time: %f", id, time.Since(writeTime).Seconds())

	if err != nil {
		return fmt.Errorf("Cannot copy data to object writer: %v", err)
	}

	// time.Sleep(3 * time.Second)
	closeTime := time.Now()
	err = ow.Close()
	log.Infof("GCS [%s] Close time: %f", id, time.Since(closeTime).Seconds())

	return err
}

var (
	imagePattern = "gs://chotot-photo-production/%d.jpg:"
)

func main() {
	if gcsClient != nil {
		defer gcsClient.Close()
	}

	ctx := context.Background()
	processing := func(ic chan string, wg *sync.WaitGroup) {
		wg.Add(1)
		for {
			if id, ok := <-ic; ok {
				log.Infof("[%s] err: %v", id, tryPutObject(ctx, id))
			} else {
				break
			}
		}
		defer wg.Done()
	}

	idChan := make(chan string)
	var wg sync.WaitGroup
	for i := 0; i < 30; i++ {
		go processing(idChan, &wg)
	}

	file, err := os.Open(file)
	if err != nil {
		log.Info("Cannot open file: ", err)
		return
	}
	defer file.Close()

	// csvReader := csv.NewReader(file)
	// numLines := 0
	// for {
	// 	record, err := csvReader.Read()
	// 	if err == io.EOF {
	// 		break
	// 	}
	//
	// 	if err != nil {
	// 		log.Info("Cannot read csv: ", err)
	// 	}
	// 	numLines++
	// 	if numLines%10000 == 0 {
	// 		log.Infof("Finish %d images", numLines)
	// 	}
	//
	// 	if record[0] != "" {
	// 		idChan <- record[0]
	// 	}
	// }

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		id := 0
		fmt.Sscanf(scanner.Text(), imagePattern, &id)
		log.Info("Image id: ", id)
		if id != 0 {
			idChan <- fmt.Sprint(id)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Info("Scanner error: ", err)
	}

	close(idChan)
	wg.Wait()

	log.Info("Finish worker")
}
