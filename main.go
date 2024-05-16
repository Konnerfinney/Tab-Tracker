package main

import (
	"Tab_Tracker/auth"
	"Tab_Tracker/db"
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
)

func main() {
	// Initialize the database connection
	database := db.ConnectDB()
	defer database.Close()

	// Ensure the database has the required tables
	db.CreateTable(database)
	// Optionally insert test records
	// db.InsertTestRecords(database)

	// Initialize the authentication system
	auth.SetupAuth()

	// Create a new router
	r := mux.NewRouter()

	// Define your routes here
	r.HandleFunc("/", HomeHandler).Methods("GET")
	r.HandleFunc("/users/{userID}", UserHandler).Methods("GET")

	// Authentication routes
	r.HandleFunc("/auth/{provider}/callback", auth.AuthenticateHandler).Methods("GET")
	r.HandleFunc("/auth/{provider}", auth.AuthenticateHandler).Methods("GET")

	// Start the server
	log.Println("listening on localhost:3000")
	http.ListenAndServe(":3000", r)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.html")
	if err != nil {
		log.Printf("Error parsing template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	t.Execute(w, nil)
}

func UserHandler(w http.ResponseWriter, r *http.Request) {
	// Extract variables from the route using mux.Vars
	vars := mux.Vars(r)
	userID := vars["userID"]

	// Write the userID to the response
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "User ID: %s", userID)
}
