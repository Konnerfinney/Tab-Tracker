package db

import (
	"Tab_Tracker/models"
	"database/sql"
	"fmt"
)

// InsertTransactionRequest inserts a new transaction request into the database.
func InsertTransactionRequest(request models.TransactionRequest, db *sql.DB) (int, error) {
	var requestID int
	err := db.QueryRow(
		"INSERT INTO transaction_requests (transaction_id, requester_id, requestee_id, status) VALUES ($1, $2, $3, $4) RETURNING request_id",
		request.TransactionID, request.RequesterID, request.RequesteeID, request.Status,
	).Scan(&requestID)
	if err != nil {
		return 0, fmt.Errorf("unable to insert transaction request: %v", err)
	}
	return requestID, nil
}

// GetTransactionRequest retrieves a transaction request by its ID.
func GetTransactionRequest(requestID int, db *sql.DB) (models.TransactionRequest, error) {
	var request models.TransactionRequest
	err := db.QueryRow(
		"SELECT request_id, transaction_id, requester_id, requestee_id, request_date, status, response_date FROM transaction_requests WHERE request_id = $1",
		requestID,
	).Scan(&request.RequestID, &request.TransactionID, &request.RequesterID, &request.RequesteeID, &request.RequestDate, &request.Status, &request.ResponseDate)
	if err != nil {
		return request, fmt.Errorf("unable to get transaction request: %v", err)
	}
	return request, nil
}
