package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"path"
	"time"
)

// Elastic represents vault which is backed by elastic search
type Elastic struct {
	baseURL string
	index   string
	idxType string
	client  *http.Client
}

// ElasticResult represents result from vault which is backed by elastic
type ElasticResult struct {
	ID     string `json:"_id"`
	Source Source `json:"_source"`
}

// Source ...
type Source struct {
	AdID           int    `json:"ad_id"`
	ActionID       int    `json:"action_id"`
	AccountID      string `json:"account_id"`
	ActionType     string `json:"action_type"`
	CurrentStateID int    `json:"current_state"`
	State          string `json:"state"`
	LockedBy       int    `json:"locked_by"`
	Queue          string `json:"queue"`
}

// NewESVault returns new *Elastic
func NewESVault(baseURL, idx, idxType string) *Elastic {
	return &Elastic{
		baseURL: baseURL,
		index:   idx,
		idxType: idxType,
		client: &http.Client{
			Timeout: time.Second * 30,
		},
	}
}

func newRequest(method, url string, body io.Reader) (*http.Request, error) {
	return http.NewRequest(method, url, body)
}

// Get val from es by key
func (e *Elastic) Get(key string) (string, error) {
	fullURL := e.baseURL + path.Join(e.index, e.idxType, key)
	req, err := newRequest("GET", fullURL, nil)
	if err != nil {
		return "", err
	}
	req.SetBasicAuth("esearch", "K7lBsDlv")
	res, err := e.client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	elasticRes := ElasticResult{}
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&elasticRes); err != nil {
		return "", err
	}
	fmt.Println(elasticRes)
	return elasticRes.Source.State, nil
}

// Del val from as by key
func (e *Elastic) Del(key string) (string, error) {
	return "", nil
}

// Set key-val to ES
func (e *Elastic) Set(key, val string, exp time.Duration) (string, error) {
	return "", nil
}

// Inc +1 for key
func (e *Elastic) Inc(key string) (int64, error) {
	return 0, nil
}
