package db

import (
	"Tab_Tracker/models"
	"database/sql"
	"fmt"
)

// HandleUser inserts a new user or updates an existing user in the database.
func HandleUser(user models.User, db *sql.DB) error {
	// Check if the user already exists
	existingUser, err := GetUserByProviderID(user.ProviderID, db)
	if err != nil {
		if err == sql.ErrNoRows {
			// User does not exist, insert new user
			_, err = InsertUser(user, db)
			if err != nil {
				return fmt.Errorf("unable to insert new user: %v", err)
			}
		} else {
			return fmt.Errorf("unable to query user: %v", err)
		}
	} else {
		// User exists, you can update user details if needed here
		fmt.Printf("User already exists: %+v\n", existingUser)
	}

	return nil
}

// InsertUser inserts a new user into the database and returns the user ID.
func InsertUser(user models.User, db *sql.DB) (int, error) {
	var userID int
	err := db.QueryRow(
		"INSERT INTO users (provider_id, username, first_name, last_name, email, provider, is_proxy_user) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING user_id",
		user.ProviderID, user.Username, user.FirstName, user.LastName, user.Email, user.Provider, user.IsProxyUser,
	).Scan(&userID)
	if err != nil {
		return 0, fmt.Errorf("unable to insert user: %v", err)
	}
	return userID, nil
}

// GetUser retrieves a user by their userID.
func GetUser(userID int, db *sql.DB) (models.User, error) {
	var user models.User
	err := db.QueryRow(
		"SELECT user_id, provider_id, username, first_name, last_name, email, provider, is_proxy_user, created_at FROM users WHERE user_id = $1",
		userID,
	).Scan(&user.UserID, &user.ProviderID, &user.Username, &user.FirstName, &user.LastName, &user.Email, &user.Provider, &user.IsProxyUser, &user.CreatedAt)
	if err != nil {
		return user, fmt.Errorf("unable to get user: %v", err)
	}
	return user, nil
}

// GetUserByProviderID retrieves a user by their providerID.
func GetUserByProviderID(providerID string, db *sql.DB) (models.User, error) {
	var user models.User
	err := db.QueryRow(
		"SELECT user_id, provider_id, username, first_name, last_name, email, provider, is_proxy_user, created_at FROM users WHERE provider_id = $1",
		providerID,
	).Scan(&user.UserID, &user.ProviderID, &user.Username, &user.FirstName, &user.LastName, &user.Email, &user.Provider, &user.IsProxyUser, &user.CreatedAt)
	if err != nil {
		return user, fmt.Errorf("unable to get user: %v", err)
	}
	return user, nil
}
