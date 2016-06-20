package onesignal

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
	"time"

	"github.com/tbalthazar/onesignal-go/testhelper"
)

var t = time.Date(2014, time.April, 1, 4, 20, 2, 3000000, time.UTC)
var sampleApp1 = &App{
	ID:                               "92911750-242d-4260-9e00-9d9034f139ce",
	Name:                             "Your app 1",
	Players:                          150,
	MessagablePlayers:                143,
	UpdatedAt:                        t,
	CreatedAt:                        t,
	GCMKey:                           "a gcm push key",
	ChromeKey:                        "A Chrome Web Push GCM key",
	ChromeWebOrigin:                  "Chrome Web Push Site URL",
	ChromeWebGCMSenderID:             "Chrome Web Push GCM Sender ID",
	ChromeWebDefaultNotificationIcon: "http://yoursite.com/chrome_notification_icon",
	ChromeWebSubDomain:               "your_site_name",
	APNSEnv:                          "sandbox",
	APNSCertificates:                 "Your apns certificate",
	SafariAPNSCertificate:            "Your Safari APNS certificate",
	SafariSiteOrigin:                 "The homename for your website for Safari Push, including http or https",
	SafariPushID:                     "The certificate bundle ID for Safari Web Push",
	SafariIcon1616:                   "http://onesignal.com/safari_packages/92911750-242d-4260-9e00-9d9034f139ce/16x16.png",
	SafariIcon3232:                   "http://onesignal.com/safari_packages/92911750-242d-4260-9e00-9d9034f139ce/16x16@2.png",
	SafariIcon6464:                   "http://onesignal.com/safari_packages/92911750-242d-4260-9e00-9d9034f139ce/32x32@2x.png",
	SafariIcon128128:                 "http://onesignal.com/safari_packages/92911750-242d-4260-9e00-9d9034f139ce/128x128.png",
	SafariIcon256256:                 "http://onesignal.com/safari_packages/92911750-242d-4260-9e00-9d9034f139ce/128x128@2x.png",
	SiteName:                         "The URL to your website for Web Push",
	BasicAuthKey:                     "NGEwMGZmMjItY2NkNy0xMWUzLTk5ZDUtMDAwYzI5NDBlNjJj",
}

var sampleApp2 = &App{
	ID:                               "e4e87830-b954-11e3-811d-f3b376925f15",
	Name:                             "Your app 2",
	Players:                          100,
	MessagablePlayers:                80,
	UpdatedAt:                        t,
	CreatedAt:                        t,
	GCMKey:                           "a gcm push key",
	ChromeKey:                        "A Chrome Web Push GCM key",
	ChromeWebOrigin:                  "Chrome Web Push Site URL",
	ChromeWebGCMSenderID:             "Chrome Web Push GCM Sender ID",
	ChromeWebDefaultNotificationIcon: "http://yoursite.com/chrome_notification_icon",
	ChromeWebSubDomain:               "your_site_name",
	APNSEnv:                          "sandbox",
	APNSCertificates:                 "Your apns certificate",
	SafariAPNSCertificate:            "Your Safari APNS certificate",
	SafariSiteOrigin:                 "The homename for your website for Safari Push, including http or https",
	SafariPushID:                     "The certificate bundle ID for Safari Web Push",
	SafariIcon1616:                   "http://onesignal.com/safari_packages/e4e87830-b954-11e3-811d-f3b376925f15/16x16.png",
	SafariIcon3232:                   "http://onesignal.com/safari_packages/e4e87830-b954-11e3-811d-f3b376925f15/16x16@2.png",
	SafariIcon6464:                   "http://onesignal.com/safari_packages/e4e87830-b954-11e3-811d-f3b376925f15/32x32@2x.png",
	SafariIcon128128:                 "http://onesignal.com/safari_packages/e4e87830-b954-11e3-811d-f3b376925f15/128x128.png",
	SafariIcon256256:                 "http://onesignal.com/safari_packages/e4e87830-b954-11e3-811d-f3b376925f15/128x128@2x.png",
	SiteName:                         "The URL to your website for Web Push",
	BasicAuthKey:                     "NGEwMGZmMjItY2NkNy0xMWUzLTk5ZDUtMDAwYzI5NDBlNjJj",
}

