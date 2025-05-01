package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sumitst05/patiently/internal/repository"
)

func init() {
	repository.InitDB()
}

func TestSignup_Success(t *testing.T) {
	r := setupTestRouter()

	// request payload
	payload := map[string]any{
		"name":     "John Doe",
		"email":    "johndoe@example.com",
		"password": "password123",
		"role":     "receptionist",
	}
	body, _ := json.Marshal(payload)

	req := httptest.NewRequest(http.MethodPost, "/api/auth/signup", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Contains(t, w.Body.String(), "id")
	assert.Contains(t, w.Body.String(), "email")
}

func TestSignup_InvalidRole(t *testing.T) {
	r := setupTestRouter()

	payload := map[string]any{
		"name":     "Jane Doe",
		"email":    "janedoe@example.com",
		"password": "password123",
		"role":     "invalidrole",
	}
	body, _ := json.Marshal(payload)

	req := httptest.NewRequest(http.MethodPost, "/api/auth/signup", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "Invalid role")
}

func TestSignin_Success(t *testing.T) {
	r := setupTestRouter()

	// first register a user
	signupPayload := map[string]any{
		"name":     "John Login",
		"email":    "johnlogin@example.com",
		"password": "password123",
		"role":     "doctor",
	}
	signupBody, _ := json.Marshal(signupPayload)

	req := httptest.NewRequest(http.MethodPost, "/api/auth/signup", bytes.NewReader(signupBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)

	// now login
	signinPayload := map[string]any{
		"email":    "johnlogin@example.com",
		"password": "password123",
	}
	signinBody, _ := json.Marshal(signinPayload)

	req = httptest.NewRequest(http.MethodPost, "/api/auth/signin", bytes.NewReader(signinBody))
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Login successful")

	// check cookie is set
	cookies := w.Result().Cookies()
	var hasToken bool
	for _, cookie := range cookies {
		if cookie.Name == "token" && cookie.Value != "" {
			hasToken = true
		}
	}
	assert.True(t, hasToken, "JWT token cookie should be set")
}

func TestSignin_InvalidCredentials(t *testing.T) {
	r := setupTestRouter()

	payload := map[string]any{
		"email":    "nonexistent@example.com",
		"password": "wrongpassword",
	}
	body, _ := json.Marshal(payload)

	req := httptest.NewRequest(http.MethodPost, "/api/auth/signin", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.Contains(t, w.Body.String(), "Invalid credentials")
}

func TestLogout(t *testing.T) {
	r := setupTestRouter()

	req := httptest.NewRequest(http.MethodPost, "/api/auth/logout", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Logged out")

	// check cookie is cleared
	cookies := w.Result().Cookies()
	var cleared bool
	for _, cookie := range cookies {
		if cookie.Name == "token" && cookie.Value == "" && cookie.MaxAge == -1 {
			cleared = true
		}
	}
	assert.True(t, cleared, "token cookie should be cleared")
}
