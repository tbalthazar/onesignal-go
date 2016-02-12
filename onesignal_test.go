package onesignal

import (
	"io/ioutil"
	"net/http"
	"testing"
)

func TestNewClient(t *testing.T) {
	key := "fake key"
	httpClient := &http.Client{}

	c := NewClient(key, httpClient)

	if got, want := c.BaseURL.String(), defaultBaseURL; got != want {
		t.Errorf("NewClient BaseURL is %v, want %v", got, want)
	}

	if got, want := c.Key, key; got != want {
		t.Errorf("NewClient Key is %v, want %v", got, want)
	}

	if got, want := c.Client, httpClient; got != want {
		t.Errorf("NewClient Client is %v, want %v", got, want)
	}
}

func TestNewRequest(t *testing.T) {
	key := "fake key"
	httpClient := &http.Client{}

	c := NewClient(key, httpClient)

	method := "GET"
	inURL, outURL := "foo", defaultBaseURL+"foo"
	inBody := struct{ Foo string }{Foo: "Bar"}
	outBody := `{"Foo":"Bar"}` + "\n"
	req, _ := c.NewRequest(method, inURL, inBody)

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
	if got, want := req.Header.Get("Authorization"), "Basic "+c.Key; got != want {
		t.Errorf("NewRequest() Authorization header is %v, want %v", got, want)
	}

}
