package tests

import (
	"Tab_Tracker/models"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEndToEnd(t *testing.T) {
	defer resetEnv(os.Environ())()
	db := SetupTestDB(t)
	defer db.Close()

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
	var createdUser map[string]int
	err := json.NewDecoder(rr.Body).Decode(&createdUser)
	assert.NoError(t, err)

	userID := createdUser["user_id"]

	req, _ = http.NewRequest("GET", "/users/"+strconv.Itoa(userID), nil)
	rr = httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	var fetchedUser models.User
	err = json.NewDecoder(rr.Body).Decode(&fetchedUser)
	assert.NoError(t, err)
	assert.Equal(t, user.Email, fetchedUser.Email)
}
