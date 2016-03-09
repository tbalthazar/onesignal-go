package onesignal

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestNewClient(t *testing.T) {
	c := NewClient(nil)

	if got, want := c.BaseURL.String(), defaultBaseURL; got != want {
		t.Errorf("NewClient BaseURL is %v, want %v", got, want)
	}

	if got, want := c.Client, http.DefaultClient; got != want {
		t.Errorf("NewClient Client is %v, want %v", got, want)
	}

	if got, want := c.Players.client, c; got != want {
		t.Errorf("NewClient.PlayersService.client is %v, want %v", got, want)
	}
}

func TestNewClient_withCustomHTTPClient(t *testing.T) {
	httpClient := &http.Client{}

	c := NewClient(httpClient)

	if got, want := c.Client, httpClient; got != want {
		t.Errorf("NewClient Client is %v, want %v", got, want)
	}
}

func TestNewRequest(t *testing.T) {
	appKey := "fake app key"
	userKey := "fake user key"
	c := NewClient(nil)
	c.AppKey = appKey
	c.UserKey = userKey

	method := "GET"
	inURL, outURL := "foo", defaultBaseURL+"foo"
	inBody := struct{ Foo string }{Foo: "Bar"}
	authKeyType := APP
	outBody := `{"Foo":"Bar"}` + "\n"
	req, _ := c.NewRequest(method, inURL, inBody, authKeyType)

	// test the HTTP method
	if got, want := req.Method, method; got != want {
		t.Errorf("NewRequest(%q) Method is %v, want %v", method, got, want)
	}

	// test the URL
	if got, want := req.URL.String(), outURL; got != want {
		t.Errorf("NewRequest(%q) URL is %v, want %v", inURL, got, want)
	}

	// test that body was JSON encoded
	body, _ := ioutil.ReadAll(req.Body)
	if got, want := string(body), outBody; got != want {
		t.Errorf("NewRequest(%q) Body is %v, want %v", inBody, got, want)
	}

	// test Content-Type header
	if got, want := req.Header.Get("Content-Type"), "application/json"; got != want {
		t.Errorf("NewRequest() Content-Type header is %v, want %v", got, want)
	}

	// test Accept header
	if got, want := req.Header.Get("Accept"), "application/json"; got != want {
		t.Errorf("NewRequest() Accept header is %v, want %v", got, want)
	}

	// test Authorization header
	if got, want := req.Header.Get("Authorization"), "Basic "+appKey; got != want {
		t.Errorf("NewRequest() Authorization header is %v, want %v", got, want)
	}
}

func TestNewRequest_userKeyType(t *testing.T) {
	appKey := "fake app key"
	userKey := "fake user key"
	c := NewClient(nil)
	c.AppKey = appKey
	c.UserKey = userKey

	req, _ := c.NewRequest("GET", "foo", nil, USER)

	// test Authorization header
	if got, want := req.Header.Get("Authorization"), "Basic "+userKey; got != want {
		t.Errorf("NewRequest() Authorization header is %v, want %v", got, want)
	}
}

func TestNewRequest_invalidJSON(t *testing.T) {
	c := NewClient(nil)

	type T struct {
		A map[int]interface{}
	}
	_, err := c.NewRequest("GET", "/", &T{}, APP)

	if err == nil {
		t.Error("Expected error to be returned.")
	}
	if err, ok := err.(*json.UnsupportedTypeError); !ok { // type assertion
		t.Errorf("Expected a UnsupportedTypeError; got %#v.", err)
	}
}

func TestNewRequest_emptyBody(t *testing.T) {
	c := NewClient(nil)

	req, err := c.NewRequest("GET", "/", nil, APP)

	if err != nil {
		t.Fatalf("NewRequest returned unexpected error: %v", err)
	}
	if req.Body != nil {
		t.Fatalf("Request contains a non-nil Body: %v", req.Body)
	}
}
