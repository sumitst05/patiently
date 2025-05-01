package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/sumitst05/patiently/internal/repository"
)

func init() {
	repository.InitDB()
}

func getAuthToken(t *testing.T, r *gin.Engine) *http.Cookie {
	payload := map[string]string{
		"email":    "johndoe@example.com",
		"password": "password123",
	}
	body, _ := json.Marshal(payload)

	req := httptest.NewRequest(http.MethodPost, "/api/auth/signin", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	for _, cookie := range w.Result().Cookies() {
		if cookie.Name == "token" {
			return cookie
		}
	}
	t.Fatal("token cookie not found")
	return nil
}

func TestCreatePatient_Success(t *testing.T) {
	r := setupTestRouter()
	token := getAuthToken(t, r)

	fmt.Println(token)

	payload := map[string]any{
		"name":    "Test Patient",
		"age":     30,
		"gender":  "male",
		"address": "test patient address",
		"phone":   "1234567890",
	}
	body, _ := json.Marshal(payload)

	req := httptest.NewRequest(http.MethodPost, "/api/patient/create", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	req.AddCookie(token)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	fmt.Println("Create response body:", w.Body.String())

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), `"name":"Test Patient"`)
	assert.Contains(t, w.Body.String(), `"created_by"`)
}

func TestGetAllPatients_Success(t *testing.T) {
	r := setupTestRouter()
	token := getAuthToken(t, r)

	req := httptest.NewRequest(http.MethodGet, "/api/patient/fetch", nil)
	req.AddCookie(token)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Test Patient")
}

func TestGetPatientById_Success(t *testing.T) {
	r := setupTestRouter()
	token := getAuthToken(t, r)

	req := httptest.NewRequest(http.MethodGet, "/api/patient/fetch/1", nil)
	req.AddCookie(token)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), `"id":1`)
}

func TestGetPatientById_NotFound(t *testing.T) {
	r := setupTestRouter()
	token := getAuthToken(t, r)

	req := httptest.NewRequest(http.MethodGet, "/api/patient/fetch/9999", nil)
	req.AddCookie(token)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Contains(t, w.Body.String(), "Patient not found")
}

func TestUpdatePatient_Success(t *testing.T) {
	r := setupTestRouter()
	token := getAuthToken(t, r)

	payload := map[string]any{
		"name": "Updated Patient",
		"age":  35,
	}
	body, _ := json.Marshal(payload)

	req := httptest.NewRequest(http.MethodPost, "/api/patient/update/1", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	req.AddCookie(token)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), `"name":"Updated Patient"`)
	assert.Contains(t, w.Body.String(), `"age":35`)
}

func TestDeletePatient_Success(t *testing.T) {
	r := setupTestRouter()
	token := getAuthToken(t, r)

	req := httptest.NewRequest(http.MethodDelete, "/api/patient/delete/1", nil)
	req.AddCookie(token)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Patient deleted successfully")
}

func TestGetPatientRegistrationHistory(t *testing.T) {
	r := setupTestRouter()
	token := getAuthToken(t, r)

	req := httptest.NewRequest(http.MethodGet, "/api/patient/fetch/1/history", nil)
	req.AddCookie(token)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
