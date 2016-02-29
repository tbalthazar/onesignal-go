package onesignal

import (
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

func (s *PlayersService) List(opt *PlayerListOptions) {
	// build the URL with the query string
	u, err := url.Parse("/players")
	if err != nil {
		log.Fatal(err)
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
	_, err = s.client.Client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
	}

	// defer resp.Body.Close()

	// var playerList PlayerList
	// dec := json.NewDecoder(resp.Body)
	// err = dec.Decode(&playerList)
	// if err != nil {
	// 	log.Fatal("JSON Decode error: ", err)
	// }

	// return &playerList
}
