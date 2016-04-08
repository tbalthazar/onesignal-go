/*
Package onesignal provides the binding for OneSignal API.

Create a new OneSignal client:

	client := onesignal.NewClient(nil)

Set the UserKey if you want to use the /apps endpoints:

	client.UserKey = "YourOneSignalUserKey"

Set the AppKey for the other endpoints:

	client.AppKey = "YourOneSignalAppKey"

List Apps

	apps, res, err := client.Apps.List()

Get App

	app, res, err := client.Apps.Get("YourAppID")

Create App

	appRequest := &onesignal.AppRequest{
		Name: "Your app 1",
	}
	app, res, err := client.Apps.Create(appRequest )

Update App

	appRequest := &onesignal.AppRequest{
		Name: "Your app 1 modified",
	}
	app, res, err := client.Apps.Update(appID, appRequest)

*/
package onesignal
