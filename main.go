package main

import (
    "fmt"
    "net/http"
    "github.com/gorilla/mux"
	//"Tab_Tracker/db"
)

func main() {
	// database := db.ConnectDB()
    // defer database.Close()
	// db.CreateTable(database)
	// db.InsertTestRecords(database)
    r := mux.NewRouter()
    // Define your routes here
    r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hello, World!"))
    })
    r.HandleFunc("/users/{userID}", UserHandler).Methods("GET")

    http.ListenAndServe(":8080", r)
}

func UserHandler(w http.ResponseWriter, r *http.Request) {
    // Extract variables from the route using mux.Vars
    vars := mux.Vars(r)
    userID := vars["userID"]

    // Write the userID to the response
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "User ID: %s", userID)
}