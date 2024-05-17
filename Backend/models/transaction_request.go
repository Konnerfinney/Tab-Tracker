package models

import "time"

// TransactionRequest represents a request for a transaction between two users.
type TransactionRequest struct {
	RequestID     int       `json:"request_id"`
	TransactionID int       `json:"transaction_id"`
	RequesterID   int       `json:"requester_id"`
	RequesteeID   int       `json:"requestee_id"`
	RequestDate   time.Time `json:"request_date"`
	Status        string    `json:"status"`
	ResponseDate  time.Time `json:"response_date,omitempty"`
}
