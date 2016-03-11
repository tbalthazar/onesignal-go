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

	opt := &onesignal.PlayerListOptions{
		AppId:  appID,
		Limit:  10,
		Offset: 0,
	}
	client := onesignal.NewClient(nil)
	client.AppKey = appKey

	listRes, res, err := client.Players.List(opt)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("--- listRes:%+v\n", listRes)
	fmt.Printf("--- res:%+v, err:%+v\n", res)
}
