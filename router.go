package main

import (
	"Tab_Tracker/auth"
	"Tab_Tracker/handlers"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/", HomeHandler).Methods("GET")
	r.HandleFunc("/users/{userID}", UserHandler).Methods("GET")

	// Authentication routes
	r.HandleFunc("/auth/{provider}/callback", auth.AuthenticateHandler).Methods("GET")
	r.HandleFunc("/auth/{provider}", auth.AuthenticateHandler).Methods("GET")

	// User routes
	r.HandleFunc("/users/{user_id}", handlers.GetUserHandler).Methods("GET")
	r.HandleFunc("/users", handlers.CreateUserHandler).Methods("POST")

	// User balance routes
	r.HandleFunc("/balances/{user_id}", handlers.GetUserBalancesHandler).Methods("GET")
	r.HandleFunc("/balances", handlers.CreateUserBalanceHandler).Methods("POST")

	// User transaction routes
	r.HandleFunc("/transactions/{transaction_id}", handlers.GetUserTransactionHandler).Methods("GET")
	r.HandleFunc("/transactions", handlers.CreateUserTransactionHandler).Methods("POST")

	// Transaction request routes
	r.HandleFunc("/requests/{request_id}", handlers.GetTransactionRequestHandler).Methods("GET")
	r.HandleFunc("/requests", handlers.CreateTransactionRequestHandler).Methods("POST")

	// Interaction routes
	r.HandleFunc("/interactions/{interaction_id}", handlers.GetInteractionHandler).Methods("GET")
	r.HandleFunc("/interactions/user/{user_id}", handlers.GetUserInteractionsHandler).Methods("GET")
	r.HandleFunc("/interactions", handlers.CreateInteractionHandler).Methods("POST")

	return r
}
