package tests

import (
	"Tab_Tracker/db"
	"Tab_Tracker/handlers"
	"Tab_Tracker/models"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func setupRouter(t *testing.T) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/users", handlers.CreateUserHandler).Methods("POST")
	r.HandleFunc("/users/{user_id}", handlers.GetUserHandler).Methods("GET")
	r.HandleFunc("/balances/{user_id}", handlers.GetUserBalancesHandler).Methods("GET")
	r.HandleFunc("/transactions", handlers.CreateUserTransactionHandler).Methods("POST")
	r.HandleFunc("/transactions/{transaction_id}", handlers.GetUserTransactionHandler).Methods("GET")
	r.HandleFunc("/requests", handlers.CreateTransactionRequestHandler).Methods("POST")
	r.HandleFunc("/requests/{request_id}", handlers.GetTransactionRequestHandler).Methods("GET")
	r.HandleFunc("/interactions", handlers.CreateInteractionHandler).Methods("POST")
	r.HandleFunc("/interactions/{interaction_id}", handlers.GetInteractionHandler).Methods("GET")
	r.HandleFunc("/interactions/user/{user_id}", handlers.GetUserInteractionsHandler).Methods("GET")
	return r
}

func TestCreateUserHandler(t *testing.T) {
	defer resetEnv(os.Environ())()
	testDB := SetupTestDB(t)
	defer testDB.Close()
	r := setupRouter(t)
	user := models.User{
		ProviderID:  "google-12345",
		Username:    "john_doe",
		FirstName:   "John",
		LastName:    "Doe",
		Email:       "john.doe@example.com",
		Provider:    "google",
		IsProxyUser: false,
	}
	userData, _ := json.Marshal(user)
	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(userData))
	rr := httptest.NewRecorder()

	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)
}

func TestGetUserHandler(t *testing.T) {
	defer resetEnv(os.Environ())()
	testDB := SetupTestDB(t)
	defer testDB.Close()

	r := setupRouter(t)

	user := models.User{
		ProviderID:  "google-12345",
		Username:    "john_doe",
		FirstName:   "John",
		LastName:    "Doe",
		Email:       "john.doe@example.com",
		Provider:    "google",
		IsProxyUser: false,
	}

	userID, err := db.InsertUser(user, testDB)
	assert.NoError(t, err)

	req, _ := http.NewRequest("GET", "/users/"+strconv.Itoa(userID), nil)
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	var fetchedUser models.User
	err = json.NewDecoder(rr.Body).Decode(&fetchedUser)
	assert.NoError(t, err)
	assert.Equal(t, user.Email, fetchedUser.Email)
}

func TestCreateOrUpdateUserBalanceHandler(t *testing.T) {
	defer resetEnv(os.Environ())()
	testDB := SetupTestDB(t)
	defer testDB.Close()

	r := setupRouter(t)

	user1 := models.User{
		ProviderID:  "google-12345",
		Username:    "john_doe",
		FirstName:   "John",
		LastName:    "Doe",
		Email:       "john.doe@example.com",
		Provider:    "google",
		IsProxyUser: false,
	}
	user2 := models.User{
		ProviderID:  "google-67890",
		Username:    "jane_doe",
		FirstName:   "Jane",
		LastName:    "Doe",
		Email:       "jane.doe@example.com",
		Provider:    "google",
		IsProxyUser: false,
	}

	userID1, err := db.InsertUser(user1, testDB)
	assert.NoError(t, err)
	userID2, err := db.InsertUser(user2, testDB)
	assert.NoError(t, err)

	balance := models.UserBalance{
		User1ID:    userID1,
		User2ID:    userID2,
		NetBalance: 100.00,
	}

	balanceData, _ := json.Marshal(balance)
	req, _ := http.NewRequest("POST", "/balances", bytes.NewBuffer(balanceData))
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)
}

func TestGetUserBalancesHandler(t *testing.T) {
	defer resetEnv(os.Environ())()
	testDB := SetupTestDB(t)
	defer testDB.Close()

	r := setupRouter(t)

	user1 := models.User{
		ProviderID:  "google-12345",
		Username:    "john_doe",
		FirstName:   "John",
		LastName:    "Doe",
		Email:       "john.doe@example.com",
		Provider:    "google",
		IsProxyUser: false,
	}
	user2 := models.User{
		ProviderID:  "google-67890",
		Username:    "jane_doe",
		FirstName:   "Jane",
		LastName:    "Doe",
		Email:       "jane.doe@example.com",
		Provider:    "google",
		IsProxyUser: false,
	}

	userID1, err := db.InsertUser(user1, testDB)
	assert.NoError(t, err)
	userID2, err := db.InsertUser(user2, testDB)
	assert.NoError(t, err)

	balance := models.UserBalance{
		User1ID:    userID1,
		User2ID:    userID2,
		NetBalance: 100.00,
	}

	_, err = db.InsertUserBalance(balance, testDB)
	assert.NoError(t, err)

	req, _ := http.NewRequest("GET", "/balances/"+strconv.Itoa(userID1), nil)
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	var fetchedBalances []models.UserBalance
	err = json.NewDecoder(rr.Body).Decode(&fetchedBalances)
	assert.NoError(t, err)
	assert.Len(t, fetchedBalances, 1)
}

