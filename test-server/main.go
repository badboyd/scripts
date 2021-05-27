// package main

// import (
// 	"io/ioutil"
// 	"log"
// 	"net/http"
// )

// func handler(w http.ResponseWriter, req *http.Request) {
// 	b, err := ioutil.ReadAll(req.Body)
// 	if err != nil {
// 		log.Println("Read body err: ", err.Error())
// 		w.WriteHeader(http.StatusBadRequest)
// 	}

// 	log.Printf("method %s request payload:%s ", req.Method, string(b))

// 	w.WriteHeader(http.StatusOK)
// }

// func main() {
// 	http.HandleFunc("/test", handler)

// 	http.ListenAndServe(":13001", nil)
// }

package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	url := "http://127.0.0.1:13000/api/v1/cmds"
	method := "POST"

	// 	payload := strings.NewReader(`{
	//     "protocol": "http",
	//     "executeIn": 30,
	//     "retryInterval": 10,
	//     "maxRetry": 3,
	//     "name": "angpau",
	//     "httpRequest": {
	//         "method": "POST",
	//         "path": "http://127.0.0.1:13001/test",
	//         "payload": "{\"bar\":\"baz\",\"foo\":\"1\"}"
	//     }
	// }`)

	b := bytes.NewReader([]byte(`
	{
    	"protocol": "http",
		"executeIn": 30,
		"retryInterval": 10,
		"maxRetry": 3,
		"name": "angpau",
		"httpRequest": {
			"method": "POST",
			"path": "http://127.0.0.1:13001/test",
			"payload": "{\"bar\":\"baz\",\"foo\":\"1\"}"
		}
	}`))

	client := &http.Client{}
	for i := 0; i < 1000; i++ {
		req, err := http.NewRequest(method, url, b)
		if err != nil {
			fmt.Println(err)
			return
		}
		req.Header.Add("Content-Type", "application/json")

		res, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer res.Body.Close()

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(body))
	}
}
