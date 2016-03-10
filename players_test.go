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
	client := NewClient(nil)
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
		  "limit":10,
		  "players":
		  [
		     {
		       "identifier":"ce777617da7f548fe7a9ab6febb56cf39fba6d382000c0395666288d961ee566",
		       "session_count":1,
		       "language":"en",
		       "timezone":-28800,
		       "game_version":"1.0",
		       "device_os":"7.0.4",
		       "device_type":0,
		       "device_model":"iPhone",
		       "ad_id":null,
		       "tags":{"a":"1","foo":"bar"},
		       "last_active":1395096859,
		       "amount_spent":0,
		       "created_at":1395096859,
		       "invalid_identifier":false,
		       "badge_count": 0
		     }
		  ]
			}`)
	})

	listRes, _, err := client.Players.List(opt)
	if err != nil {
		t.Errorf("List returned an error: %v", err)
	}

	player := Player{
		Identifier:   "ce777617da7f548fe7a9ab6febb56cf39fba6d382000c0395666288d961ee566",
		SessionCount: 1,
		Language:     "en",
		Timezone:     -28800,
		GameVersion:  "1.0",
		DeviceOS:     "7.0.4",
		DeviceType:   0,
		DeviceModel:  "iPhone",
		Tags: map[string]string{
			"a":   "1",
			"foo": "bar",
		},
		LastActive:        1395096859,
		AmountSpent:       0,
		CreatedAt:         1395096859,
		InvalidIdentifier: false,
		BadgeCount:        0,
	}
	want := &PlayerListResponse{
		TotalCount: 2,
		Offset:     0,
		Limit:      10,
		Players:    []Player{player},
	}
	if !reflect.DeepEqual(listRes, want) {
		t.Errorf("List returned %+v, want %+v", listRes, want)
	}

	if requestSent == false {
		t.Errorf("Request has not been sent")
	}
}
