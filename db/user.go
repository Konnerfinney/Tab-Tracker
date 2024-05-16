package db

import (
	"Tab_Tracker/models"
	"database/sql"
	"fmt"
)

// HandleUser inserts a new user or updates an existing user in the database.
func HandleUser(user models.User, db *sql.DB) error {
	var existingUser models.User
	err := db.QueryRow(
		"SELECT id, provider_id, first_name, last_name, email, provider FROM users WHERE provider_id=$1",
		user.ProviderID,
	).Scan(&existingUser.UserID, &existingUser.ProviderID, &existingUser.FirstName, &existingUser.LastName, &existingUser.Email, &existingUser.Provider)

	if err != nil {
		if err == sql.ErrNoRows {
			// User does not exist, insert new user
			_, err = db.Exec(
				"INSERT INTO users (provider_id, first_name, last_name, email, provider) VALUES ($1, $2, $3, $4, $5)",
				user.ProviderID, user.FirstName, user.LastName, user.Email, user.Provider,
			)
			if err != nil {
				return fmt.Errorf("unable to insert new user: %v", err)
			}
		} else {
			return fmt.Errorf("unable to query user: %v", err)
		}
	}

	// User exists or newly created
	return nil
}
