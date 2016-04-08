// Package onesignal provides the binding for OneSignal API.
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

// AuthKeyType specifies the token used to authentify the requests
type AuthKeyType uint

const (
	APP AuthKeyType = iota
	USER
)

// A Client manages communication with the OneSignal API.
type Client struct {
	BaseURL *url.URL
	AppKey  string
	UserKey string
	Client  *http.Client

	Apps          *AppsService
	Players       *PlayersService
	Notifications *NotificationsService
}

// SuccessResponse  wraps the standard http.Response for several API methods
// that just return a Success flag.
type SuccessResponse struct {
	Success bool `json:"success"`
}

// ErrorResponse reports one or more errors caused by an API request.
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

	c.Apps = &AppsService{client: c}
	c.Players = &PlayersService{client: c}
	c.Notifications = &NotificationsService{client: c}

	return c
}

// NewRequest creates an API request. path is a relative URL, like "/apps". The
// value pointed to by body is JSON encoded and included as the request body.
// The AuthKeyType will determine which authorization token (APP or USER) is
// used for the request.
func (c *Client) NewRequest(method, path string, body interface{}, authKeyType AuthKeyType) (*http.Request, error) {
	// build the URL
	u, err := url.Parse(c.BaseURL.String() + path)
	if err != nil {
		return nil, err
	}

	// JSON encode the body
	var buf io.ReadWriter
	if body != nil {
		b := new(bytes.Buffer)
		err := json.NewEncoder(b).Encode(body)
		if err != nil {
			return nil, err
		}
		buf = b
		// log.Println("Body is: " + b.String())
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
		// log.Println("Authorization header is AppKey")
	} else {
		req.Header.Add("Authorization", "Basic "+c.UserKey)
		// log.Println("Authorization header is UserKey")
	}

	return req, nil
}

// Do sends an API request and returns the API response. The API response is
// JSON decoded and stored in the value pointed to by v, or returned as an
// error if an API error has occurred.
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

	// // log body for debug
	// b := new(bytes.Buffer)
	// b.ReadFrom(resp.Body)
	// log.Println("response body: ", b.String())

	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&v)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

// CheckResponse checks the API response for errors, and returns them if
// present.  A response is considered an error if it has a status code outside
// the 200 range.  API error responses are expected to have either no response
// body, or a JSON response body that maps to ErrorResponse.  Any other
// response body will be silently ignored.
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
