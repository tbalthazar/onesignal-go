package onesignal

import (
	"net/http"
	"net/url"
	"strconv"
)

// NotificationsService handles communication with the notification related
// methods of the OneSignal API.
type NotificationsService struct {
	client *Client
}

// Notification  represents a OneSignal notification.
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

// NotificationRequest represents a request to create a notification.
type NotificationRequest struct {
	AppID                  string            `json:"app_id"`
	Contents               map[string]string `json:"contents,omitempty"`
	Headings               map[string]string `json:"headings,omitempty"`
	IsIOS                  bool              `json:"isIos,omitempty"`
	IsAndroid              bool              `json:"isAndroid,omitempty"`
	IsWP                   bool              `json:"isWP,omitempty"`
	IsADM                  bool              `json:"isAdm,omitempty"`
	IsChrome               bool              `json:"isChrome,omitempty"`
	IsChromeWeb            bool              `json:"isChromeWeb,omitempty"`
	IsSafari               bool              `json:"isSafari,omitempty"`
	IsAnyWeb               bool              `json:"isAnyWeb,omitempty"`
	IncludedSegments       []string          `json:"included_segments,omitempty"`
	ExcludedSegments       []string          `json:"excluded_segments,omitempty"`
	IncludePlayerIDs       []string          `json:"include_player_ids,omitempty"`
	IncludeExternalUserIDs []string          `json:"include_external_user_ids,omitempty"`
	IncludeIOSTokens       []string          `json:"include_ios_tokens,omitempty"`
	IncludeAndroidRegIDs   []string          `json:"include_android_reg_ids,omitempty"`
	IncludeWPURIs          []string          `json:"include_wp_uris,omitempty"`
	IncludeWPWNSURIs       []string          `json:"include_wp_wns_uris,omitempty"`
	IncludeAmazonRegIDs    []string          `json:"include_amazon_reg_ids,omitempty"`
	IncludeChromeRegIDs    []string          `json:"include_chrome_reg_ids,omitempty"`
	IncludeChromeWebRegIDs []string          `json:"include_chrome_web_reg_ids,omitempty"`
	AppIDs                 []string          `json:"app_ids,omitempty"`
	Tags                   interface{}       `json:"tags,omitempty"`
	IOSBadgeType           string            `json:"ios_badgeType,omitempty"`
	IOSBadgeCount          int               `json:"ios_badgeCount,omitempty"`
	IOSSound               string            `json:"ios_sound,omitempty"`
	AndroidSound           string            `json:"android_sound,omitempty"`
	ADMSound               string            `json:"adm_sound,omitempty"`
	WPSound                string            `json:"wp_sound,omitempty"`
	WPWNSSound             string            `json:"wp_wns_sound,omitempty"`
	Data                   interface{}       `json:"data,omitempty"`
	Buttons                interface{}       `json:"buttons,omitempty"`
	SmallIcon              string            `json:"small_icon,omitempty"`
	LargeIcon              string            `json:"large_icon,omitempty"`
	BigPicture             string            `json:"big_picture,omitempty"`
	ADMSmallIcon           string            `json:"adm_small_icon,omitempty"`
	ADMLargeIcon           string            `json:"adm_large_icon,omitempty"`
	ADMBigPicture          string            `json:"adm_big_picture,omitempty"`
	ChromeIcon             string            `json:"chrome_icon,omitempty"`
	ChromeBigPicture       string            `json:"chrome_big_picture,omitempty"`
	ChromeWebIcon          string            `json:"chrome_web_icon,omitempty"`
	FirefoxIcon            string            `json:"firefox_icon,omitempty"`
	URL                    string            `json:"url,omitempty"`
	SendAfter              string            `json:"send_after,omitempty"`
	DelayedOption          string            `json:"delayed_option,omitempty"`
	DeliveryTimeOfDay      string            `json:"delivery_time_of_day,omitempty"`
	AndroidLEDColor        string            `json:"android_led_color,omitempty"`
	AndroidAccentColor     string            `json:"android_accent_color,omitempty"`
	AndroidVisibility      int               `json:"android_visibility,omitempty"`
	ContentAvailable       bool              `json:"content_available,omitempty"`
	AndroidBackgroundData  bool              `json:"android_background_data,omitempty"`
	AmazonBackgroundData   bool              `json:"amazon_background_data,omitempty"`
	TemplateID             string            `json:"template_id,omitempty"`
	AndroidGroup           string            `json:"android_group,omitempty"`
	AndroidGroupMessage    interface{}       `json:"android_group_message,omitempty"`
	ADMGroup               string            `json:"adm_group,omitempty"`
	ADMGroupMessage        interface{}       `json:"adm_group_message,omitempty"`
	Filters                interface{}       `json:"filters,omitempty"`
}

