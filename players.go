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

type PlayerRequest struct {
	AppID        string            `json:"app_id"`
	DeviceType   int               `json:"device_type"`
	Identifier   string            `json:"identifier"`
	Language     string            `json:"language"`
	Timezone     int               `json:"timezone"`
	GameVersion  string            `json:"game_version"`
	DeviceOS     string            `json:"device_os"`
	DeviceModel  string            `json:"device_model"`
	AdID         string            `json:"ad_id"`
	SDK          string            `json:"sdk"`
	SessionCount int               `json:"session_count"`
	Tags         map[string]string `json:"tags"`
	AmountSpent  float32           `json:"amount_spent"`
	CreatedAt    int               `json:"created_at"`
	Playtime     int               `json:"playtime"`
	BadgeCount   int               `json:"badge_count"`
	LastActive   int               `json:"last_active"`
	TestType     int               `json:"test_type"`
}

type PlayerListOptions struct {
	AppID  string `json:"app_id"`
	Limit  int    `json:"limit"`
	Offset int    `json:"offset"`
}

type PlayerListResponse struct {
	TotalCount int `json:"total_count"`
	Offset     int `json:"offset"`
	Limit      int `json:"limit"`
	Players    []Player
}

type PlayerCreateResponse struct {
	Success bool   `json:"success"`
	ID      string `json:"id"`
}

func (s *PlayersService) List(opt *PlayerListOptions) (*PlayerListResponse, *http.Response, error) {
	// build the URL with the query string
	u, err := url.Parse("/players")
	if err != nil {
		return nil, nil, err
	}
	q := u.Query()
	q.Set("app_id", opt.AppID)
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

	err = CheckResponse(resp)
	if err != nil {
		return nil, resp, err
	}

	var plResp PlayerListResponse
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&plResp)
	if err != nil {
		return nil, nil, err
	}
	return &plResp, resp, nil
}

func (s *PlayersService) Create(player *PlayerRequest) (*PlayerCreateResponse, *http.Response, error) {
	// build the URL with the query string
	u, err := url.Parse("/players")
	if err != nil {
		return nil, nil, err
	}

	// create the request
	req, err := s.client.NewRequest("POST", u.String(), player, APP)
	if err != nil {
		return nil, nil, err
	}

	// send the request
	resp, err := s.client.Client.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()

	err = CheckResponse(resp)
	if err != nil {
		return nil, resp, err
	}

	var plResp PlayerCreateResponse
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&plResp)
	if err != nil {
		return nil, resp, err
	}
	return &plResp, resp, nil
}
