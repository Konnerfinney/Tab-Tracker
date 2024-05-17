package handlers

import (
	"Tab_Tracker/db"
	"Tab_Tracker/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetInteractionHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	interactionID, err := strconv.Atoi(vars["interaction_id"])
	if err != nil {
		http.Error(w, "Invalid interaction ID", http.StatusBadRequest)
		return
	}

	dbConn := db.ConnectDB()
	defer dbConn.Close()

	interaction, err := db.GetInteraction(interactionID, dbConn)
	if err != nil {
		http.Error(w, "Interaction not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(interaction)
}

func GetUserInteractionsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, err := strconv.Atoi(vars["user_id"])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	dbConn := db.ConnectDB()
	defer dbConn.Close()

	interactions, err := db.GetInteractionsByUser(userID, dbConn)
	if err != nil {
		http.Error(w, "Failed to retrieve interactions", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(interactions)
}

func CreateInteractionHandler(w http.ResponseWriter, r *http.Request) {
	var interaction models.Interaction
	err := json.NewDecoder(r.Body).Decode(&interaction)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	dbConn := db.ConnectDB()
	defer dbConn.Close()

	interactionID, err := db.InsertInteraction(interaction, dbConn)
	if err != nil {
		http.Error(w, "Failed to create interaction", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]int{"interaction_id": interactionID})
}