var sampleAppRequest = &AppRequest{
	Name:                             "Your app 1",
	GCMKey:                           "a gcm push key",
	ChromeKey:                        "A Chrome Web Push GCM key",
	ChromeWebOrigin:                  "Chrome Web Push Site URL",
	ChromeWebGCMSenderID:             "Chrome Web Push GCM Sender ID",
	ChromeWebDefaultNotificationIcon: "http://yoursite.com/chrome_notification_icon",
	ChromeWebSubDomain:               "your_site_name",
	APNSEnv:                          "sandbox",
	SafariSiteOrigin:                 "The homename for your website for Safari Push, including http or https",
	SafariIcon1616:                   "http://onesignal.com/safari_packages/92911750-242d-4260-9e00-9d9034f139ce/16x16.png",
	SafariIcon3232:                   "http://onesignal.com/safari_packages/92911750-242d-4260-9e00-9d9034f139ce/16x16@2.png",
	SafariIcon6464:                   "http://onesignal.com/safari_packages/92911750-242d-4260-9e00-9d9034f139ce/32x32@2x.png",
	SafariIcon128128:                 "http://onesignal.com/safari_packages/92911750-242d-4260-9e00-9d9034f139ce/128x128.png",
	SafariIcon256256:                 "http://onesignal.com/safari_packages/92911750-242d-4260-9e00-9d9034f139ce/128x128@2x.png",
	SiteName:                         "The URL to your website for Web Push",
}

func TestAppsService_List(t *testing.T) {
	setup()
	defer teardown()

	requestSent := false

	mux.HandleFunc("/apps", func(w http.ResponseWriter, r *http.Request) {
		requestSent = true

		testMethod(t, r, "GET")
		testHeader(t, r, "Authorization", "Basic "+client.UserKey)

		fmt.Fprint(w, testhelper.LoadFixture(t, "app-list-response.json"))
	})

	apps, _, err := client.Apps.List()
	if err != nil {
		t.Errorf("List returned an error: %v", err)
	}

	want := []App{*sampleApp1, *sampleApp2}
	if !reflect.DeepEqual(apps, want) {
		t.Errorf("List returned %+v, want %+v", apps, want)
	}

	if requestSent == false {
		t.Errorf("Request has not been sent")
	}
}

func TestAppsService_Get(t *testing.T) {
	setup()
	defer teardown()

	requestSent := false
	appID := "92911750-242d-4260-9e00-9d9034f139ce"

	mux.HandleFunc("/apps/"+appID, func(w http.ResponseWriter, r *http.Request) {
		requestSent = true

		testMethod(t, r, "GET")
		testHeader(t, r, "Authorization", "Basic "+client.UserKey)

		fmt.Fprint(w, testhelper.LoadFixture(t, "app-get-response.json"))
	})

	app, _, err := client.Apps.Get(appID)
	if err != nil {
		t.Errorf("Get returned an error: %v", err)
	}

	want := sampleApp1
	if !reflect.DeepEqual(app, want) {
		t.Errorf("Get returned %+v, want %+v", app, want)
	}

	if requestSent == false {
		t.Errorf("Request has not been sent")
	}
}

func TestAppsService_Create(t *testing.T) {
	setup()
	defer teardown()

	requestSent := false
	appRequest := sampleAppRequest

	mux.HandleFunc("/apps", func(w http.ResponseWriter, r *http.Request) {
		requestSent = true

		testMethod(t, r, "POST")
		testHeader(t, r, "Authorization", "Basic "+client.UserKey)

		testBody(t, r, &AppRequest{}, appRequest)

		fmt.Fprint(w, testhelper.LoadFixture(t, "app-get-response.json"))
	})

	createRes, _, err := client.Apps.Create(appRequest)
	if err != nil {
		t.Errorf("Create returned an error: %v", err)
	}

	want := sampleApp1
	if !reflect.DeepEqual(want, createRes) {
		t.Errorf("Request response: %+v, want %+v", createRes, want)
	}

	if requestSent == false {
		t.Errorf("Request has not been sent")
	}
}

func TestAppsService_Update(t *testing.T) {
	setup()
	defer teardown()

	requestSent := false
	appRequest := sampleAppRequest
	appID := "id123"

	mux.HandleFunc("/apps/"+appID, func(w http.ResponseWriter, r *http.Request) {
		requestSent = true

		testMethod(t, r, "PUT")
		testHeader(t, r, "Authorization", "Basic "+client.UserKey)

		testBody(t, r, &AppRequest{}, appRequest)

		fmt.Fprint(w, testhelper.LoadFixture(t, "app-get-response.json"))
	})

	updateRes, _, err := client.Apps.Update(appID, appRequest)
	if err != nil {
		t.Errorf("Update returned an error: %v", err)
	}

	want := sampleApp1
	if !reflect.DeepEqual(want, updateRes) {
		t.Errorf("Request response: %+v, want %+v", updateRes, want)
	}

	if requestSent == false {
		t.Errorf("Request has not been sent")
	}
}
