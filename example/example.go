package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/tbalthazar/onesignal-go"
)

var (
	appID   string
	appKey  string
	userKey string
)

func ListApps(client *onesignal.Client) {
	fmt.Println("### ListApps ###")

	apps, res, err := client.Apps.List()
	if err != nil {
		fmt.Printf("--- res:%+v, err:%+v\n", res)
		log.Fatal(err)
	}
	fmt.Printf("--- apps:%+v\n", apps)
	fmt.Println()
}

func GetApps(appID string, client *onesignal.Client) {
	fmt.Println("### GetApps ###")

	app, res, err := client.Apps.Get(appID)
	if err != nil {
		fmt.Printf("--- res:%+v, err:%+v\n", res)
		log.Fatal(err)
	}
	fmt.Printf("--- app:%+v\n", app)
	fmt.Println()
}

func CreateApps(client *onesignal.Client) string {
	fmt.Println("### CreateApps ###")
	app := &onesignal.AppRequest{
		Name: "Your app 1",
	}
	createRes, res, err := client.Apps.Create(app)
	if err != nil {
		fmt.Printf("--- res:%+v, err:%+v\n", res)
		log.Fatal(err)
	}
	fmt.Printf("--- createRes:%+v\n", createRes)
	fmt.Printf("--- appID: %s\n", createRes.ID)
	fmt.Println()

	return createRes.ID
}

func UpdateApps(appID string, client *onesignal.Client) {
	fmt.Println("### UpdateApps ###")
	app := &onesignal.AppRequest{
		Name: "Your app 1 modified",
	}
	updateRes, res, err := client.Apps.Update(appID, app)
	if err != nil {
		fmt.Printf("--- res:%+v, err:%+v\n", res)
		log.Fatal(err)
	}
	fmt.Printf("--- updateRes:%+v\n", updateRes)
	fmt.Println()
}

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

