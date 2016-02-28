package onesignal

import (
	"log"
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
	req, err := s.client.NewRequest("GET", "/players", nil)
	if err != nil {
		log.Fatal("Do: ", err)
	}

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
