/*
Package onesignal provides the binding for OneSignal API.

Create a new OneSignal client:

	client := onesignal.NewClient(nil)

Set the UserKey if you want to use the /apps endpoints:

	client.UserKey = "YourOneSignalUserKey"

Set the AppKey for the other endpoints:

	client.AppKey = "YourOneSignalAppKey"

Apps

List apps:

	apps, res, err := client.Apps.List()

Get an app:

	app, res, err := client.Apps.Get("YourAppID")

Create an app:

	appRequest := &onesignal.AppRequest{
		Name: "Your app 1",
	}
	app, res, err := client.Apps.Create(appRequest )

Update an app:

	appRequest := &onesignal.AppRequest{
		Name: "Your app 1 modified",
	}
	app, res, err := client.Apps.Update(appID, appRequest)

Players

List players:

	listOpt := &onesignal.PlayerListOptions{
		AppID:  appID,
		Limit:  10,
		Offset: 0,
	}
	listRes, res, err := client.Players.List(listOpt)

Get player:

	player, res, err := client.Players.Get("playerID")

Create player:

	playerRequest := &onesignal.PlayerRequest{
		AppID:        "appID",
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
			"foo": "bar",
		},
		AmountSpent: 1.99,
		CreatedAt:   1395096859,
		Playtime:    12,
		BadgeCount:  1,
		LastActive:  1395096859,
		TestType:    1,
	}
	createRes, res, err := client.Players.Create(playerRequest)

Create a new session for a player:

	opt := &onesignal.PlayerOnSessionOptions{
		Identifier:  "FakeIdentifier",
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
	successRes, res, err := client.Players.OnSession(playerID, opt)

Create a new purchase for a player:

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
	successRes, res, err := client.Players.OnPurchase(playerID, opt)

Increment the total session length for a player:

	opt := &onesignal.PlayerOnFocusOptions{
		State:      "ping",
		ActiveTime: 60,
	}
	successRes, res, err := client.Players.OnFocus(playerID, opt)

Generate a link to download a CSV list of all the players:

	opt := &onesignal.PlayerCSVExportOptions{
		AppID: appID,
	}
	CSVEXportRes, res, err := client.Players.CSVExport(opt)

Update a player:

	player := &onesignal.PlayerRequest{
		GameVersion: "1.3",
	}
	successRes, res, err := client.Players.Update(playerID, player)

Notifications

List notifications:

	listOpt := &onesignal.NotificationListOptions{
		AppID:  appID,
		Limit:  10,
		Offset: 0,
	}
	listRes, res, err := client.Notifications.List(listOpt)

Get a notification:

	opt := &onesignal.NotificationGetOptions{
		AppID: appID,
	}
	notification, res, err := client.Notifications.Get(notificationID, opt)

Create a notification:

	playerID := "aPlayerID"
	notificationReq := &onesignal.NotificationRequest{
		AppID:            appID,
		Contents:         map[string]string{"en": "English message"},
		IsIOS:            true,
		IncludePlayerIDs: []string{playerID},
	}
	createRes, res, err := client.Notifications.Create(notificationReq)

Update a notification:

	opt := &onesignal.NotificationUpdateOptions{
		AppID:  appID,
		Opened: true,
	}
	successRes, res, err := client.Notifications.Update(notificationID, opt)

Delete a notification:

	opt := &onesignal.NotificationDeleteOptions{
		AppID: appID,
	}
	successRes, res, err := client.Notifications.Delete(notificationID, opt)

*/
package onesignal
