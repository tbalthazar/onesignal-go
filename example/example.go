package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/tbalthazar/onesignal-go"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	appID := os.Getenv("ONESIGNAL_APP_ID")
	appKey := os.Getenv("ONESIGNAL_API_KEY")
	client := onesignal.NewClient(nil)
	client.AppKey = appKey

	// --- List
	listOpt := &onesignal.PlayerListOptions{
		AppID:  appID,
		Limit:  10,
		Offset: 0,
	}

	listRes, res, err := client.Players.List(listOpt)
	if err != nil {
		fmt.Printf("--- res:%+v, err:%+v\n", res)
		log.Fatal(err)
	}
	fmt.Printf("--- listRes:%+v\n", listRes)
	fmt.Printf("--- nbPlayers: %d\n\n", len(listRes.Players))

	// --- Create
	player := &onesignal.PlayerRequest{
		AppID:        appID,
		DeviceType:   0,
		Identifier:   "fakeidentifier",
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
		AmountSpent: 1.99,
		CreatedAt:   1395096859,
		Playtime:    12,
		BadgeCount:  1,
		LastActive:  1395096859,
		TestType:    1,
	}
	createRes, res, err := client.Players.Create(player)
	if err != nil {
		fmt.Printf("--- res:%+v, err:%+v\n", res)
		log.Fatal(err)
	}
	fmt.Printf("--- createRes:%+v\n\n", createRes)

	// --- List
	listRes, res, err = client.Players.List(listOpt)
	if err != nil {
		fmt.Printf("--- res: %+v\n", res)
		log.Fatal(err)
	}
	fmt.Printf("--- listRes:%+v\n", listRes)
	fmt.Printf("--- nbPlayers: %d\n\n", len(listRes.Players))
}