func TestCreateTransactionHandler(t *testing.T) {
	defer resetEnv(os.Environ())()
	testDB := SetupTestDB(t)
	defer testDB.Close()

	r := setupRouter(t)

	user1 := models.User{
		ProviderID:  "google-12345",
		Username:    "john_doe",
		FirstName:   "John",
		LastName:    "Doe",
		Email:       "john.doe@example.com",
		Provider:    "google",
		IsProxyUser: false,
	}
	user2 := models.User{
		ProviderID:  "google-67890",
		Username:    "jane_doe",
		FirstName:   "Jane",
		LastName:    "Doe",
		Email:       "jane.doe@example.com",
		Provider:    "google",
		IsProxyUser: false,
	}

	userID1, err := db.InsertUser(user1, testDB)
	assert.NoError(t, err)
	userID2, err := db.InsertUser(user2, testDB)
	assert.NoError(t, err)

	transaction := models.UserTransaction{
		CreditorID:  userID1,
		DebtorID:    userID2,
		Amount:      50.00,
		Description: "Dinner payment",
	}

	transactionData, _ := json.Marshal(transaction)
	req, _ := http.NewRequest("POST", "/transactions", bytes.NewBuffer(transactionData))
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)
}

func TestGetTransactionHandler(t *testing.T) {
	defer resetEnv(os.Environ())()
	testDB := SetupTestDB(t)
	defer testDB.Close()

	r := setupRouter(t)

	user1 := models.User{
		ProviderID:  "google-12345",
		Username:    "john_doe",
		FirstName:   "John",
		LastName:    "Doe",
		Email:       "john.doe@example.com",
		Provider:    "google",
		IsProxyUser: false,
	}
	user2 := models.User{
		ProviderID:  "google-67890",
		Username:    "jane_doe",
		FirstName:   "Jane",
		LastName:    "Doe",
		Email:       "jane.doe@example.com",
		Provider:    "google",
		IsProxyUser: false,
	}

	userID1, err := db.InsertUser(user1, testDB)
	assert.NoError(t, err)
	userID2, err := db.InsertUser(user2, testDB)
	assert.NoError(t, err)

	transaction := models.UserTransaction{
		CreditorID:  userID1,
		DebtorID:    userID2,
		Amount:      50.00,
		Description: "Dinner payment",
	}

	transactionID, err := db.InsertUserTransaction(transaction, testDB)
	assert.NoError(t, err)

	req, _ := http.NewRequest("GET", "/transactions/"+strconv.Itoa(transactionID), nil)
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	var fetchedTransaction models.UserTransaction
	err = json.NewDecoder(rr.Body).Decode(&fetchedTransaction)
	assert.NoError(t, err)
	assert.Equal(t, transaction.Description, fetchedTransaction.Description)
}

func TestCreateTransactionRequestHandler(t *testing.T) {
	defer resetEnv(os.Environ())()
	testDB := SetupTestDB(t)
	defer testDB.Close()

	r := setupRouter(t)

	user1 := models.User{
		ProviderID:  "google-12345",
		Username:    "john_doe",
		FirstName:   "John",
		LastName:    "Doe",
		Email:       "john.doe@example.com",
		Provider:    "google",
		IsProxyUser: false,
	}
	user2 := models.User{
		ProviderID:  "google-67890",
		Username:    "jane_doe",
		FirstName:   "Jane",
		LastName:    "Doe",
		Email:       "jane.doe@example.com",
		Provider:    "google",
		IsProxyUser: false,
	}

	userID1, err := db.InsertUser(user1, testDB)
	assert.NoError(t, err)
	userID2, err := db.InsertUser(user2, testDB)
	assert.NoError(t, err)

	transaction := models.UserTransaction{
		CreditorID:  userID1,
		DebtorID:    userID2,
		Amount:      50.00,
		Description: "Dinner payment",
	}

	transactionID, err := db.InsertUserTransaction(transaction, testDB)
	assert.NoError(t, err)

	request := models.TransactionRequest{
		TransactionID: transactionID,
		RequesterID:   userID1,
		RequesteeID:   userID2,
		Status:        "pending",
	}

	requestData, _ := json.Marshal(request)
	req, _ := http.NewRequest("POST", "/requests", bytes.NewBuffer(requestData))
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)
}

