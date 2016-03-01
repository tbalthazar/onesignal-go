package onesignal

import (
	"encoding/json"
	"log"
	"net/url"
	"strconv"
)

type PlayersService struct {
	client *Client
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
}

func (s *PlayersService) List(opt *PlayerListOptions) (*PlayerListResponse, error) {
	// build the URL with the query string
	u, err := url.Parse("/players")
	if err != nil {
		return nil, err
	}
	q := u.Query()
	q.Set("app_id", opt.AppId)
	q.Set("limit", strconv.Itoa(opt.Limit))
	q.Set("offset", strconv.Itoa(opt.Limit))
	u.RawQuery = q.Encode()

	// create the request
	req, err := s.client.NewRequest("GET", u.String(), nil)
	if err != nil {
		log.Fatal("Do: ", err)
	}

	// send the request
	resp, err := s.client.Client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
	}
	defer resp.Body.Close()

	// return nil, nil

	var response PlayerListResponse
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
