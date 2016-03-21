package onesignal

import (
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"testing"
)

func sampleNotificationListResponse() string {
	return `{
		"total_count":2,
	  "offset":0,
	  "limit":10,
	  "notifications":
	  [
	    {
	     "id":"481a2734-6b7d-11e4-a6ea-4b53294fa671",
	     "successful":15,
	     "failed":1,
	     "converted":3,
	     "remaining":0,
	     "queued_at":1415914655,
	     "send_after":1415914655,
	     "canceled": false,
	     "url": "https://yourWebsiteToOpen.com",
	     "data":null,
	     "headings":{
	       "en":"English and default langauge heading",
	       "es":"Spanish language heading"
	     },     
	     "contents":{
	       "en":"English and default content",
	       "es":"Hola"
	     }
	    },
	    {
	     "id":"b6b326a8-40aa-13e5-b91b-bf8bc3fa26f7",
	     "successful":5,
	     "failed":2,
	     "converted":0,
	     "remaining":0,
	     "queued_at":1415915123,
	     "send_after":1415915123,
	     "canceled": false,
			 "url": null,
	     "data":{
	       "foo":"bar",
	       "your":"custom metadata"
	     },
	     "headings":{
	       "en":"English and default langauge heading",
	       "es":"Spanish language heading"
	     },
	     "contents":{
	       "en":"English and default content",
	       "es":"Hola"
	     }
	    }
	  ]
	}`
}

func sampleNotification1() *Notification {
	return &Notification{
		ID:         "481a2734-6b7d-11e4-a6ea-4b53294fa671",
		Successful: 15,
		Failed:     1,
		Converted:  3,
		Remaining:  0,
		QueuedAt:   1415914655,
		SendAfter:  1415914655,
		Canceled:   false,
		URL:        "https://yourWebsiteToOpen.com",
		Data:       nil,
		Headings: map[string]string{
			"en": "English and default langauge heading",
			"es": "Spanish language heading",
		},
		Contents: map[string]string{
			"en": "English and default content",
			"es": "Hola",
		},
	}
}

func sampleNotification2() *Notification {
	return &Notification{
		ID:         "b6b326a8-40aa-13e5-b91b-bf8bc3fa26f7",
		Successful: 5,
		Failed:     2,
		Converted:  0,
		Remaining:  0,
		QueuedAt:   1415915123,
		SendAfter:  1415915123,
		Canceled:   false,
		// URL:        nil,
		Data: map[string]string{
			"foo":  "bar",
			"your": "custom metadata",
		},
		Headings: map[string]string{
			"en": "English and default langauge heading",
			"es": "Spanish language heading",
		},
		Contents: map[string]string{
			"en": "English and default content",
			"es": "Hola",
		},
	}
}

func TestNotificationsService_List(t *testing.T) {
	setup()
	defer teardown()

	requestSent := false

	// NotificationListOptions
	opt := &NotificationListOptions{
		AppID:  "fake-app-id",
		Limit:  10,
		Offset: 0,
	}

	mux.HandleFunc("/notifications", func(w http.ResponseWriter, r *http.Request) {
		requestSent = true

		testMethod(t, r, "GET")
		testHeader(t, r, "Authorization", "Basic "+client.AppKey)

		// test URL/query string
		u, _ := url.Parse("/notifications")
		u.Scheme = ""
		u.Host = ""
		q := u.Query()
		q.Set("app_id", opt.AppID)
		q.Set("limit", strconv.Itoa(opt.Limit))
		q.Set("offset", strconv.Itoa(opt.Offset))
		u.RawQuery = q.Encode()
		want := u.String()
		if got := r.URL.String(); got != want {
			t.Errorf("URL: got %v, want %v", got, want)
		}

		fmt.Fprint(w, sampleNotificationListResponse())
	})

	listRes, _, err := client.Notifications.List(opt)
	if err != nil {
		t.Errorf("List returned an error: %v", err)
	}

	notification1 := *sampleNotification1()
	notification2 := *sampleNotification2()
	want := &NotificationListResponse{
		TotalCount:    2,
		Offset:        0,
		Limit:         10,
		Notifications: []Notification{notification1, notification2},
	}
	if !reflect.DeepEqual(listRes, want) {
		t.Errorf("List returned %+v, want %+v", listRes, want)
	}

	if requestSent == false {
		t.Errorf("Request has not been sent")
	}
}
