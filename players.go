package onesignal

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
)

type PlayersService struct {
	client *Client
}

type Player struct {
	Identifier        string            `json:"identifier"`
	SessionCount      int               `json:"session_count"`
	Language          string            `json:"language"`
	Timezone          int               `json:"timezone"`
	GameVersion       string            `json:"game_version"`
	DeviceOS          string            `json:"device_os"`
	DeviceType        int               `json:"device_type"`
	DeviceModel       string            `json:"device_model"`
	AdID              string            `json:"ad_id"`
	Tags              map[string]string `json:"tags"`
	LastActive        int               `json:"last_active"`
	AmountSpent       float32           `json:"amount_spent"`
	CreatedAt         int               `json:"created_at"`
	InvalidIdentifier bool              `json:"invalid_identifier"`
	BadgeCount        int               `json:"badge_count"`
}

type PlayerListOptions struct {
	AppId  string `json:"app_id"`
	Limit  int    `json:"limit"`
	Offset int    `json:"offset"`
}

type PlayerListResponse struct {
	TotalCount int `json:"total_count"`
	Offset     int `json:"offset"`
	Limit      int `json:"limit"`
	Players    []Player
}

func (s *PlayersService) List(opt *PlayerListOptions) (*PlayerListResponse, *http.Response, error) {
	// build the URL with the query string
	u, err := url.Parse("/players")
	if err != nil {
		return nil, nil, err
	}
	q := u.Query()
	q.Set("app_id", opt.AppId)
	q.Set("limit", strconv.Itoa(opt.Limit))
	q.Set("offset", strconv.Itoa(opt.Limit))
	u.RawQuery = q.Encode()

	// create the request
	req, err := s.client.NewRequest("GET", u.String(), nil, APP)
	if err != nil {
		return nil, nil, err
	}

	// send the request
	resp, err := s.client.Client.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		var response PlayerListResponse
		dec := json.NewDecoder(resp.Body)
		err = dec.Decode(&response)
		if err != nil {
			return nil, nil, err
		}
		return &response, resp, nil
	} else {
		var errResp ErrorResponse
		dec := json.NewDecoder(resp.Body)
		err := dec.Decode(&errResp)
		if err != nil {
			return nil, nil, err
		}
		return nil, resp, &errResp
	}
}