// NotificationCreateResponse wraps the standard http.Response for the
// NotificationsService.Create method
type NotificationCreateResponse struct {
	ID         string      `json:"id"`
	Recipients int         `json:"recipients"`
	Errors     interface{} `json:"errors"`
}

// NotificationListOptions specifies the parameters to the
// NotificationsService.List method
type NotificationListOptions struct {
	AppID  string `json:"app_id"`
	Limit  int    `json:"limit"`
	Offset int    `json:"offset"`
}

// NotificationListResponse wraps the standard http.Response for the
// NotificationsService.List method
type NotificationListResponse struct {
	TotalCount    int `json:"total_count"`
	Offset        int `json:"offset"`
	Limit         int `json:"limit"`
	Notifications []Notification
}

// NotificationUpdateOptions specifies the parameters to the
// NotificationsService.Get method
type NotificationGetOptions struct {
	AppID string `json:"app_id"`
}

// NotificationUpdateOptions specifies the parameters to the
// NotificationsService.Update method
type NotificationUpdateOptions struct {
	AppID  string `json:"app_id"`
	Opened bool   `json:"opened"`
}

// NotificationDeleteOptions specifies the parameters to the
// NotificationsService.Delete method
type NotificationDeleteOptions struct {
	AppID string `json:"app_id"`
}

// List the notifications.
//
// OneSignal API docs:
// https://documentation.onesignal.com/docs/notifications-view-notifications
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

// Get a single notification.
//
// OneSignal API docs:
// https://documentation.onesignal.com/docs/notificationsid-view-notification
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

// Create a notification.
//
// OneSignal API docs:
// https://documentation.onesignal.com/docs/notifications-create-notification
func (s *NotificationsService) Create(opt *NotificationRequest) (*NotificationCreateResponse, *http.Response, error) {
	// build the URL
	u, err := url.Parse("/notifications")
	if err != nil {
		return nil, nil, err
	}

	// create the request
	req, err := s.client.NewRequest("POST", u.String(), opt, APP)
	if err != nil {
		return nil, nil, err
	}

	createRes := &NotificationCreateResponse{}
	resp, err := s.client.Do(req, createRes)
	if err != nil {
		return nil, resp, err
	}

	return createRes, resp, err
}

// Update a notification.
//
// OneSignal API docs:
// https://documentation.onesignal.com/docs/notificationsid-track-open
func (s *NotificationsService) Update(notificationID string, opt *NotificationUpdateOptions) (*SuccessResponse, *http.Response, error) {
	// build the URL
	u, err := url.Parse("/notifications/" + notificationID)
	if err != nil {
		return nil, nil, err
	}

	// create the request
	req, err := s.client.NewRequest("PUT", u.String(), opt, APP)
	if err != nil {
		return nil, nil, err
	}

	updateRes := &SuccessResponse{}
	resp, err := s.client.Do(req, updateRes)
	if err != nil {
		return nil, resp, err
	}

	return updateRes, resp, err
}

// Delete a notification.
//
// OneSignal API docs:
// https://documentation.onesignal.com/docs/notificationsid-cancel-notification
func (s *NotificationsService) Delete(notificationID string, opt *NotificationDeleteOptions) (*SuccessResponse, *http.Response, error) {
	// build the URL
	u, err := url.Parse("/notifications/" + notificationID)
	if err != nil {
		return nil, nil, err
	}

	// create the request
	req, err := s.client.NewRequest("DELETE", u.String(), opt, APP)
	if err != nil {
		return nil, nil, err
	}

	deleteRes := &SuccessResponse{}
	resp, err := s.client.Do(req, deleteRes)
	if err != nil {
		return nil, resp, err
	}

	return deleteRes, resp, err
}
