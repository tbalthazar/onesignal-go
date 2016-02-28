package onesignal

import (
	"net/http"
	"net/http/httptest"
	"net/url"
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

	mux.HandleFunc("/players", func(w http.ResponseWriter, r *http.Request) {
		requestSent = true
		want := "GET"
		if got := r.Method; got != want {
			t.Errorf("Request method: %v, want %v", got, want)
		}
		// fmt.Fprint(w, `[{"number":1}]`)
	})

	opt := &PlayerListOptions{
		AppId:  "fake-app-id",
		Limit:  10,
		Offset: 0,
	}
	client.Players.List(opt)

	if requestSent == false {
		t.Errorf("Request has not been sent")
	}
}
