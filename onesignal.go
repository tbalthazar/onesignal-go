package onesignal

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

const (
	defaultBaseURL = "https://onesignal.com/api/v1/"
)

type Client struct {
	BaseURL *url.URL
	Key     string
	Client  *http.Client

	Players *PlayersService
}

// NewClient returns a new OneSignal API client.
func NewClient(key string, client *http.Client) *Client {
	baseURL, _ := url.Parse(defaultBaseURL)
	c := &Client{
		BaseURL: baseURL,
		Key:     key,
		Client:  client,
	}

	c.Players = &PlayersService{client: c}

	return c
}

// NewRequest creates an API request.
func (c *Client) NewRequest(method, path string, body interface{}) (*http.Request, error) {
	// build the URL
	u, err := url.Parse(c.BaseURL.String() + path)
	if err != nil {
		return nil, err
	}

	// JSON encode the body
	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	// create the request
	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	// headers
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", "Basic "+c.Key)

	return req, nil
}
