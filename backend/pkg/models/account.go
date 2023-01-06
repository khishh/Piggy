package models

type Account struct {
	ID           string `gorm:"primary" json:"id"`
	ItemId       string `json:"item_id"`
	Name         string `json:"name"`
	OfficialName string `json:"official_name"`
	SubType      string `json:"sub_type"`
	Type         string `json:"type"`
}
