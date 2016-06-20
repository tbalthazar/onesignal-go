package onesignal

import (
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"testing"

	"github.com/tbalthazar/onesignal-go/testhelper"
)

var sampleNotification1 = &Notification{
	ID:         "481a2734-6b7d-11e4-a6ea-4b53294fa671",
	Successful: 15,
	Failed:     1,
	Converted:  3,
	Remaining:  0,
	QueuedAt:   1415914655,
	SendAfter:  1415914655,
	Canceled:   false,
	URL:        "https://yourWebsiteToOpen.com",
	Headings: map[string]string{
		"en": "English and default langauge heading",
		"es": "Spanish language heading",
	},
	Contents: map[string]string{
		"en": "English and default content",
		"es": "Hola",
	},
}

var sampleNotification2 = &Notification{
	ID:         "b6b326a8-40aa-13e5-b91b-bf8bc3fa26f7",
	Successful: 5,
	Failed:     2,
	Converted:  0,
	Remaining:  0,
	QueuedAt:   1415915123,
	SendAfter:  1415915123,
	Canceled:   false,
	Data: map[string]interface{}{
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

var sampleNotificationRequest = &NotificationRequest{
	AppID:            "id123",
	Contents:         map[string]string{"en": "English message"},
	IsIOS:            true,
	IncludePlayerIDs: []string{"playerid123"},
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

		fmt.Fprint(w, testhelper.LoadFixture(t, "notification-list-response.json"))
	})

	listRes, _, err := client.Notifications.List(opt)
	if err != nil {
		t.Errorf("List returned an error: %v", err)
	}

	notification1 := *sampleNotification1
	notification2 := *sampleNotification2
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

func TestNotificationsService_Get(t *testing.T) {
	setup()
	defer teardown()

	requestSent := false

	// NotificationGetOptions
	opt := &NotificationGetOptions{
		AppID: "fake-app-id",
	}
	notification := sampleNotification2

	mux.HandleFunc("/notifications/"+notification.ID, func(w http.ResponseWriter, r *http.Request) {
		requestSent = true

		testMethod(t, r, "GET")
		testHeader(t, r, "Authorization", "Basic "+client.AppKey)

		// test URL/query string
		u, _ := url.Parse("/notifications/" + notification.ID)
		u.Scheme = ""
		u.Host = ""
		q := u.Query()
		q.Set("app_id", opt.AppID)
		u.RawQuery = q.Encode()
		want := u.String()
		if got := r.URL.String(); got != want {
			t.Errorf("URL: got %v, want %v", got, want)
		}

		fmt.Fprint(w, testhelper.LoadFixture(t, "notification-get-response.json"))
	})

	getRes, _, err := client.Notifications.Get(notification.ID, opt)
	if err != nil {
		t.Errorf("Get returned an error: %v", err)
	}

	if !reflect.DeepEqual(getRes, notification) {
		t.Errorf("Get returned %+v, want %+v", getRes, notification)
	}

	if requestSent == false {
		t.Errorf("Request has not been sent")
	}
}

func TestNotificationsService_Create(t *testing.T) {
	setup()
	defer teardown()

	requestSent := false

	notificationRequest := sampleNotificationRequest

	mux.HandleFunc("/notifications", func(w http.ResponseWriter, r *http.Request) {
		requestSent = true

		testMethod(t, r, "POST")
		testHeader(t, r, "Authorization", "Basic "+client.AppKey)

		testBody(t, r, &NotificationRequest{}, notificationRequest)

		fmt.Fprint(w, `{
			"id": "notif-fake-id",
			"recipients": 1
		}`)
	})

	want := &NotificationCreateResponse{
		ID:         "notif-fake-id",
		Recipients: 1,
	}
	createRes, _, err := client.Notifications.Create(notificationRequest)
	if err != nil {
		t.Errorf("Create returned an error: %v", err)
	}

	if !reflect.DeepEqual(createRes, want) {
		t.Errorf("Get returned %+v, want %+v", createRes, want)
	}

	if requestSent == false {
		t.Errorf("Request has not been sent")
	}
}

func TestNotificationsService_Create_returnsError(t *testing.T) {
	setup()
	defer teardown()

	notificationRequest := sampleNotificationRequest

	mux.HandleFunc("/notifications", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, `{
		  "errors":
		  [
				"Notification content must not be null for any languages."
		  ]
			}`)
	})

	_, resp, err := client.Notifications.Create(notificationRequest)
	errResp, ok := err.(*ErrorResponse)
	if !ok {
		t.Errorf("Error should be of type ErrorResponse but is %v: %+v", reflect.TypeOf(err), err)
	}

	want := "Notification content must not be null for any languages."
	if got := errResp.Messages[0]; want != got {
		t.Errorf("Error message: %v, want %v", got, want)
	}

	if got, want := resp.StatusCode, http.StatusBadRequest; want != got {
		t.Errorf("Status code: %d, want %d", got, want)
	}
}

