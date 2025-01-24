package models

import "time"

// Transaction represents a transaction entity
type Transaction struct {
	Id          string
	Description string
	Amount      float64
	Category    string
	Date        time.Time
}

// ToRow converts a Transaction to a row that can be appended to Google Sheets
func (t Transaction) ToRow() []interface{} {
	return []interface{}{t.Id, t.Description, t.Amount, t.Category, t.Date.Format(time.DateOnly)}
}