func TestGetTransactionRequestHandler(t *testing.T) {
	defer resetEnv(os.Environ())()
	testDB := SetupTestDB(t)
	defer testDB.Close()

	r := setupRouter(t)

	user1 := models.User{
		ProviderID:  "google-12345",
		Username:    "john_doe",
		FirstName:   "John",
		LastName:    "Doe",
		Email:       "john.doe@example.com",
		Provider:    "google",
		IsProxyUser: false,
	}
	user2 := models.User{
		ProviderID:  "google-67890",
		Username:    "jane_doe",
		FirstName:   "Jane",
		LastName:    "Doe",
		Email:       "jane.doe@example.com",
		Provider:    "google",
		IsProxyUser: false,
	}

	userID1, err := db.InsertUser(user1, testDB)
	assert.NoError(t, err)
	userID2, err := db.InsertUser(user2, testDB)
	assert.NoError(t, err)

	transaction := models.UserTransaction{
		CreditorID:  userID1,
		DebtorID:    userID2,
		Amount:      50.00,
		Description: "Dinner payment",
	}

	transactionID, err := db.InsertUserTransaction(transaction, testDB)
	assert.NoError(t, err)

	request := models.TransactionRequest{
		TransactionID: transactionID,
		RequesterID:   userID1,
		RequesteeID:   userID2,
		Status:        "pending",
	}

	requestID, err := db.InsertTransactionRequest(request, testDB)
	assert.NoError(t, err)

	req, _ := http.NewRequest("GET", "/requests/"+strconv.Itoa(requestID), nil)
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	var fetchedRequest models.TransactionRequest
	err = json.NewDecoder(rr.Body).Decode(&fetchedRequest)
	assert.NoError(t, err)
	assert.Equal(t, request.Status, fetchedRequest.Status)
}

func TestCreateInteractionHandler(t *testing.T) {
	defer resetEnv(os.Environ())()
	testDB := SetupTestDB(t)
	defer testDB.Close()

	r := setupRouter(t)

	timestamp, err := time.Parse(time.RFC3339, "2023-05-16T10:00:00Z")
	assert.NoError(t, err)

	interaction := models.Interaction{
		InteractionStarter:       1,
		InteractionReceiver:      2,
		InteractionCount:         1,
		LastInteractionTimestamp: timestamp,
	}

	interactionData, _ := json.Marshal(interaction)
	req, _ := http.NewRequest("POST", "/interactions", bytes.NewBuffer(interactionData))
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)
}

func TestGetInteractionHandler(t *testing.T) {
	defer resetEnv(os.Environ())()
	testDB := SetupTestDB(t)
	defer testDB.Close()

	r := setupRouter(t)

	timestamp, err := time.Parse(time.RFC3339, "2023-05-16T10:00:00Z")
	assert.NoError(t, err)

	interaction := models.Interaction{
		InteractionStarter:       1,
		InteractionReceiver:      2,
		InteractionCount:         1,
		LastInteractionTimestamp: timestamp,
	}

	interactionID, err := db.InsertInteraction(interaction, testDB)
	assert.NoError(t, err)

	req, _ := http.NewRequest("GET", "/interactions/"+strconv.Itoa(interactionID), nil)
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	var fetchedInteraction models.Interaction
	err = json.NewDecoder(rr.Body).Decode(&fetchedInteraction)
	assert.NoError(t, err)
	assert.Equal(t, interaction.InteractionCount, fetchedInteraction.InteractionCount)
}

func TestGetUserInteractionsHandler(t *testing.T) {
	defer resetEnv(os.Environ())()
	testDB := SetupTestDB(t)
	defer testDB.Close()

	r := setupRouter(t)

	timestamp, err := time.Parse(time.RFC3339, "2023-05-16T10:00:00Z")
	assert.NoError(t, err)

	interaction1 := models.Interaction{
		InteractionStarter:       1,
		InteractionReceiver:      2,
		InteractionCount:         1,
		LastInteractionTimestamp: timestamp,
	}
	interaction2 := models.Interaction{
		InteractionStarter:       1,
		InteractionReceiver:      3,
		InteractionCount:         1,
		LastInteractionTimestamp: timestamp,
	}

	_, err = db.InsertInteraction(interaction1, testDB)
	assert.NoError(t, err)
	_, err = db.InsertInteraction(interaction2, testDB)
	assert.NoError(t, err)

	req, _ := http.NewRequest("GET", "/interactions/user/1", nil)
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	var fetchedInteractions []models.Interaction
	err = json.NewDecoder(rr.Body).Decode(&fetchedInteractions)
	assert.NoError(t, err)
	assert.Len(t, fetchedInteractions, 2)
}
