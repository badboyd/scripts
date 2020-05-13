package main

import (
	"bytes"
	"database/sql"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	_ "github.com/lib/pq"

	"git.chotot.org/ad-service/spine/pkg/types"
)

var (
	dbClient *sql.DB

	insertImage = `
		INSERT INTO ad_media (
			ad_media_id,
			upload_time,
			media_type,
			seq_no,
 			hide,
			digest
		)
		VALUES (
			$1,
			CURRENT_TIMESTAMP,
			$2,
			$3,
			'f',
			$4
		)
	`
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("not enough params")
		return
	}
	brainstemURL := args[1]
	brainstemURLProd := args[2]
	from, _ := strconv.Atoi(args[3])
	limit, _ := strconv.Atoi(args[4])

	connStr := "user=postgres dbname=blocketdb host=10.60.7.10 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	dbClient = db
	defer dbClient.Close()

	fileName := args[0]
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	r := csv.NewReader(file)
	records, err := r.ReadAll()
	if err != nil {
		fmt.Println(err)
		return
	}

	var cnt types.Counter
	var wg sync.WaitGroup

	queue := make(chan string, 10)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go startWorker(i, queue, brainstemURL, brainstemURLProd, &cnt, &wg)
	}

	if from == 0 {
		from = 1
	}

	to := limit + from
	if limit >= len(records) || to >= len(records) || limit == 0 {
		to = len(records)
	}
	log.Printf("Start processing msg from %d to %d", from, to)
	for i := from; i < to; i++ {
		if i%100 == 0 {
			log.Printf("Finish %d ads", i)
		}
		queue <- records[i][0]
	}
	close(queue)
	wg.Wait()

	log.Printf("Finish processing %d msg", cnt.Value())
}

func startWorker(id int, queue chan string, brainstemURL, brainstemURLProd string, cnt *types.Counter, wg *sync.WaitGroup) {
	log.Printf("Start worker [%d]", id)
	defer wg.Done()
	for adID := range queue {
		in := loadAd(adID, brainstemURLProd)
		uploadImage(in)
		// fmt.Println(in)
		if in != nil {
			if err := newAd(in, brainstemURL); err != nil {
				log.Printf("Worker [%d] error with ad_id [%v]: %v", id, adID, err)
			}
		}
		cnt.Inc()
	}
	log.Printf("finish worker [%d]", id)
}

func uploadImage(in map[string]interface{}) {
	for i := 0; i < 12; i++ {
		if img := in[fmt.Sprintf("image_id%d", i)]; img != nil {
			imgID := strings.TrimLeft(strings.TrimRight(img.(string), ".jpg"), "0")
			_, err := dbClient.Exec(insertImage, imgID, "image", i, in[fmt.Sprintf("digest_%d", i)].(string))
			log.Println(err)
		}
	}
}

func loadAd(adID, brainstemBaseURL string) map[string]interface{} {
	v := url.Values{}
	v.Set("ad_id", adID)

	path := brainstemBaseURL + "/api/v1/private/flashad/loadad"
	req, err := newHTTPRequest("GET", path, nil)
	if err != nil {
		log.Println(err)
		return nil
	}
	req.URL.RawQuery = v.Encode()

	tmp := make(map[string]interface{})
	_, err = httpClientDo(req, &tmp)
	if err != nil {
		log.Println(err)
		return nil
	}

	for k, v := range tmp["ad"].(map[string]interface{}) {
		tmp[k] = v
	}
	delete(tmp, "ad")

	for k, v := range tmp["params"].(map[string]interface{}) {
		tmp[k] = v
	}
	delete(tmp, "params")

	if tmp["images"] != nil {
		images := tmp["images"].(map[string]interface{})
		for k, v := range images {
			tmp["image_id"+k] = v.(map[string]interface{})["name"]
			tmp["digest_"+k] = v.(map[string]interface{})["digest"]
		}
		delete(tmp, "images")
	}

	for k, v := range tmp["users"].(map[string]interface{}) {
		tmp[k] = v
	}
	delete(tmp, "users")

	tmp["landed_type"] = 1
	tmp["apartment_type"] = 1
	tmp["mobile_model"] = 1
	tmp["commercial_type"] = 1
	tmp["verified_account"] = 1
	tmp["source"] = "android"
	return tmp
}

func newAd(in map[string]interface{}, brainstemBaseURL string) error {
	if in == nil {
		return fmt.Errorf("NIL_INPUT")
	}
	delete(in, "ad_id")
	delete(in, "list_id")

	path := brainstemBaseURL + "/api/v1/private/flashad/new"
	req, err := newHTTPRequest("POST", path, in)
	if err != nil {
		return err
	}

	_, err = httpClientDo(req, nil)
	return err
}

func editAd(in map[string]interface{}, brainstemBaseURL string) error {
	if in == nil {
		return fmt.Errorf("NIL_INPUT")
	}

	path := brainstemBaseURL + "/api/v1/private/flashad/edit"
	req, err := newHTTPRequest("POST", path, in)
	if err != nil {
		return err
	}

	_, err = httpClientDo(req, nil)
	return err
}

var httpClient = &http.Client{
	Timeout: 5 * time.Second,
}

// Should we add the token ?
func newHTTPRequest(method, path string, body interface{}) (*http.Request, error) {
	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		if err := json.NewEncoder(buf).Encode(body); err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, path, buf)
	if err != nil {
		return nil, err
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Owner", "badboyd")

	return req, nil
}

func httpClientDo(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer func() {
		io.Copy(ioutil.Discard, resp.Body)
		resp.Body.Close()
	}()

	if resp.StatusCode/100 != 2 {
		body, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf("%v", string(body))
	}

	if v != nil {
		d := json.NewDecoder(resp.Body)
		d.UseNumber()
		d.Decode(v)
	}

	return resp, err
}
