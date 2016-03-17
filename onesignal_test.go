package onesignal

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"strings"
	"testing"
)

var (
	mux    *http.ServeMux
	server *httptest.Server
	client *Client
)

func setup() {
	// create a test server and a mux
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)

	// create a client, giving it the test server URL
	client = NewClient(nil)
	url, _ := url.Parse(server.URL)
	client.BaseURL = url
}

func teardown() {
	server.Close()
}

func testMethod(t *testing.T, r *http.Request, want string) {
	if got := r.Method; got != want {
		t.Errorf("Request method: %v, want %v", got, want)
	}
}

func testBody(t *testing.T, r *http.Request, body interface{}, want interface{}) {
	json.NewDecoder(r.Body).Decode(body)
	if !reflect.DeepEqual(body, want) {
		t.Errorf("Request body: %+v, want %+v", body, want)
	}
}

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

	if got, want := c.Apps.client, c; got != want {
		t.Errorf("NewClient.AppsService.client is %v, want %v", got, want)
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

func TestDo(t *testing.T) {
	setup()
	defer teardown()

	type foo struct {
		A string
	}

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"A":"a"}`)
	})

	req, _ := client.NewRequest("GET", "/", nil, APP)
	body := new(foo)
	client.Do(req, body)

	want := &foo{"a"}
	if !reflect.DeepEqual(body, want) {
		t.Errorf("Response body = %v, want %v", body, want)
	}
}

func TestDo_httpError(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Bad Request", 400)
	})

	req, _ := client.NewRequest("GET", "/", nil, APP)
	_, err := client.Do(req, nil)

	_, ok := err.(*ErrorResponse)
	if !ok {
		t.Errorf("Error should be of type ErrorResponse but is %v: %+v", reflect.TypeOf(err), err)
	}
}

func TestCheckResponse_ok(t *testing.T) {
	r := &http.Response{
		StatusCode: http.StatusOK,
	}

	err := CheckResponse(r)
	if err != nil {
		t.Fatalf("CheckResponse shouldn't return an error, but returned: %+v", err)
	}
}

func TestCheckResponse_badRequest(t *testing.T) {
	r := &http.Response{
		StatusCode: http.StatusBadRequest,
		Body: ioutil.NopCloser(strings.NewReader(`{
			"errors":
			[
				"Invalid or missing authentication token"
			]
		}`)),
	}

	err, ok := CheckResponse(r).(*ErrorResponse)
	if !ok {
		t.Errorf("CheckResponse return value should be of type ErrorResponse but is %v: %+v", reflect.TypeOf(err), err)
	}

	if err == nil {
		t.Fatalf("CheckResponse should return an error")
	}

	if len(err.Messages) == 0 {
		t.Fatalf("CheckResponse ErrorResponse should contain messages")
	}

	want := "Invalid or missing authentication token"
	if got := err.Messages[0]; want != got {
		t.Errorf("Error message: %v, want %v", got, want)
	}
}

func TestCheckResponse_noBody(t *testing.T) {
	r := &http.Response{
		StatusCode: http.StatusBadRequest,
		Body:       ioutil.NopCloser(strings.NewReader("")),
	}

	err, ok := CheckResponse(r).(*ErrorResponse)
	if !ok {
		t.Errorf("CheckResponse return value should be of type ErrorResponse but is %v: %+v", reflect.TypeOf(err), err)
	}

	if err == nil {
		t.Fatalf("CheckResponse should return an error")
	}

	if len(err.Messages) != 1 {
		t.Fatalf("CheckResponse ErrorResponse should contain 1 message")
	}

	want := "Couldn't decode response body JSON"
	if got := err.Messages[0]; want != got {
		t.Errorf("Error message: %v, want %v", got, want)
	}
}
