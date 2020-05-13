// package main
//
// import (
// 	"encoding/csv"
// 	"fmt"
// 	"strconv"
// 	"strings"
// 	// "log"
// 	spine "git.chotot.org/ad-service/spine/client"
// 	"git.chotot.org/ad-service/test/trans"
//
// 	"os"
// )
//
// var (
// 	adType = map[string]string{
// 		"sell": "s",
// 		"buy":  "k",
// 		"let":  "u",
// 		"rent": "h",
// 		"swap": "b",
// 	}
// )
//
// var (
// 	defaultIgnoreFields = []string{}
//
// 	spineCli *spine.Client
// )
//
// func main() {
// 	args := os.Args[1:]
// 	if len(args) < 4 {
// 		fmt.Println("not enough params")
// 		return
// 	}
//
// 	transHost := args[0]
// 	transPort, _ := strconv.Atoi(args[1])
//
// 	spineUrl := args[2]
// 	spineCli = spine.NewClient(nil, spine.BaseURL(spineUrl))
//
// 	fileName := args[3]
// 	username := args[4]
// 	passwd := args[5]
//
// 	ignoredFields := map[string]int{}
//
// 	if len(args) > 6 {
// 		for _, v := range strings.Split(args[6], ",") {
// 			ignoredFields[v] = 1
// 		}
// 	} else {
// 		for _, v := range defaultIgnoreFields {
// 			ignoredFields[v] = 1
// 		}
// 	}
//
// 	trans.Init(transHost, transPort)
//
// 	file, err := os.Open(fileName)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	defer file.Close()
//
// 	r := csv.NewReader(file)
// 	records, err := r.ReadAll()
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
//
// 	token := Authenticate(username, passwd)
// 	params := records[0]
// 	for i := 1; i < len(records); i++ {
// 		token = AcceptAd(token, params, records[i], ignoredFields)
// 		hideAd(params, records[i])
// 	}
//
// }
//
// func AcceptAd(token string, params []string, values []string, ignoredFields map[string]int) string {
// 	cmd := new(trans.Command)
// 	adId := 0
// 	for i, v := range params {
// 		switch v {
// 		case "type":
// 			cmd.AddParam(v, adType[values[i]])
// 		case "deviation", "list_id":
// 		case "status":
// 		default:
// 			if _, ok := ignoredFields[v]; !ok && values[i] != "" {
// 				if v == "ad_id" {
// 					adId, _ = strconv.Atoi(values[i])
// 				}
// 				cmd.AddParam(v, values[i])
// 			}
// 		}
// 	}
//
// 	cmd.AddParam("action_id", "1")
// 	r := spine.GetLastAcceptedActionIdRequest{
// 		AdId: adId,
// 	}
// 	res := spine.GetLastAcceptedActionIdResponse{}
// 	err := spineCli.GetLastAcceptedActionId(r, &res)
// 	if err == nil && res.ActionId != nil {
// 		cmd.AddParam("action_id", fmt.Sprintf("%v", *res.ActionId))
// 	}
//
// 	cmd.AddParam("do_not_send_mail", "1")
// 	cmd.AddParam("token", token)
// 	cmd.AddParam("remote_addr", "10.50.1.12")
// 	cmd.AddParam("utf8", "yes")
// 	cmd.AddParam("action", "accept")
//
// 	fmt.Println(cmd)
// 	token, err = cmd.SendCommand("review")
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	return token
// }
//
// func Authenticate(username string, passwd string) string {
// 	cmd := new(trans.Command)
// 	cmd.AddParam("username", username)
// 	cmd.AddParam("passwd", passwd)
// 	cmd.AddParam("remote_addr", "10.50.1.12")
//
// 	token, err := cmd.SendCommand("authenticate")
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println(token)
// 	return token
// }
//
// func hideAd(params []string, values []string) {
// 	cmd := new(trans.Command)
// 	for i, v := range params {
// 		switch v {
// 		case "status":
// 			if values[i] == "active" {
// 				return
// 			}
// 		case "ad_id":
// 			cmd.AddParam("ad_id", values[i])
// 		case "list_id":
// 			cmd.AddParam("list_id", values[i])
// 		}
// 	}
//
// 	cmd.AddParam("newstatus", "hidden")
// 	cmd.AddParam("hidden_reason", "2")
//
// 	fmt.Println(cmd)
//
// 	_, err := cmd.SendCommand("user_ad_status_change")
// 	if err != nil {
// 		fmt.Println(err)
// 	}
//
// }

