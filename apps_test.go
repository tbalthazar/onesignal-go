package onesignal

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
	"time"
)

func sampleAppListResponse() string {
	return `[
		{
		  "id": "92911750-242d-4260-9e00-9d9034f139ce",
		  "name": "Your app 1",
		  "players": 150,
		  "messagable_players": 143,
		  "updated_at": "2014-04-01T04:20:02.003Z",
		  "created_at": "2014-04-01T04:20:02.003Z",
		  "gcm_key": "a gcm push key",
		  "chrome_key": "A Chrome Web Push GCM key",
		  "chrome_web_origin": "Chrome Web Push Site URL",
		  "chrome_web_gcm_sender_id": "Chrome Web Push GCM Sender ID",
		  "chrome_web_default_notification_icon": "http://yoursite.com/chrome_notification_icon",
		  "chrome_web_sub_domain": "your_site_name",
		  "apns_env": "sandbox",
		  "apns_certificates": "Your apns certificate",
		  "safari_apns_cetificate": "Your Safari APNS certificate",
		  "safari_site_origin": "The homename for your website for Safari Push, including http or https",
		  "safari_push_id": "The certificate bundle ID for Safari Web Push",
		  "safari_icon_16_16": "http://onesignal.com/safari_packages/92911750-242d-4260-9e00-9d9034f139ce/16x16.png",
		  "safari_icon_32_32": "http://onesignal.com/safari_packages/92911750-242d-4260-9e00-9d9034f139ce/16x16@2.png",
		  "safari_icon_64_64": "http://onesignal.com/safari_packages/92911750-242d-4260-9e00-9d9034f139ce/32x32@2x.png",
		  "safari_icon_128_128": "http://onesignal.com/safari_packages/92911750-242d-4260-9e00-9d9034f139ce/128x128.png",
		  "safari_icon_256_256": "http://onesignal.com/safari_packages/92911750-242d-4260-9e00-9d9034f139ce/128x128@2x.png",
		  "site_name": "The URL to your website for Web Push",
		  "basic_auth_key": "NGEwMGZmMjItY2NkNy0xMWUzLTk5ZDUtMDAwYzI5NDBlNjJj"
		},
		{
		  "id": "e4e87830-b954-11e3-811d-f3b376925f15",
		  "name": "Your app 2",
		  "players": 100,
		  "messagable_players": 80,
		  "updated_at": "2014-04-01T04:20:02.003Z",
		  "created_at": "2014-04-01T04:20:02.003Z",
		  "gcm_key": "a gcm push key",
		  "chrome_key": "A Chrome Web Push GCM key",
		  "chrome_web_origin": "Chrome Web Push Site URL",
		  "chrome_web_gcm_sender_id": "Chrome Web Push GCM Sender ID",
		  "chrome_web_default_notification_icon": "http://yoursite.com/chrome_notification_icon",
		  "chrome_web_sub_domain": "your_site_name",
		  "apns_env": "sandbox",
		  "apns_certificates": "Your apns certificate",
		  "safari_apns_cetificate": "Your Safari APNS certificate",
		  "safari_site_origin": "The homename for your website for Safari Push, including http or https",
		  "safari_push_id": "The certificate bundle ID for Safari Web Push",
		  "safari_icon_16_16": "http://onesignal.com/safari_packages/e4e87830-b954-11e3-811d-f3b376925f15/16x16.png",
		  "safari_icon_32_32": "http://onesignal.com/safari_packages/e4e87830-b954-11e3-811d-f3b376925f15/16x16@2.png",
		  "safari_icon_64_64": "http://onesignal.com/safari_packages/e4e87830-b954-11e3-811d-f3b376925f15/32x32@2x.png",
		  "safari_icon_128_128": "http://onesignal.com/safari_packages/e4e87830-b954-11e3-811d-f3b376925f15/128x128.png",
		  "safari_icon_256_256": "http://onesignal.com/safari_packages/e4e87830-b954-11e3-811d-f3b376925f15/128x128@2x.png",
		  "site_name": "The URL to your website for Web Push",
		  "basic_auth_key": "NGEwMGZmMjItY2NkNy0xMWUzLTk5ZDUtMDAwYzI5NDBlNjJj"
		}
	]`
}

func sampleApp1() *App {
	t := time.Date(2014, time.April, 1, 4, 20, 2, 3000000, time.UTC)
	return &App{
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
}

func sampleApp2() *App {
	t := time.Date(2014, time.April, 1, 4, 20, 2, 3000000, time.UTC)
	return &App{
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
}

func TestAppsService_List(t *testing.T) {
	setup()
	defer teardown()

	requestSent := false

	mux.HandleFunc("/apps", func(w http.ResponseWriter, r *http.Request) {
		requestSent = true

		testMethod(t, r, "GET")

		fmt.Fprint(w, sampleAppListResponse())
	})

	apps, _, err := client.Apps.List()
	if err != nil {
		t.Errorf("List returned an error: %v", err)
	}

	want := []App{*sampleApp1(), *sampleApp2()}
	if !reflect.DeepEqual(apps, want) {
		t.Errorf("List returned %+v, want %+v", apps, want)
	}

	if requestSent == false {
		t.Errorf("Request has not been sent")
	}
}
