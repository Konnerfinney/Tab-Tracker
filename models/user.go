package models

type User struct {
	UserID     int    `json:"user_id"`
	ProviderID string `json:"provider_id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Email      string `json:"email"`
	Provider   string `json:"provider"`
}
