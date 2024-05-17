package auth

import (
	"Tab_Tracker/db"
	"Tab_Tracker/models"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
)

// SetupAuth initializes the authentication system
func SetupAuth() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	key := "Secret-session-key" // Replace with your SESSION_SECRET or similar
	maxAge := 86400 * 30        // 30 days
	isProd := false             // Set to true when serving over https

	store := sessions.NewCookieStore([]byte(key))
	store.MaxAge(maxAge)
	store.Options.Path = "/"
	store.Options.HttpOnly = true // HttpOnly should always be enabled
	store.Options.Secure = isProd

	gothic.Store = store

	clientID := os.Getenv("GOOGLE_CLIENT_ID")
	clientSecret := os.Getenv("GOOGLE_CLIENT_SECRET")
	redirectURL := "http://localhost:3000/auth/google/callback"

	if clientID == "" || clientSecret == "" {
		log.Fatal("Google client ID or secret not set in environment")
	}

	goth.UseProviders(
		google.New(
			clientID,
			clientSecret,
			redirectURL,
			"email", "profile",
		),
	)
}

// AuthenticateHandler handles the authentication routes
func AuthenticateHandler(w http.ResponseWriter, r *http.Request) {
	// Extract the provider from the URL path
	provider := mux.Vars(r)["provider"]
	if provider == "" {
		http.Error(w, "Provider not specified", http.StatusBadRequest)
		return
	}

	// If the URL path contains "callback", handle the callback
	if r.URL.Path == fmt.Sprintf("/auth/%s/callback", provider) {
		googleCallback(w, r)
		return
	}

	// Otherwise, begin the auth process
	gothic.BeginAuthHandler(w, r)
}

func googleCallback(res http.ResponseWriter, req *http.Request) {
	user, err := gothic.CompleteUserAuth(res, req)
	if err != nil {
		log.Printf("Error completing user auth: %v", err)
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	newUser := models.User{
		ProviderID: user.UserID,
		FirstName:  user.FirstName,
		LastName:   user.LastName,
		Email:      user.Email,
		Provider:   user.Provider,
	}

	dbConn := db.ConnectDB()
	defer dbConn.Close() // Close the database connection when done

	err = db.HandleUser(newUser, dbConn)
	if err != nil {
		log.Printf("Error handling user: %v", err)
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	t, err := template.ParseFiles("templates/success.html")
	if err != nil {
		log.Printf("Error parsing template: %v", err)
		http.Error(res, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	t.Execute(res, user)
}
