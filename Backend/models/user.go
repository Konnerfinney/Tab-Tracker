package models

import "time"

// User represents a user in the system.
type User struct {
	UserID      int       `json:"user_id"`
	ProviderID  string    `json:"provider_id"`
	Username    string    `json:"username"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Email       string    `json:"email"`
	Provider    string    `json:"provider"`
	IsProxyUser bool      `json:"is_proxy_user"`
	CreatedAt   time.Time `json:"created_at"`
}
