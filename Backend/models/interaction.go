package models

import "time"

// Interaction represents an interaction between two users.
type Interaction struct {
	InteractionID            int       `json:"interaction_id"`
	InteractionStarter       int       `json:"interaction_starter"`
	InteractionReceiver      int       `json:"interaction_receiver"`
	InteractionCount         int       `json:"interaction_count"`
	LastInteractionTimestamp time.Time `json:"last_interaction_timestamp"`
}
