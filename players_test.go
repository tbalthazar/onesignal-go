package onesignal

import (
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"testing"
)

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

func samplePlayerResponse() string {
	return `{
		"id": "id123",
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
		"playtime":0,
	  "amount_spent":0.0,
	  "created_at":1395096859,
	  "invalid_identifier":false,
	  "badge_count": 0,
		"sdk": "fake-sdk"
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

func samplePlayerOnSessionOptions() *PlayerOnSessionOptions {
	return &PlayerOnSessionOptions{
		Identifier:  "ce777617da7f548fe7a9ab6febb56cf39fba6d382000c0395666288d961ee566",
		Language:    "en",
		Timezone:    -28800,
		GameVersion: "1.0",
		DeviceOS:    "7.0.4",
		AdID:        "fake-ad-id",
		SDK:         "fake-sdk",
		Tags: map[string]string{
			"a":   "1",
			"foo": "bar",
		},
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
		testHeader(t, r, "Authorization", "Basic "+client.AppKey)

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

func TestPlayersService_Get(t *testing.T) {
	setup()
	defer teardown()

	requestSent := false

	mux.HandleFunc("/players/id123", func(w http.ResponseWriter, r *http.Request) {
		requestSent = true

		testMethod(t, r, "GET")
		testHeader(t, r, "Authorization", "Basic "+client.AppKey)

		fmt.Fprint(w, samplePlayerResponse())
	})

	player, _, err := client.Players.Get("id123")
	want := samplePlayer()

	if err != nil {
		t.Errorf("Shouldn't have returned an error: %+v", err)
	}

	if !reflect.DeepEqual(want, player) {
		t.Errorf("Request response: %+v, want %+v", player, want)
	}

	if requestSent == false {
		t.Errorf("Request has not been sent")
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
		testHeader(t, r, "Authorization", "Basic "+client.AppKey)

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

func TestPlayersService_OnSession(t *testing.T) {
	requestSent := false

	setup()
	defer teardown()

	opt := samplePlayerOnSessionOptions()

	mux.HandleFunc("/players/id123/on_session", func(w http.ResponseWriter, r *http.Request) {
		requestSent = true

		testMethod(t, r, "POST")
		testHeader(t, r, "Authorization", "Basic "+client.AppKey)

		testBody(t, r, &PlayerOnSessionOptions{}, opt)

		fmt.Fprint(w, `{
			"success": true
		}`)
	})

	onSessionRes, _, err := client.Players.OnSession("id123", opt)
	want := &PlayerOnSessionResponse{
		Success: true,
	}

	if err != nil {
		t.Errorf("Shouldn't have returned an error: %+v", err)
	}

	if !reflect.DeepEqual(want, onSessionRes) {
		t.Errorf("Request response: %+v, want %+v", onSessionRes, want)
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
		testHeader(t, r, "Authorization", "Basic "+client.AppKey)

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
