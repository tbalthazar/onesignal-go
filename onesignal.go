package onesignal

import (
	"net/http"
	"net/url"
)

const (
	defaultBaseURL = "https://onesignal.com/api/v1"
)

type Client struct {
	BaseURL *url.URL
	Key     string
	Client  *http.Client
}

func NewClient(key string, client *http.Client) *Client {
	baseURL, _ := url.Parse(defaultBaseURL)
	c := &Client{
		BaseURL: baseURL,
		Key:     key,
		Client:  client,
	}

	return c
}