// package main
//
// import (
// 	"fmt"
// 	"os"
// 	"path/filepath"
// )
//
// func run() ([]string, error) {
// 	searchDir := "./sql"
//
// 	fileList := make([]string, 0)
// 	e := filepath.Walk(searchDir, func(path string, f os.FileInfo, err error) error {
// 		if filepath.Ext(path) == ".tmpl" {
// 			fileList = append(fileList, path)
// 		}
//
// 		return err
// 	})
//
// 	if e != nil {
// 		panic(e)
// 	}
//
// 	for _, file := range fileList {
// 		fmt.Println(file)
// 	}
//
// 	return fileList, nil
// }
//
// func main() {
// 	run()
// }

//
// // Copyright 2015 The Prometheus Authors
// // Licensed under the Apache License, Version 2.0 (the "License");
// // you may not use this file except in compliance with the License.
// // You may obtain a copy of the License at
// //
// // http://www.apache.org/licenses/LICENSE-2.0
// //
// // Unless required by applicable law or agreed to in writing, software
// // distributed under the License is distributed on an "AS IS" BASIS,
// // WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// // See the License for the specific language governing permissions and
// // limitations under the License.
//
// // A minimal example of how to include Prometheus instrumentation.
// // package main
// //
// // import (
// // 	"flag"
// // 	"log"
// // 	"net/http"
// //
// // 	"github.com/prometheus/client_golang/prometheus/promhttp"
// // )
// //
// // var addr = flag.String("listen-address", ":8080", "The address to listen on for HTTP requests.")
// //
// // func main() {
// // 	flag.Parse()
// // 	http.Handle("/metrics", promhttp.Handler())
// // 	log.Fatal(http.ListenAndServe(*addr, nil))
// // }
//
// // package main
// //
// // import (
// // 	"encoding/json"
// // 	"fmt"
// // 	"io/ioutil"
// // 	"net/http"
// // 	"os"
// // 	"strings"
// // 	"time"
// // )
// //
// // func main() {
// // 	args := os.Args[1:]
// // 	// fmt.Println(args)
// // 	token := args[0]
// //
// // 	go tmp(token)
// //
// // 	for {
// // 		time.Sleep(1 * time.Second)
// // 	}
// // }
// //
// // func tmp(token string) {
// // 	url := "http://127.0.0.1:5657/v1/admin/validate_token"
// // 	// token := "X599e90685a20cc2c000000004d93a55300000000"
// // 	for {
// // 		time.Sleep(10 * time.Millisecond)
// // 		t := time.Now()
// // 		tmp := fmt.Sprintf("{\n    \"token\":\"%v\",\n    \"valid_minutes\": 60,\n    \"remote_addr\": \"127.0.0.1\"\n}", token)
// // 		payload := strings.NewReader(tmp)
// //
// // 		req, _ := http.NewRequest("POST", url, payload)
// //
// // 		req.Header.Add("content-type", "application/json")
// //
// // 		res, _ := http.DefaultClient.Do(req)
// //
// // 		defer res.Body.Close()
// // 		body, _ := ioutil.ReadAll(res.Body)
// // 		fmt.Println(string(body))
// //
// // 		m := make(map[string]interface{})
// // 		json.Unmarshal([]byte(body), &m)
// // 		fmt.Println(m)
// // 		token = fmt.Sprintf("%v", m["token"])
// //
// // 		fmt.Println(token)
// // 		fmt.Printf("runtime %f\n", time.Since(t).Seconds())
// // 	}
// //
// // }
// // package main
// //
// // import (
// // 	"database/sql"
// // 	"fmt"
// // 	"log"
// // 	"os"
// // 	"time"
// //
// // 	_ "github.com/lib/pq"
// // )
// //
// // var (
// // 	db  *sql.DB
// // 	err error
// // )
// //
// // func main() {
// // 	// args := os.Args[1:]
// // 	token := os.Args[1]
// //
// // 	db, err = sql.Open("postgres", "user=postgres dbname=blocketdb host=10.60.7.55 port=5432 sslmode=disable search_path=bpv,public")
// // 	if err != nil {
// // 		log.Fatal(err)
// // 	}
// //
// // 	db.SetMaxOpenConns(1)
// // 	db.SetMaxIdleConns(1)
// //
// // 	for {
// // 		t := time.Now()
// // 		token = do(token)
// // 		log.Printf("Runtime: %f", time.Since(t).Seconds())
// // 	}
// // }
// //
// // func do(token string) string {
// // 	q := fmt.Sprintf("SELECT * FROM validate_token('%v','127.0.0.1',INTERVAL '60 minutes','mamaphp_get_next_ad',NULL)", token)
// // 	rows, err := db.Query(q)
// // 	if err != nil {
// // 		log.Fatal(err)
// // 	}
// // 	defer rows.Close()
// // 	var (
// // 		adminId    int
// // 		t          string
// // 		remoteAddr *string
// // 		e          *string
// // 	)
// // 	for rows.Next() {
// // 		if err := rows.Scan(&adminId, &t, &remoteAddr, &e); err != nil {
// // 			log.Fatal(err)
// // 		}
// //
// // 		log.Println(adminId, t, remoteAddr, e)
// // 	}
// // 	return t
// // }