func TestNotificationsService_Create_invalidPlayerIds(t *testing.T) {
	setup()
	defer teardown()

	notificationRequest := sampleNotificationRequest

	mux.HandleFunc("/notifications", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{
			"errors": {
				"invalid_player_ids" : [
					"5fdc92b2-3b2a-11e5-ac13-8fdccfe4d986",
					"00cb73f8-5815-11e5-ba69-f75522da5528"
				]
			}
		}`)
	})

	createResp, _, err := client.Notifications.Create(notificationRequest)
	if err != nil {
		t.Errorf("Create returned an error: %v", err)
	}

	ids := []interface{}{
		"5fdc92b2-3b2a-11e5-ac13-8fdccfe4d986",
		"00cb73f8-5815-11e5-ba69-f75522da5528",
	}
	errors := map[string]interface{}{
		"invalid_player_ids": ids,
	}
	want := &NotificationCreateResponse{
		Errors: errors,
	}
	if !reflect.DeepEqual(createResp, want) {
		t.Errorf("Errors: %v, want %v", createResp, want)
	}
}

func TestNotificationsService_Create_noSubscribedPlayers(t *testing.T) {
	setup()
	defer teardown()

	notificationRequest := sampleNotificationRequest

	mux.HandleFunc("/notifications", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{
			"id": "",
			"recipients": 0,
			"errors": [
				"All included players are not subscribed"
			]
		}`)
	})

	createResp, _, err := client.Notifications.Create(notificationRequest)
	if err != nil {
		t.Errorf("Create returned an error: %v", err)
	}

	errors := []interface{}{
		"All included players are not subscribed",
	}
	want := &NotificationCreateResponse{
		ID:         "",
		Recipients: 0,
		Errors:     errors,
	}
	if !reflect.DeepEqual(createResp, want) {
		t.Errorf("Errors: %v, want %v", createResp, want)
	}
}

func TestNotificationsService_Update(t *testing.T) {
	setup()
	defer teardown()

	requestSent := false

	notifID := "notif-fake-id"
	opt := &NotificationUpdateOptions{
		AppID:  "id123",
		Opened: true,
	}

	mux.HandleFunc("/notifications/"+notifID, func(w http.ResponseWriter, r *http.Request) {
		requestSent = true

		testMethod(t, r, "PUT")
		testHeader(t, r, "Authorization", "Basic "+client.AppKey)

		testBody(t, r, &NotificationUpdateOptions{}, opt)

		fmt.Fprint(w, `{
			"success": true
		}`)
	})

	want := &SuccessResponse{
		Success: true,
	}
	updateRes, _, err := client.Notifications.Update(notifID, opt)
	if err != nil {
		t.Errorf("Update returned an error: %v", err)
	}

	if !reflect.DeepEqual(updateRes, want) {
		t.Errorf("Get returned %+v, want %+v", updateRes, want)
	}

	if requestSent == false {
		t.Errorf("Request has not been sent")
	}
}

func TestNotificationsService_Delete(t *testing.T) {
	setup()
	defer teardown()

	requestSent := false

	notifID := "notif-fake-id"
	opt := &NotificationDeleteOptions{
		AppID: "id123",
	}

	mux.HandleFunc("/notifications/"+notifID, func(w http.ResponseWriter, r *http.Request) {
		requestSent = true

		testMethod(t, r, "DELETE")
		testHeader(t, r, "Authorization", "Basic "+client.AppKey)

		testBody(t, r, &NotificationDeleteOptions{}, opt)

		fmt.Fprint(w, `{
			"success": true
		}`)
	})

	want := &SuccessResponse{
		Success: true,
	}
	deleteRes, _, err := client.Notifications.Delete(notifID, opt)
	if err != nil {
		t.Errorf("Delete returned an error: %v", err)
	}

	if !reflect.DeepEqual(deleteRes, want) {
		t.Errorf("Delete returned %+v, want %+v", deleteRes, want)
	}

	if requestSent == false {
		t.Errorf("Request has not been sent")
	}
}
