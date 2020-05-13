package main

import (
	"encoding/csv"
	"encoding/hex"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

var (
	idQueue            chan string
	secretBin, saltBin []byte

	irisURLFmt        = "/%s/%s.%s"
	irisBaseURL       = "https://cdn.chotot.com"
	irisFormat        = "jpg"
	irisViewOption    = "preset:view"
	irisListingOption = "preset:listing"
	mirageBaseURL     = "https://static.chotot.com.vn/1/images/"

	salt, secret string
	mirageSecret string
	csvFile      string

	irisViewSize      = metrics{T: "view", S: "Iris", MT: "Size"}
	irisViewTime      = metrics{T: "view", S: "Iris", MT: "ms"}
	irisListingSize   = metrics{T: "listing", S: "Iris", MT: "Size"}
	irisListingTime   = metrics{T: "listing", S: "Iris", MT: "ms"}
	mirageViewSize    = metrics{T: "view", S: "Mirage", MT: "Size"}
	mirageViewTime    = metrics{T: "view", S: "Mirage", MT: "ms"}
	mirageListingSize = metrics{T: "listing", S: "Mirage", MT: "Size"}
	mirageListingTime = metrics{T: "listing", S: "Mirage", MT: "ms"}

	prefix = "plain"
)

type metrics struct {
	m   sync.Mutex
	T   string
	S   string
	MT  string
	cnt int64
}

func (mt *metrics) count(delta int64) {
	mt.m.Lock()
	defer mt.m.Unlock()

	mt.cnt += delta
}

func (mt *metrics) print() {
	log.Printf("Type %s Service %s cnt %d MT %s \n", mt.T, mt.S, mt.cnt, mt.MT)
}

func init() {
	flag.StringVar(&salt, "salt", "test", "Salt key")
	flag.StringVar(&secret, "secret", "test", "Secret")
	flag.StringVar(&mirageSecret, "mirage_secret", "test", "Mirage secret")
	flag.StringVar(&csvFile, "csv", "adMedia.csv", "Csv file of ad_media")
	flag.Parse()

	idQueue = make(chan string, 100)

	var err error

	log.Println("Configuration: ", salt, secret, mirageSecret, csvFile)
	if saltBin, err = hex.DecodeString(salt); err != nil {
		log.Fatal("Salt expected to be hex-encoded string")
	}
	if secretBin, err = hex.DecodeString(secret); err != nil {
		log.Fatal("Key expected to be hex-encoded string")
	}
}

func startWorker(in chan string, wg *sync.WaitGroup) {
	// ctx := context.Background()
	defer wg.Done()
	for {
		if id, ok := <-in; !ok {
			break
		} else {
			getIrisImages(id)
			getMirageImages(id)
		}
	}
}

func storeMetrics(size, time int64, imageType, service string) {
	switch {
	case imageType == "listing" && service == "iris":
		irisListingSize.count(int64(size))
		irisListingTime.count(int64(time))
	case imageType == "view" && service == "iris":
		irisViewSize.count(int64(size))
		irisViewTime.count(int64(time))
	case imageType == "listing" && service == "mirage":
		mirageListingSize.count(int64(size))
		mirageListingTime.count(int64(time))
	case imageType == "view" && service == "mirage":
		mirageViewSize.count(int64(size))
		mirageViewTime.count(int64(time))
	}
}

func getIrisImages(id string) {
	irisThumbURL := getIrisThumbURL(fmt.Sprint(id))
	irisImageURL := getIrisImageURL(fmt.Sprint(id))
	getImage(irisThumbURL, "listing", "iris")
	getImage(irisImageURL, "view", "iris")
}

var (
	mirageThumbOption = imageSetting{
		MaxWidth:  200,
		MaxHeight: 150,
		Quality:   80,
		Watermark: false,
	}
	mirageViewOption = imageSetting{
		MaxWidth:  800,
		MaxHeight: 640,
		Quality:   70,
		Watermark: true,
	}
)

func getMirageImages(id string) {
	thumbOption := getMirageImageOption(mirageThumbOption)
	thumbOption.Id = fmt.Sprintf("%010s", id)
	mirageThumbURL := getMirageLink(thumbOption)

	viewOption := getMirageImageOption(mirageViewOption)
	viewOption.Id = fmt.Sprintf("%010s", id)
	mirageViewURL := getMirageLink(viewOption)

	getImage(mirageThumbURL, "listing", "mirage")
	getImage(mirageViewURL, "view", "mirage")
}

func getImage(url, imageType, service string) {
	log.Println("Get image: ", url)
	now := time.Now()

	res, err := http.Get(url)
	if err != nil {
		log.Println("Cannot get http: ", err)
		return
	}
	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println("Cannot get image: ", err)
		return
	}
	storeMetrics(int64(len(data)), time.Since(now).Nanoseconds()/1e6, imageType, service)
}

func main() {
	log.Print("Start worker")

	wg := sync.WaitGroup{}
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go startWorker(idQueue, &wg)
	}

	file, err := os.Open(csvFile)
	if err != nil {
		log.Println("Cannot open file: ", err)
		return
	}
	defer file.Close()

	r := csv.NewReader(file)
	records, err := r.ReadAll()
	if err != nil {
		log.Println("Cannot read all records: ", err)
		return
	}

	for i, v := range records {
		if i%100 == 0 {
			log.Printf("Finish %d images", i)
		}
		idQueue <- v[0]
	}

	close(idQueue)
	wg.Wait()

	log.Print("Stop worker")
	time.Sleep(5 * time.Second)

	irisListingTime.print()
	irisListingSize.print()

	irisViewTime.print()
	irisViewSize.print()

	mirageListingTime.print()
	mirageListingSize.print()

	mirageViewTime.print()
	mirageViewSize.print()
}
