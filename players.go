package onesignal

type PlayersService struct {
	client *Client
}

type PlayerListOptions struct {
	AppId  string `json:"app_id"`
	limit  int    `json:"limit"`
	offset int    `json:"offset"`
}