// package main
//
// import (
// 	"fmt"
// 	"io/ioutil"
// 	"net/http"
// 	"strings"
// 	"time"
// )
//
// func main() {
//
// 	for i := 0; i < 100000; i++ {
// 		start := time.Now()
// 		url := "http://127.0.0.1:13000/filter"
//
// 		payload := strings.NewReader("{\n\t\"ad_id\": 39492131,\n\t\"action_id\": 1,\n\t\"application\": \"clear\"\n}")
//
// 		req, _ := http.NewRequest("POST", url, payload)
//
// 		req.Header.Add("content-type", "application/json")
//
// 		res, _ := http.DefaultClient.Do(req)
//
// 		defer res.Body.Close()
// 		body, _ := ioutil.ReadAll(res.Body)
//
// 		fmt.Println(res)
// 		fmt.Println(string(body))
// 		fmt.Println(time.Since(start))
//
// 		time.Sleep(100 * time.Millisecond)
// 	}
//
// }

package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	spine "git.chotot.org/ad-service/spine/client"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("not enough params")
		return
	}
	spineURL := args[1]
	spineCli := spine.NewClient(nil, spine.BaseURL(spineURL))

	username := args[2]
	passwd := args[3]

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

	listID, _ := strconv.Atoi(args[4])
	token := authenticate(username, passwd, spineCli)
	for i := 1; i < len(records); i++ {
		phone := "0" + records[i][0]
		log.Printf("Add %v to list.", phone)
		if token, err = addBlockedItem(token, phone, listID, spineCli); err != nil {
			log.Println(err)
			token = authenticate(username, passwd, spineCli)
		}
	}
}

func authenticate(username, passwd string, spineCli *spine.Client) string {
	req := spine.AuthenticateRequest{
		Username: username,
		Password: passwd,
	}

	res := spine.AuthenticateResponse{}
	if err := spineCli.Authenticate(req, &res); err != nil {
		log.Println(err)
	}

	return res.Token
}

func addBlockedItem(token, value string, listID int, spineCli *spine.Client) (string, error) {
	req := spine.AddBlockedItemRequest{
		Token:  token,
		ListID: listID,
		Value:  value,
	}

	res := spine.BlockedItemActionResponse{}
	err := spineCli.AddBlockedItem(req, &res)
	if err != nil {
		log.Println(err)
	}
	return res.Token, err
}

// package main
//
// import (
// 	"fmt"
//
// 	cconf "git.chotot.org/ad-service/cconf/client"
// )
//
// func main() {
// 	configClient := cconf.NewClient(cconf.BaseURL("http://127.0.0.1:9090"))
//
// 	for {
// 		fmt.Println(configClient.GetString("category_settings.params.1.2010.s.value", ""))
// 	}
// }
