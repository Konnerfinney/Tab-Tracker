package db

import (
	"Tab_Tracker/models"
	"database/sql"
	"fmt"
)

// InsertUserTransaction inserts a new transaction into the database.
func InsertUserTransaction(transaction models.UserTransaction, db *sql.DB) (int, error) {
	var transactionID int
	err := db.QueryRow(
		"INSERT INTO user_transactions (creditor_id, debtor_id, amount, description) VALUES ($1, $2, $3, $4) RETURNING transaction_id",
		transaction.CreditorID, transaction.DebtorID, transaction.Amount, transaction.Description,
	).Scan(&transactionID)
	if err != nil {
		return 0, fmt.Errorf("unable to insert user transaction: %v", err)
	}
	return transactionID, nil
}

// GetUserTransaction retrieves a transaction by its ID.
func GetUserTransaction(transactionID int, db *sql.DB) (models.UserTransaction, error) {
	var transaction models.UserTransaction
	err := db.QueryRow(
		"SELECT transaction_id, creditor_id, debtor_id, amount, description, created_at FROM user_transactions WHERE transaction_id = $1",
		transactionID,
	).Scan(&transaction.TransactionID, &transaction.CreditorID, &transaction.DebtorID, &transaction.Amount, &transaction.Description, &transaction.CreatedAt)
	if err != nil {
		return transaction, fmt.Errorf("unable to get user transaction: %v", err)
	}
	return transaction, nil
}
