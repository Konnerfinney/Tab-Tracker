package db

import (
	"Tab_Tracker/models"
	"database/sql"
	"fmt"
)

// InsertInteraction inserts a new interaction into the database and returns the interaction ID.
func InsertInteraction(interaction models.Interaction, db *sql.DB) (int, error) {
	var interactionID int
	err := db.QueryRow(
		"INSERT INTO interactions (interaction_starter, interaction_receiver, interaction_count, last_interaction_timestamp) VALUES ($1, $2, $3, $4) RETURNING interaction_id",
		interaction.InteractionStarter, interaction.InteractionReceiver, interaction.InteractionCount, interaction.LastInteractionTimestamp,
	).Scan(&interactionID)
	if err != nil {
		return 0, fmt.Errorf("unable to insert interaction: %v", err)
	}
	return interactionID, nil
}

// GetInteraction retrieves an interaction by its ID.
func GetInteraction(interactionID int, db *sql.DB) (models.Interaction, error) {
	var interaction models.Interaction
	err := db.QueryRow(
		"SELECT interaction_id, interaction_starter, interaction_receiver, interaction_count, last_interaction_timestamp FROM interactions WHERE interaction_id = $1",
		interactionID,
	).Scan(&interaction.InteractionID, &interaction.InteractionStarter, &interaction.InteractionReceiver, &interaction.InteractionCount, &interaction.LastInteractionTimestamp)
	if err != nil {
		return interaction, fmt.Errorf("unable to get interaction: %v", err)
	}
	return interaction, nil
}

// GetInteractionsByUser retrieves interactions involving a specific user.
func GetInteractionsByUser(userID int, db *sql.DB) ([]models.Interaction, error) {
	rows, err := db.Query(
		"SELECT interaction_id, interaction_starter, interaction_receiver, interaction_count, last_interaction_timestamp FROM interactions WHERE interaction_starter = $1 OR interaction_receiver = $1",
		userID,
	)
	if err != nil {
		return nil, fmt.Errorf("unable to query interactions: %v", err)
	}
	defer rows.Close()

	var interactions []models.Interaction
	for rows.Next() {
		var interaction models.Interaction
		if err := rows.Scan(&interaction.InteractionID, &interaction.InteractionStarter, &interaction.InteractionReceiver, &interaction.InteractionCount, &interaction.LastInteractionTimestamp); err != nil {
			return nil, fmt.Errorf("unable to scan interaction: %v", err)
		}
		interactions = append(interactions, interaction)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating through interactions: %v", err)
	}

	return interactions, nil
}
