package onesignal

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"strconv"
	"testing"
)

func TestList(t *testing.T) {
	requestSent := false

	// create a test server and a mux
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)
	defer server.Close()

	// create a client, giving it the test server URL
	key := "fake-key"
	client := NewClient(key, http.DefaultClient)
	url, _ := url.Parse(server.URL)
	client.BaseURL = url

	// PlayerListOptions
	opt := &PlayerListOptions{
		AppId:  "fake-app-id",
		Limit:  10,
		Offset: 0,
	}

	mux.HandleFunc("/players", func(w http.ResponseWriter, r *http.Request) {
		requestSent = true

		// test method
		want := "GET"
		if got := r.Method; got != want {
			t.Errorf("Request method: %v, want %v", got, want)
		}

		// test URL/query string
		u, _ := url.Parse("/players")
		u.Scheme = ""
		u.Host = ""
		q := u.Query()
		q.Set("app_id", opt.AppId)
		q.Set("limit", strconv.Itoa(opt.Limit))
		q.Set("offset", strconv.Itoa(opt.Limit))
		u.RawQuery = q.Encode()
		want = u.String()
		if got := r.URL.String(); got != want {
			t.Errorf("URL: got %v, want %v", got, want)
		}
		fmt.Fprint(w, `{
			"total_count":2,
			"offset":0,
			"limit":10
			}`)
	})

	res, err := client.Players.List(opt)
	if err != nil {
		t.Errorf("List returned an error: %v", err)
	}

	want := &PlayerListResponse{
		TotalCount: 2,
		Offset:     0,
		Limit:      10,
	}
	if !reflect.DeepEqual(res, want) {
		t.Errorf("List returned %+v, want %+v", res, want)
	}

	if requestSent == false {
		t.Errorf("Request has not been sent")
	}
}
