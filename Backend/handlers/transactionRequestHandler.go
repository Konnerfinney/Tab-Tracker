package handlers

import (
	"Tab_Tracker/db"
	"Tab_Tracker/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetTransactionRequestHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	requestID, err := strconv.Atoi(vars["request_id"])
	if err != nil {
		http.Error(w, "Invalid request ID", http.StatusBadRequest)
		return
	}

	dbConn := db.ConnectDB()
	defer dbConn.Close()

	request, err := db.GetTransactionRequest(requestID, dbConn)
	if err != nil {
		http.Error(w, "Request not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(request)
}

func CreateTransactionRequestHandler(w http.ResponseWriter, r *http.Request) {
	var request models.TransactionRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	dbConn := db.ConnectDB()
	defer dbConn.Close()

	requestID, err := db.InsertTransactionRequest(request, dbConn)
	if err != nil {
		http.Error(w, "Failed to create request", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]int{"request_id": requestID})
}
