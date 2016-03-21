package onesignal

import (
	"net/http"
	"net/url"
	"time"
)

type AppsService struct {
	client *Client
}

type App struct {
	ID                               string    `json:"id"`
	Name                             string    `json:"name"`
	Players                          int       `json:"players"`
	MessagablePlayers                int       `json:"messagable_players"`
	UpdatedAt                        time.Time `json:"updated_at"`
	CreatedAt                        time.Time `json:"created_at"`
	GCMKey                           string    `json:"gcm_key"`
	ChromeKey                        string    `json:"chrome_key"`
	ChromeWebOrigin                  string    `json:"chrome_web_origin"`
	ChromeWebGCMSenderID             string    `json:"chrome_web_gcm_sender_id"`
	ChromeWebDefaultNotificationIcon string    `json:"chrome_web_default_notification_icon"`
	ChromeWebSubDomain               string    `json:"chrome_web_sub_domain"`
	APNSEnv                          string    `json:"apns_env"`
	APNSCertificates                 string    `json:"apns_certificates"`
	SafariAPNSCertificate            string    `json:"safari_apns_cetificate"`
	SafariSiteOrigin                 string    `json:"safari_site_origin"`
	SafariPushID                     string    `json:"safari_push_id"`
	SafariIcon1616                   string    `json:"safari_icon_16_16"`
	SafariIcon3232                   string    `json:"safari_icon_32_32"`
	SafariIcon6464                   string    `json:"safari_icon_64_64"`
	SafariIcon128128                 string    `json:"safari_icon_128_128"`
	SafariIcon256256                 string    `json:"safari_icon_256_256"`
	SiteName                         string    `json:"site_name"`
	BasicAuthKey                     string    `json:"basic_auth_key"`
}

type AppRequest struct {
	Name                             string `json:"name",omitempty`
	GCMKey                           string `json:"gcm_key,omitempty"`
	ChromeKey                        string `json:"chrome_key,omitempty"`
	ChromeWebOrigin                  string `json:"chrome_web_origin,omitempty"`
	ChromeWebGCMSenderID             string `json:"chrome_web_gcm_sender_id,omitempty"`
	ChromeWebDefaultNotificationIcon string `json:"chrome_web_default_notification_icon,omitempty"`
	ChromeWebSubDomain               string `json:"chrome_web_sub_domain,omitempty"`
	APNSEnv                          string `json:"apns_env,omitempty"`
	APNSCertificates                 string `json:"apns_certificates,omitempty"`
	SafariAPNSCertificate            string `json:"safari_apns_cetificate,omitempty"`
	SafariSiteOrigin                 string `json:"safari_site_origin,omitempty"`
	SafariPushID                     string `json:"safari_push_id,omitempty"`
	SafariIcon1616                   string `json:"safari_icon_16_16,omitempty"`
	SafariIcon3232                   string `json:"safari_icon_32_32,omitempty"`
	SafariIcon6464                   string `json:"safari_icon_64_64,omitempty"`
	SafariIcon128128                 string `json:"safari_icon_128_128,omitempty"`
	SafariIcon256256                 string `json:"safari_icon_256_256,omitempty"`
	SiteName                         string `json:"site_name,omitempty"`
	BasicAuthKey                     string `json:"basic_auth_key,omitempty"`
}

func (s *AppsService) List() ([]App, *http.Response, error) {
	// build the URL
	u, err := url.Parse("/apps")
	if err != nil {
		return nil, nil, err
	}

	// create the request
	req, err := s.client.NewRequest("GET", u.String(), nil, USER)
	if err != nil {
		return nil, nil, err
	}

	apps := new([]App)
	resp, err := s.client.Do(req, apps)
	if err != nil {
		return nil, resp, err
	}

	return *apps, resp, err
}

func (s *AppsService) Get(appID string) (*App, *http.Response, error) {
	// build the URL
	u, err := url.Parse("/apps/" + appID)
	if err != nil {
		return nil, nil, err
	}

	// create the request
	req, err := s.client.NewRequest("GET", u.String(), nil, USER)
	if err != nil {
		return nil, nil, err
	}

	app := &App{}
	resp, err := s.client.Do(req, app)
	if err != nil {
		return nil, resp, err
	}

	return app, resp, err
}

func (s *AppsService) Create(opt *AppRequest) (*App, *http.Response, error) {
	// build the URL
	u, err := url.Parse("/apps")
	if err != nil {
		return nil, nil, err
	}

	// create the request
	req, err := s.client.NewRequest("POST", u.String(), opt, USER)
	if err != nil {
		return nil, nil, err
	}

	app := &App{}
	resp, err := s.client.Do(req, app)
	if err != nil {
		return nil, resp, err
	}

	return app, resp, err
}

func (s *AppsService) Update(appID string, opt *AppRequest) (*App, *http.Response, error) {
	// build the URL
	u, err := url.Parse("/apps/" + appID)
	if err != nil {
		return nil, nil, err
	}

	// create the request
	req, err := s.client.NewRequest("PUT", u.String(), opt, USER)
	if err != nil {
		return nil, nil, err
	}

	app := &App{}
	resp, err := s.client.Do(req, app)
	if err != nil {
		return nil, resp, err
	}

	return app, resp, err
}
