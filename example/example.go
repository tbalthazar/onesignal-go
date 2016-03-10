package main

import (
	"fmt"
	"github.com/tbalthazar/onesignal-go"
	"log"
)

func main() {
	fmt.Println("--- hey")
	opt := &onesignal.PlayerListOptions{
		AppId:  "fake-app-id",
		Limit:  10,
		Offset: 0,
	}
	client := onesignal.NewClient(nil)
	listRes, res, err := client.Players.List(opt)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("--- listRes:%+v", listRes)
	fmt.Printf("--- res:%+v, err:%+v\n", res)
	fmt.Printf("--- err:%+v\n", err)
}
