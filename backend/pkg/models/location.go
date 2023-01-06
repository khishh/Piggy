package models

type Location struct {
	Address    string  `json:"address"`
	City       string  `json:"city"`
	Country    string  `json:"country"`
	Latitude   float64 `gorm:"type:real" json:"lat"`
	Longitude  float64 `gorm:"type:real" json:"lon"`
	PostalCode string  `json:"postal_code"`
	Region     string  `json:"region"`
}
