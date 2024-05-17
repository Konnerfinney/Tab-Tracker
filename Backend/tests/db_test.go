package tests

import (
	"Tab_Tracker/db"
	"Tab_Tracker/models"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func TestInsertUser(t *testing.T) {
	testDB := SetupTestDB(t)
	defer testDB.Close()

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
	assert.NotZero(t, userID)
}

func TestGetUser(t *testing.T) {
	testDB := SetupTestDB(t)
	defer testDB.Close()

	user := models.User{
		ProviderID:  "google-12345",
		Username:    "john_doe",
		FirstName:   "John",
		LastName:    "Doe",
		Email:       "john.doe@example.com",
		Provider:    "google",
		IsProxyUser: false,
	}
	log.Printf("Test!")
	userID, err := db.InsertUser(user, testDB)
	assert.NoError(t, err)

	fetchedUser, err := db.GetUser(userID, testDB)
	assert.NoError(t, err)
	assert.Equal(t, user.Email, fetchedUser.Email)

}

func TestInsertUserBalance(t *testing.T) {
	defer resetEnv(os.Environ())()
	testDB := SetupTestDB(t)
	defer testDB.Close()

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

	balanceID, err := db.InsertUserBalance(balance, testDB)
	assert.NoError(t, err)
	assert.NotZero(t, balanceID)
}

func TestGetUserBalance(t *testing.T) {
	defer resetEnv(os.Environ())()
	testDB := SetupTestDB(t)
	defer testDB.Close()

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

	balanceID, err := db.InsertUserBalance(balance, testDB)
	assert.NoError(t, err)
	assert.NotZero(t, balanceID)

	fetchedBalance, err := db.GetUserBalance(userID1, userID2, testDB)
	assert.NoError(t, err)
	assert.Equal(t, balance.NetBalance, fetchedBalance.NetBalance)
}

func TestGetUserBalances(t *testing.T) {
	defer resetEnv(os.Environ())()
	testDB := SetupTestDB(t)
	defer testDB.Close()

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
	user3 := models.User{
		ProviderID:  "google-54321",
		Username:    "alice_smith",
		FirstName:   "Alice",
		LastName:    "Smith",
		Email:       "alice.smith@example.com",
		Provider:    "google",
		IsProxyUser: false,
	}

	userID1, err := db.InsertUser(user1, testDB)
	assert.NoError(t, err)
	userID2, err := db.InsertUser(user2, testDB)
	assert.NoError(t, err)
	userID3, err := db.InsertUser(user3, testDB)
	assert.NoError(t, err)

	balance1 := models.UserBalance{
		User1ID:    userID1,
		User2ID:    userID2,
		NetBalance: 100.00,
	}
	balance2 := models.UserBalance{
		User1ID:    userID1,
		User2ID:    userID3,
		NetBalance: -50.00,
	}

	_, err = db.InsertUserBalance(balance1, testDB)
	assert.NoError(t, err)
	_, err = db.InsertUserBalance(balance2, testDB)
	assert.NoError(t, err)

	balances, err := db.GetUserBalances(userID1, testDB)
	assert.NoError(t, err)
	assert.Len(t, balances, 2)
}
