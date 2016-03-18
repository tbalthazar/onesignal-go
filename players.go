package onesignal

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

type PlayersService struct {
	client *Client
}

type Player struct {
	ID                string            `json:"id"`
	Playtime          int               `json:"playtime"`
	SDK               string            `json:"sdk"`
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
	AppID             string            `json:"app_id,omitempty"`
	DeviceType        int               `json:"device_type,omitempty"`
	Identifier        string            `json:"identifier,omitempty"`
	Language          string            `json:"language,omitempty"`
	Timezone          int               `json:"timezone,omitempty"`
	GameVersion       string            `json:"game_version,omitempty"`
	DeviceOS          string            `json:"device_os,omitempty"`
	DeviceModel       string            `json:"device_model,omitempty"`
	AdID              string            `json:"ad_id,omitempty"`
	SDK               string            `json:"sdk,omitempty"`
	SessionCount      int               `json:"session_count,omitempty"`
	Tags              map[string]string `json:"tags,omitempty"`
	AmountSpent       float32           `json:"amount_spent,omitempty"`
	CreatedAt         int               `json:"created_at,omitempty"`
	Playtime          int               `json:"playtime,omitempty"`
	BadgeCount        int               `json:"badge_count,omitempty"`
	LastActive        int               `json:"last_active,omitempty"`
	TestType          int               `json:"test_type,omitempty"`
	NotificationTypes string            `json:"notification_types,omitempty"`
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

type PlayerUpdateResponse struct {
	Success bool `json:"success"`
}

type PlayerOnSessionOptions struct {
	Identifier  string            `json:"identifier,omitempty"`
	Language    string            `json:"language,omitempty"`
	Timezone    int               `json:"timezone,omitempty"`
	GameVersion string            `json:"game_version,omitempty"`
	DeviceOS    string            `json:"device_os,omitempty"`
	AdID        string            `json:"ad_id,omitempty"`
	SDK         string            `json:"sdk,omitempty"`
	Tags        map[string]string `json:"tags,omitempty"`
}

type PlayerOnSessionResponse struct {
	Success bool `json:"success"`
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
	q.Set("offset", strconv.Itoa(opt.Offset))
	u.RawQuery = q.Encode()

	// create the request
	req, err := s.client.NewRequest("GET", u.String(), nil, APP)
	if err != nil {
		return nil, nil, err
	}

	plResp := &PlayerListResponse{}
	resp, err := s.client.Do(req, plResp)
	if err != nil {
		return nil, resp, err
	}

	return plResp, resp, err
}

func (s *PlayersService) Get(playerID string) (*Player, *http.Response, error) {
	// build the URL
	path := fmt.Sprintf("/players/%s", playerID)
	u, err := url.Parse(path)
	if err != nil {
		return nil, nil, err
	}

	// create the request
	req, err := s.client.NewRequest("GET", u.String(), nil, APP)
	if err != nil {
		return nil, nil, err
	}

	plResp := new(Player)
	resp, err := s.client.Do(req, plResp)
	if err != nil {
		return nil, resp, err
	}
	plResp.ID = playerID

	return plResp, resp, err
}

func (s *PlayersService) Create(player *PlayerRequest) (*PlayerCreateResponse, *http.Response, error) {
	// build the URL
	u, err := url.Parse("/players")
	if err != nil {
		return nil, nil, err
	}

	// create the request
	req, err := s.client.NewRequest("POST", u.String(), player, APP)
	if err != nil {
		return nil, nil, err
	}

	plResp := &PlayerCreateResponse{}
	resp, err := s.client.Do(req, plResp)
	if err != nil {
		return nil, resp, err
	}

	return plResp, resp, err
}

func (s *PlayersService) OnSession(playerID string, opt *PlayerOnSessionOptions) (*PlayerOnSessionResponse, *http.Response, error) {
	// build the URL
	path := fmt.Sprintf("/players/%s/on_session", playerID)
	u, err := url.Parse(path)
	if err != nil {
		return nil, nil, err
	}

	// create the request
	req, err := s.client.NewRequest("POST", u.String(), opt, APP)
	if err != nil {
		return nil, nil, err
	}

	plResp := &PlayerOnSessionResponse{}
	resp, err := s.client.Do(req, plResp)
	if err != nil {
		return nil, resp, err
	}

	return plResp, resp, err
}

func (s *PlayersService) Update(playerID string, player *PlayerRequest) (*PlayerUpdateResponse, *http.Response, error) {
	// build the URL
	path := fmt.Sprintf("/players/%s", playerID)
	u, err := url.Parse(path)
	if err != nil {
		return nil, nil, err
	}

	// create the request
	req, err := s.client.NewRequest("PUT", u.String(), player, APP)
	if err != nil {
		return nil, nil, err
	}

	plResp := &PlayerUpdateResponse{}
	resp, err := s.client.Do(req, plResp)
	if err != nil {
		return nil, resp, err
	}

	return plResp, resp, err
}
