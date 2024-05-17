package handlers

import (
	"Tab_Tracker/db"
	"Tab_Tracker/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetUserTransactionHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	transactionID, err := strconv.Atoi(vars["transaction_id"])
	if err != nil {
		http.Error(w, "Invalid transaction ID", http.StatusBadRequest)
		return
	}

	dbConn := db.ConnectDB()
	defer dbConn.Close()

	transaction, err := db.GetUserTransaction(transactionID, dbConn)
	if err != nil {
		http.Error(w, "Transaction not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(transaction)
}

func CreateUserTransactionHandler(w http.ResponseWriter, r *http.Request) {
	var transaction models.UserTransaction
	err := json.NewDecoder(r.Body).Decode(&transaction)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	dbConn := db.ConnectDB()
	defer dbConn.Close()

	transactionID, err := db.InsertUserTransaction(transaction, dbConn)
	if err != nil {
		http.Error(w, "Failed to create transaction", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]int{"transaction_id": transactionID})
}