func GetPlayer(playerID string, client *onesignal.Client) {
	fmt.Println("### GetPlayer " + playerID + " ###")

	player, res, err := client.Players.Get(playerID)
	if err != nil {
		fmt.Printf("--- res:%+v, err:%+v\n", res)
		log.Fatal(err)
	}
	fmt.Printf("--- player:%+v\n", player)
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

func OnSessionPlayer(playerID string, client *onesignal.Client) {
	fmt.Println("### OnSessionPlayer " + playerID + " ###")
	opt := &onesignal.PlayerOnSessionOptions{
		Identifier:  "ce777617da7f548fe7a9ab6febb56cf39fba6d382000c0395666288d961ee566",
		Language:    "en",
		Timezone:    -28800,
		GameVersion: "1.0",
		DeviceOS:    "7.0.4",
		AdID:        "fake-ad-id",
		SDK:         "fake-sdk",
		Tags: map[string]string{
			"a":   "1",
			"foo": "bar",
		},
	}

	onSessionRes, res, err := client.Players.OnSession(playerID, opt)
	if err != nil {
		fmt.Printf("--- res:%+v, err:%+v\n", res)
		log.Fatal(err)
	}
	fmt.Printf("--- onSessionRes:%+v\n", onSessionRes)
	fmt.Println()
}

func OnPurchasePlayer(playerID string, client *onesignal.Client) {
	fmt.Println("### OnPurchasePlayer " + playerID + " ###")
	p1 := onesignal.Purchase{
		SKU:    "foosku1",
		Amount: 1.99,
		ISO:    "BEL",
	}
	p2 := onesignal.Purchase{
		SKU:    "foosku2",
		Amount: 2.99,
		ISO:    "GER",
	}
	opt := &onesignal.PlayerOnPurchaseOptions{
		Purchases: []onesignal.Purchase{p1, p2},
		Existing:  true,
	}

	onPurchaseRes, res, err := client.Players.OnPurchase(playerID, opt)
	if err != nil {
		fmt.Printf("--- res:%+v, err:%+v\n", res)
		log.Fatal(err)
	}
	fmt.Printf("--- onPurchaseRes:%+v\n", onPurchaseRes)
	fmt.Println()
}

func OnFocusPlayer(playerID string, client *onesignal.Client) {
	fmt.Println("### OnFocusPlayer " + playerID + " ###")
	opt := &onesignal.PlayerOnFocusOptions{
		State:      "ping",
		ActiveTime: 60,
	}

	onFocusRes, res, err := client.Players.OnFocus(playerID, opt)
	if err != nil {
		fmt.Printf("--- res:%+v, err:%+v\n", res)
		log.Fatal(err)
	}
	fmt.Printf("--- onFocusRes:%+v\n", onFocusRes)
	fmt.Println()
}

func CSVExportPlayer(client *onesignal.Client) {
	fmt.Println("### CSVExportPlayer ###")
	opt := &onesignal.PlayerCSVExportOptions{
		AppID: appID,
	}

	CSVEXportRes, res, err := client.Players.CSVExport(opt)
	if err != nil {
		fmt.Printf("--- res:%+v, err:%+v\n", res)
		log.Fatal(err)
	}
	fmt.Printf("--- CSVEXportRes:%+v\n", CSVEXportRes)
	fmt.Println()
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

func ListNotifications(client *onesignal.Client) []onesignal.Notification {
	fmt.Println("### ListNotifications ###")
	listOpt := &onesignal.NotificationListOptions{
		AppID:  appID,
		Limit:  10,
		Offset: 0,
	}

	listRes, res, err := client.Notifications.List(listOpt)
	if err != nil {
		fmt.Printf("--- res:%+v, err:%+v\n", res)
		log.Fatal(err)
	}
	fmt.Printf("--- listRes:%+v\n", listRes)
	fmt.Printf("--- nbNotifications: %d\n", len(listRes.Notifications))
	fmt.Println()
	return listRes.Notifications
}

func GetNotifications(notificationID string, client *onesignal.Client) *onesignal.Notification {
	fmt.Println("### GetNotifications " + notificationID + " ###")
	opt := &onesignal.NotificationGetOptions{
		AppID: appID,
	}

	getRes, res, err := client.Notifications.Get(notificationID, opt)
	if err != nil {
		fmt.Printf("--- res:%+v, err:%+v\n", res)
		log.Fatal(err)
	}
	fmt.Printf("--- getRes:%+v\n", getRes)
	fmt.Println()
	return getRes
}

func CreateNotifications(client *onesignal.Client) string {
	fmt.Println("### CreateNotifications ###")
	playerID := "6c54762c-4daa-4166-8aa4-1c7f0b1ad78b" // valid
	// playerID := "83823c5f-53ce-4e35-be6a-a3f27e5d838f" // invalid
	notificationReq := &onesignal.NotificationRequest{
		AppID:            appID,
		Contents:         map[string]string{"en": "English message"},
		IsIOS:            true,
		IncludePlayerIDs: []string{playerID},
	}

	createRes, res, err := client.Notifications.Create(notificationReq)
	if err != nil {
		fmt.Printf("--- res:%+v, err:%+v\n", res)
		log.Fatal(err)
	}
	fmt.Printf("--- createRes:%+v\n", createRes)
	fmt.Println()

	return createRes.ID
}

func UpdateNotifications(notificationID string, client *onesignal.Client) {
	fmt.Println("### UpdateNotifications " + notificationID + " ###")
	opt := &onesignal.NotificationUpdateOptions{
		AppID:  appID,
		Opened: true,
	}

	updateRes, res, err := client.Notifications.Update(notificationID, opt)
	if err != nil {
		fmt.Printf("--- res:%+v, err:%+v\n", res)
		log.Fatal(err)
	}
	fmt.Printf("--- updateRes:%+v\n", updateRes)
	fmt.Println()
}

func DeleteNotifications(notificationID string, client *onesignal.Client) {
	fmt.Println("### DeleteNotifications " + notificationID + " ###")
	opt := &onesignal.NotificationDeleteOptions{
		AppID: appID,
	}

	deleteRes, res, err := client.Notifications.Delete(notificationID, opt)
	if err != nil {
		fmt.Printf("--- res:%+v, err:%+v\n", res)
		log.Fatal(err)
	}
	fmt.Printf("--- deleteRes:%+v\n", deleteRes)
	fmt.Println()
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	appID = os.Getenv("ONESIGNAL_APP_ID")
	appKey = os.Getenv("ONESIGNAL_API_KEY")
	userKey = os.Getenv("ONESIGNAL_USER_AUTH_KEY")
	client := onesignal.NewClient(nil)
	client.AppKey = appKey
	client.UserKey = userKey

	// apps
	// ListApps(client)
	// GetApps(appID, client)
	// appID = CreateApps(client)
	// GetApps(appID, client)
	// UpdateApps(appID, client)
	// GetApps(appID, client)

	// players
	// ListPlayers(client)
	// playerID := CreatePlayer(client)
	// GetPlayer(playerID, client)
	// OnSessionPlayer(playerID, client)
	// OnPurchasePlayer(playerID, client)
	// OnFocusPlayer(playerID, client)
	// GetPlayer(playerID, client)
	// ListPlayers(client)
	// CSVExportPlayer(client)
	// UpdatePlayer(playerID, client)
	// ListPlayers(client)

	// notifications
	// notif := ListNotifications(client)[0]
	// GetNotifications(notif.ID, client)
	notifID := CreateNotifications(client)
	// GetNotifications(notifID, client)
	// UpdateNotifications(notifID, client)
	// GetNotifications(notifID, client)
	DeleteNotifications(notifID, client)
}
