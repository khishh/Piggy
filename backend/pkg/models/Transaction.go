package models

type Transaction struct {
	ID             string   `gorm:"primary" json:"id"`
	AccountId      string   `json:"account_id"`
	Amount         float64  `gorm:"type:real" json:"amount"`
	AuthorizedDate string   `json:"authorized_date"`
	Category       []string `gorm:"type:text[]" json:"category"`
	Date           string   `gorm:"type:date" json:"date"`
	CurrencyCode   string   `json:"iso_currency_code"`
	Location       Location `gorm:"embedded"`
}
