package onesignal

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"strconv"
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

func samplePlayerRequest() *PlayerRequest {
	return &PlayerRequest{
		Identifier:   "fake-identifier",
		Language:     "fake-language",
		Timezone:     -28800,
		GameVersion:  "1.0",
		DeviceOS:     "iOS",
		DeviceModel:  "iPhone5,2",
		AdID:         "fake-ad-id",
		SDK:          "fake-sdk",
		SessionCount: 1,
		Tags: map[string]string{
			"a":   "1",
			"foo": "bar",
		},
		AmountSpent:       0,
		CreatedAt:         1395096859,
		Playtime:          12,
		BadgeCount:        1,
		LastActive:        1395096859,
		TestType:          1,
		NotificationTypes: "2",
	}
}

func samplePlayerListResponse() string {
	return `{
		  "total_count":2,
		  "offset":0,
		  "limit":10,
		  "players":
		  [
		     {
					 "id": "id123",
					 "playtime": 0,
					 "sdk": "fake-sdk",
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
		       "amount_spent":0.0,
		       "created_at":1395096859,
		       "invalid_identifier":false,
		       "badge_count": 0
		     }
		  ]
	}`
}

func samplePlayer() *Player {
	return &Player{
		ID:           "id123",
		Playtime:     0,
		SDK:          "fake-sdk",
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
		AmountSpent:       0.0,
		CreatedAt:         1395096859,
		InvalidIdentifier: false,
		BadgeCount:        0,
	}
}

func TestList(t *testing.T) {
	setup()
	defer teardown()

	requestSent := false

	// PlayerListOptions
	opt := &PlayerListOptions{
		AppID:  "fake-app-id",
		Limit:  10,
		Offset: 0,
	}

	mux.HandleFunc("/players", func(w http.ResponseWriter, r *http.Request) {
		requestSent = true

		testMethod(t, r, "GET")

		// test URL/query string
		u, _ := url.Parse("/players")
		u.Scheme = ""
		u.Host = ""
		q := u.Query()
		q.Set("app_id", opt.AppID)
		q.Set("limit", strconv.Itoa(opt.Limit))
		q.Set("offset", strconv.Itoa(opt.Offset))
		u.RawQuery = q.Encode()
		want := u.String()
		if got := r.URL.String(); got != want {
			t.Errorf("URL: got %v, want %v", got, want)
		}

		fmt.Fprint(w, samplePlayerListResponse())
	})

	listRes, _, err := client.Players.List(opt)
	if err != nil {
		t.Errorf("List returned an error: %v", err)
	}

	player := *samplePlayer()
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

func TestList_returnsError(t *testing.T) {
	setup()
	defer teardown()

	// PlayerListOptions
	opt := &PlayerListOptions{
		AppID:  "fake-app-id",
		Limit:  10,
		Offset: 0,
	}

	mux.HandleFunc("/players", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, `{
		  "errors":
		  [
				"Invalid or missing authentication token"
		  ]
			}`)
	})

	_, resp, err := client.Players.List(opt)
	errResp, ok := err.(*ErrorResponse)
	if !ok {
		t.Errorf("Error should be of type ErrorResponse but is %v: %+v", reflect.TypeOf(err), err)
	}

	want := "Invalid or missing authentication token"
	if got := errResp.Messages[0]; want != got {
		t.Errorf("Error message: %v, want %v", got, want)
	}

	if got, want := resp.StatusCode, http.StatusBadRequest; want != got {
		t.Errorf("Status code: %d, want %d", got, want)
	}
}

func TestCreate(t *testing.T) {
	requestSent := false

	setup()
	defer teardown()

	playerRequest := samplePlayerRequest()

	mux.HandleFunc("/players", func(w http.ResponseWriter, r *http.Request) {
		requestSent = true

		testMethod(t, r, "POST")

		testBody(t, r, &PlayerRequest{}, playerRequest)

		fmt.Fprint(w, `{
			"success": true,
			"id": "ffffb794-ba37-11e3-8077-031d62f86ebf"
		}`)
	})

	createRes, _, _ := client.Players.Create(playerRequest)
	want := &PlayerCreateResponse{
		Success: true,
		ID:      "ffffb794-ba37-11e3-8077-031d62f86ebf",
	}
	if !reflect.DeepEqual(want, createRes) {
		t.Errorf("Request response: %+v, want %+v", createRes, want)
	}

	if requestSent == false {
		t.Errorf("Request has not been sent")
	}
}

func TestUpdate(t *testing.T) {
	requestSent := false

	setup()
	defer teardown()

	playerRequest := samplePlayerRequest()

	mux.HandleFunc("/players/fake-id", func(w http.ResponseWriter, r *http.Request) {
		requestSent = true

		testMethod(t, r, "PUT")

		testBody(t, r, &PlayerRequest{}, playerRequest)

		fmt.Fprint(w, `{
			"success": true
		}`)
	})

	updateRes, _, _ := client.Players.Update("fake-id", playerRequest)
	want := &PlayerUpdateResponse{
		Success: true,
	}
	if !reflect.DeepEqual(want, updateRes) {
		t.Errorf("Request response: %+v, want %+v", updateRes, want)
	}

	if requestSent == false {
		t.Errorf("Request has not been sent")
	}
}
