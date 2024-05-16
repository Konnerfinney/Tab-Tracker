package models

import "time"

// UserTransaction represents a transaction between two users.
type UserTransaction struct {
	TransactionID int       `json:"transaction_id"`
	CreditorID    int       `json:"creditor_id"`
	DebtorID      int       `json:"debtor_id"`
	Amount        float64   `json:"amount"`
	Description   string    `json:"description"`
	CreatedAt     time.Time `json:"created_at"`
}
