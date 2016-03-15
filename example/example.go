package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/tbalthazar/onesignal-go"
)

var (
	appID  string
	appKey string
)

func ListPlayers(client *onesignal.Client) {
	fmt.Println("### ListPlayers ###")
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
	fmt.Printf("--- nbPlayers: %d\n", len(listRes.Players))
	fmt.Println()
}

func CreatePlayer(client *onesignal.Client) (playerID string) {
	fmt.Println("### CreatePlayer ###")
	player := &onesignal.PlayerRequest{
		AppID:        appID,
		DeviceType:   1,
		Identifier:   "fakeidentifier2",
		Language:     "fake-language",
		Timezone:     -28800,
		GameVersion:  "1.0",
		DeviceOS:     "iOS",
		DeviceModel:  "iPhone5,2",
		AdID:         "fake-ad-id2",
		SDK:          "fake-sdk",
		SessionCount: 1,
		Tags: map[string]string{
			"a":   "1",
			"foo": "barr",
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
	fmt.Printf("--- createRes:%+v\n", createRes)
	fmt.Printf("--- playerID: %s\n", createRes.ID)
	fmt.Println()

	return createRes.ID
}

func UpdatePlayer(playerID string, client *onesignal.Client) {
	fmt.Println("### UpdatePlayer " + playerID + " ###")
	player := &onesignal.PlayerRequest{
		GameVersion: "1.3",
	}

	updateRes, res, err := client.Players.Update(playerID, player)
	if err != nil {
		fmt.Printf("--- res:%+v, err:%+v\n", res)
		log.Fatal(err)
	}
	fmt.Printf("--- updateRes:%+v\n", updateRes)
	fmt.Println()
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	appID = os.Getenv("ONESIGNAL_APP_ID")
	appKey = os.Getenv("ONESIGNAL_API_KEY")
	client := onesignal.NewClient(nil)
	client.AppKey = appKey

	ListPlayers(client)
	playerID := CreatePlayer(client)
	ListPlayers(client)
	UpdatePlayer(playerID, client)
	ListPlayers(client)
}
