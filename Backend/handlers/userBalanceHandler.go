package handlers

import (
	"Tab_Tracker/db"
	"Tab_Tracker/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetUserBalancesHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, err := strconv.Atoi(vars["user_id"])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	dbConn := db.ConnectDB()
	defer dbConn.Close()

	balances, err := db.GetUserBalances(userID, dbConn)
	if err != nil {
		http.Error(w, "Failed to retrieve balances", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(balances)
}

func CreateUserBalanceHandler(w http.ResponseWriter, r *http.Request) {
	var balance models.UserBalance
	err := json.NewDecoder(r.Body).Decode(&balance)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	dbConn := db.ConnectDB()
	defer dbConn.Close()

	balanceID, err := db.InsertUserBalance(balance, dbConn)
	if err != nil {
		http.Error(w, "Failed to create balance", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]int{"balance_id": balanceID})
}
