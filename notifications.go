package onesignal

import (
	"net/http"
	"net/url"
	"strconv"
)

type NotificationsService struct {
	client *Client
}

type Notification struct {
	ID         string            `json:"id"`
	Successful int               `json:"successful"`
	Failed     int               `json:"failed"`
	Converted  int               `json:"converted"`
	Remaining  int               `json:"remaining"`
	QueuedAt   int               `json:"queued_at"`
	SendAfter  int               `json:"send_after"`
	URL        string            `json:"url"`
	Data       interface{}       `json:"data"`
	Canceled   bool              `json:"canceled"`
	Headings   map[string]string `json:"headings"`
	Contents   map[string]string `json:"contents"`
}

// Options passed to the List method
type NotificationListOptions struct {
	AppID  string `json:"app_id"`
	Limit  int    `json:"limit"`
	Offset int    `json:"offset"`
}

// Response from the List method
type NotificationListResponse struct {
	TotalCount    int `json:"total_count"`
	Offset        int `json:"offset"`
	Limit         int `json:"limit"`
	Notifications []Notification
}

// Options passed to the Get method
type NotificationGetOptions struct {
	AppID string `json:"app_id"`
}

func (s *NotificationsService) List(opt *NotificationListOptions) (*NotificationListResponse, *http.Response, error) {
	// build the URL with the query string
	u, err := url.Parse("/notifications")
	if err != nil {
		return nil, nil, err
	}
	q := u.Query()
	q.Set("app_id", opt.AppID)
	q.Set("limit", strconv.Itoa(opt.Limit))
	q.Set("offset", strconv.Itoa(opt.Offset))
	u.RawQuery = q.Encode()

	// create the request
	req, err := s.client.NewRequest("GET", u.String(), nil, APP)
	if err != nil {
		return nil, nil, err
	}

	notifResp := &NotificationListResponse{}
	resp, err := s.client.Do(req, notifResp)
	if err != nil {
		return nil, resp, err
	}

	return notifResp, resp, err
}

func (s *NotificationsService) Get(notificationID string, opt *NotificationGetOptions) (*Notification, *http.Response, error) {
	// build the URL with the query string
	u, err := url.Parse("/notifications/" + notificationID)
	if err != nil {
		return nil, nil, err
	}
	q := u.Query()
	q.Set("app_id", opt.AppID)
	u.RawQuery = q.Encode()

	// create the request
	req, err := s.client.NewRequest("GET", u.String(), nil, APP)
	if err != nil {
		return nil, nil, err
	}

	notif := &Notification{}
	resp, err := s.client.Do(req, notif)
	if err != nil {
		return nil, resp, err
	}

	return notif, resp, err
}
