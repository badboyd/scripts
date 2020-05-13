package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"time"
)

// Client is spine client
type Client struct {
	BaseURL   *url.URL
	UserAgent string

	cache      map[string]string
	httpClient *http.Client
}

// CommonResponse default response for spine
type CommonResponse struct {
	Message string `json:"message,omitempty"`
}

// ErrorResponse error response for spine client
type ErrorResponse struct {
	Message      string `json:"message,omitempty"`
	DebugMessage string `json:"debug_message,omitempty"`
}

func (c ErrorResponse) Error() string {
	return c.Message
}

// ClientOption set spine client option
type ClientOption func(*Client)

var defaultTransport = &http.Transport{
	Dial: (&net.Dialer{
		Timeout:   30 * time.Second,
		KeepAlive: 30 * time.Second,
	}).Dial,
	MaxIdleConns:        100,
	MaxIdleConnsPerHost: 100,
	IdleConnTimeout:     90 * time.Second,
	TLSHandshakeTimeout: 10 * time.Second,
}

// ErrConnClosed connection close error
var ErrConnClosed = errors.New("CONNECTION_CLOSED")

var defaultHTTPClient = &http.Client{
	Transport: defaultTransport,
	Timeout:   30 * time.Second,
}

// NewClient create new spine client
func NewClient(httpClient *http.Client, options ...ClientOption) *Client {
	if httpClient == nil {
		httpClient = defaultHTTPClient
	}
	client := Client{httpClient: httpClient, cache: make(map[string]string)}

	for _, option := range options {
		option(&client)
	}

	return &client
}

// BaseURL set spine baseURL
func BaseURL(u string) func(*Client) {
	return func(c *Client) {
		if baseURL, err := url.Parse(u); err == nil {
			c.BaseURL = baseURL
		} else {
			panic(err)
		}
	}
}

// UserAgent set spine UserAgent
func UserAgent(u string) func(*Client) {
	return func(c *Client) {
		c.UserAgent = u
	}
}

// Should we add the token ?
func (c *Client) newRequest(method, path string, body interface{}) (*http.Request, error) {
	rel := &url.URL{Path: path}
	u := c.BaseURL.ResolveReference(rel)

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		if err := json.NewEncoder(buf).Encode(body); err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", c.UserAgent)
	req.Header.Set("Owner", "badboyd")

	return req, nil
}

func (c *Client) do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer func() {
		io.Copy(ioutil.Discard, resp.Body)
		resp.Body.Close()
	}()

	d := json.NewDecoder(resp.Body)
	d.UseNumber()

	if resp.StatusCode/100 != 2 {
		errRes := ErrorResponse{}
		d.Decode(&errRes)
		return resp, errRes
	}

	if v != nil {
		d.Decode(v)
	}

	return resp, err
}
