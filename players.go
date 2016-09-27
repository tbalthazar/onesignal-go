package onesignal

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

// PlayersService handles communication with the player related
// methods of the OneSignal API.
type PlayersService struct {
	client *Client
}

// Player represents a OneSignal player.
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

// PlayerRequest represents a request to create/update a player.
type PlayerRequest struct {
	AppID             string            `json:"app_id"`
	DeviceType        int               `json:"device_type"`
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

// PlayerListOptions specifies the parameters to the PlayersService.List method
type PlayerListOptions struct {
	AppID  string `json:"app_id"`
	Limit  int    `json:"limit"`
	Offset int    `json:"offset"`
}

// PlayerListResponse wraps the standard http.Response for the
// PlayersService.List method
type PlayerListResponse struct {
	TotalCount int `json:"total_count"`
	Offset     int `json:"offset"`
	Limit      int `json:"limit"`
	Players    []Player
}

// PlayerCreateResponse wraps the standard http.Response for the
// PlayersService.Create method
type PlayerCreateResponse struct {
	Success bool   `json:"success"`
	ID      string `json:"id"`
}

// PlayerOnSessionOptions specifies the parameters to the
// PlayersService.OnSession method
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

// Purchase represents a purchase in the options of the
// PlayersService.OnPurchase method
type Purchase struct {
	SKU    string  `json:"sku"`
	Amount float32 `json:"amount"`
	ISO    string  `json:"iso"`
}

// PlayerOnPurchaseOptions specifies the parameters to the
// PlayersService.OnPurchase method
type PlayerOnPurchaseOptions struct {
	Purchases []Purchase `json:"purchases"`
	Existing  bool       `json:"existing,omitempty"`
}

// PlayerOnFocusOptions specifies the parameters to the
// PlayersService.OnFocus method
type PlayerOnFocusOptions struct {
	State      string `json:"state"`
	ActiveTime int    `json:"active_time"`
}

// PlayerCSVExportOptions specifies the parameters to the
// PlayersService.CSVExport method
type PlayerCSVExportOptions struct {
	AppID string `json:"app_id"`
}

// PlayerCSVExportResponse wraps the standard http.Response for the
// PlayersService.CSVExport method
type PlayerCSVExportResponse struct {
	CSVFileURL string `json:"csv_file_url"`
}

// List the players.
//
// OneSignal API docs: https://documentation.onesignal.com/docs/players-view-devices
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

// Get a single player.
//
// OneSignal API docs: https://documentation.onesignal.com/docs/playersid
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

// Create a player.
//
// OneSignal API docs:
// https://documentation.onesignal.com/docs/players-add-a-device
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

// Create a new session for a player.
//
// OneSignal API docs:
// https://documentation.onesignal.com/docs/playersidon_session
func (s *PlayersService) OnSession(playerID string, opt *PlayerOnSessionOptions) (*SuccessResponse, *http.Response, error) {
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

	plResp := &SuccessResponse{}
	resp, err := s.client.Do(req, plResp)
	if err != nil {
		return nil, resp, err
	}

	return plResp, resp, err
}

// Create a new purchase for a player.
//
// OneSignal API docs: https://documentation.onesignal.com/docs/on_purchase
func (s *PlayersService) OnPurchase(playerID string, opt *PlayerOnPurchaseOptions) (*SuccessResponse, *http.Response, error) {
	// build the URL
	path := fmt.Sprintf("/players/%s/on_purchase", playerID)
	u, err := url.Parse(path)
	if err != nil {
		return nil, nil, err
	}

	// create the request
	req, err := s.client.NewRequest("POST", u.String(), opt, APP)
	if err != nil {
		return nil, nil, err
	}

	plResp := &SuccessResponse{}
	resp, err := s.client.Do(req, plResp)
	if err != nil {
		return nil, resp, err
	}

	return plResp, resp, err
}

// Increment the total session length for a player.
//
// OneSignal API docs:
// https://documentation.onesignal.com/docs/playersidon_focus
func (s *PlayersService) OnFocus(playerID string, opt *PlayerOnFocusOptions) (*SuccessResponse, *http.Response, error) {
	// build the URL
	path := fmt.Sprintf("/players/%s/on_focus", playerID)
	u, err := url.Parse(path)
	if err != nil {
		return nil, nil, err
	}

	// create the request
	req, err := s.client.NewRequest("POST", u.String(), opt, APP)
	if err != nil {
		return nil, nil, err
	}

	plResp := &SuccessResponse{}
	resp, err := s.client.Do(req, plResp)
	if err != nil {
		return nil, resp, err
	}

	return plResp, resp, err
}

// Generate a link to download a CSV list of all the players.
//
// OneSignal API docs:
// https://documentation.onesignal.com/docs/players_csv_export
func (s *PlayersService) CSVExport(opt *PlayerCSVExportOptions) (*PlayerCSVExportResponse, *http.Response, error) {
	// build the URL with the query string
	u, err := url.Parse("/players/csv_export")
	if err != nil {
		return nil, nil, err
	}
	q := u.Query()
	q.Set("app_id", opt.AppID)
	u.RawQuery = q.Encode()

	// create the request
	req, err := s.client.NewRequest("POST", u.String(), opt, APP)
	if err != nil {
		return nil, nil, err
	}

	plResp := &PlayerCSVExportResponse{}
	resp, err := s.client.Do(req, plResp)
	if err != nil {
		return nil, resp, err
	}

	return plResp, resp, err
}

// Update a player.
//
// OneSignal API docs: https://documentation.onesignal.com/docs/playersid-1
func (s *PlayersService) Update(playerID string, player *PlayerRequest) (*SuccessResponse, *http.Response, error) {
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

	plResp := &SuccessResponse{}
	resp, err := s.client.Do(req, plResp)
	if err != nil {
		return nil, resp, err
	}

	return plResp, resp, err
}
