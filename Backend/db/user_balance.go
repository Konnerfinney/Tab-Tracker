package db

import (
	"Tab_Tracker/models"
	"database/sql"
	"fmt"
)

// InsertUserBalance inserts a new user balance into the database.
func InsertUserBalance(balance models.UserBalance, db *sql.DB) (int, error) {
	var balanceID int
	err := db.QueryRow(
		"INSERT INTO user_balances (user1_id, user2_id, net_balance) VALUES ($1, $2, $3) RETURNING balance_id",
		balance.User1ID, balance.User2ID, balance.NetBalance,
	).Scan(&balanceID)
	if err != nil {
		return 0, fmt.Errorf("unable to insert user balance: %v", err)
	}
	return balanceID, nil
}

// GetUserBalance retrieves the balance between two users.
func GetUserBalance(user1ID, user2ID int, db *sql.DB) (models.UserBalance, error) {
	var balance models.UserBalance
	err := db.QueryRow(
		"SELECT balance_id, user1_id, user2_id, net_balance FROM user_balances WHERE user1_id = $1 AND user2_id = $2",
		user1ID, user2ID,
	).Scan(&balance.BalanceID, &balance.User1ID, &balance.User2ID, &balance.NetBalance)
	if err != nil {
		return balance, fmt.Errorf("unable to get user balance: %v", err)
	}
	return balance, nil
}

// GetUserBalances retrieves all balances for a specific user.
func GetUserBalances(userID int, db *sql.DB) ([]models.UserBalance, error) {
	rows, err := db.Query(
		"SELECT balance_id, user1_id, user2_id, net_balance FROM user_balances WHERE user1_id = $1 OR user2_id = $1",
		userID,
	)
	if err != nil {
		return nil, fmt.Errorf("unable to query user balances: %v", err)
	}
	defer rows.Close()

	var balances []models.UserBalance
	for rows.Next() {
		var balance models.UserBalance
		if err := rows.Scan(&balance.BalanceID, &balance.User1ID, &balance.User2ID, &balance.NetBalance); err != nil {
			return nil, fmt.Errorf("unable to scan user balance: %v", err)
		}
		balances = append(balances, balance)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating through user balances: %v", err)
	}

	return balances, nil
}
