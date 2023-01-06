package models

type Item struct {
	ID          string `gorm:"primary" json:"id"`
	UserSubId   string `json:"user_sub"`
	AccessToken string `json:"access_token"`
	RequestId   string `json:"request_id"` // used for troubleshooting
	LastCursor  string `json:"last_cursor"`
}
