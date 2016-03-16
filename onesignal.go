package onesignal

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
)

const (
	defaultBaseURL = "https://onesignal.com/api/v1/"
)

type AuthKeyType uint

const (
	APP AuthKeyType = iota
	USER
)

type Client struct {
	BaseURL *url.URL
	AppKey  string
	UserKey string
	Client  *http.Client

	Players *PlayersService
}

type ErrorResponse struct {
	Messages []string `json:"errors"`
}

func (e *ErrorResponse) Error() string {
	msg := "OneSignal returned those error messages:\n - "
	return msg + strings.Join(e.Messages, "\n - ")
}

// NewClient returns a new OneSignal API client.
func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	baseURL, err := url.Parse(defaultBaseURL)
	if err != nil {
		log.Fatal(err)
	}

	c := &Client{
		BaseURL: baseURL,
		Client:  httpClient,
	}

	c.Players = &PlayersService{client: c}

	return c
}

// NewRequest creates an API request.
func (c *Client) NewRequest(method, path string, body interface{}, authKeyType AuthKeyType) (*http.Request, error) {
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

	if authKeyType == APP {
		req.Header.Add("Authorization", "Basic "+c.AppKey)
	} else {
		req.Header.Add("Authorization", "Basic "+c.UserKey)
	}

	return req, nil
}

func (c *Client) Do(r *http.Request, v interface{}) (*http.Response, error) {
	// send the request
	resp, err := c.Client.Do(r)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	err = CheckResponse(resp)
	if err != nil {
		return resp, err
	}

	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&v)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func CheckResponse(r *http.Response) error {
	switch r.StatusCode {
	case http.StatusOK:
		return nil
	case http.StatusInternalServerError:
		return &ErrorResponse{
			Messages: []string{"Internal Server Error"},
		}
	default:
		var errResp ErrorResponse
		dec := json.NewDecoder(r.Body)
		err := dec.Decode(&errResp)
		if err != nil {
			errResp.Messages = []string{"Couldn't decode response body JSON"}
		}
		return &errResp
	}
}
