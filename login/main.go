package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Token struct {
	Token     string `json:"token"`
	AccountID int    `json:"account_id"`
}

func main() {
	args := os.Args[1:]
	if len(args) < 3 {
		fmt.Println("not enough params")
		return
	}
	loginUrl := args[0]
	userName := args[1]
	passWord := args[2]
	login(userName, passWord, loginUrl)
}

func login(username string, pw string, loginUrl string) (token string, err error) {
	values := map[string]string{"phone": username, "password": pw}
	jsonValue, _ := json.Marshal(values)
	resp, err := http.Post(loginUrl, "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	var s Token
	json.Unmarshal(body, &s)
	return s.Token, err
}
