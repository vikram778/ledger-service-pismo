// ledger_controller_test.go
package controllers

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"ledger-service-pismo/pkg/errs"
	"ledger-service-pismo/pkg/models"
)

type mockLedgerService struct{}

func (m *mockLedgerService) CreateAccount(ctx interface{}, account models.Account) (int64, error) {
	return 1, nil
}

func (m *mockLedgerService) GetAccountByID(ctx interface{}, accountID int64) (*models.Account, error) {
	if accountID == 1 {
		return &models.Account{AccountID: 1, DocumentNumber: "TestAccount"}, nil
	}
	return nil, errs.ErrorAccountNotExist
}

func (m *mockLedgerService) CreateTransaction(ctx interface{}, transaction models.Transaction) (int64, error) {
	return 1, nil
}

func TestLedgerController_CreateAccountHandler(t *testing.T) {
	mockService := &mockLedgerService{}
	controller := NewLedgerController(mockService)

	// Prepare a request with a JSON payload
	account := models.Account{AccountID: 1, DocumentNumber: "123"}
	payload, _ := json.Marshal(account)
	request, _ := http.NewRequest("POST", "/create-account", bytes.NewBuffer(payload))
	request.Header.Set("Content-Type", "application/json")

	// Create a response recorder
	recorder := httptest.NewRecorder()

	// Call the handler function
	controller.CreateAccountHandler(recorder, request)

	// Assertions
	assert.Equal(t, http.StatusOK, recorder.Code)

	// Parse the response body to check if it contains the expected data
	var status string
	err := json.Unmarshal(recorder.Body.Bytes(), &status)
	assert.NoError(t, err)
	assert.Equal(t, "account created successfully!!!", status)
}

func TestLedgerController_GetAccountHandler(t *testing.T) {
	mockService := &mockLedgerService{}
	controller := NewLedgerController(mockService)

	// Prepare a request with a path parameter
	request, _ := http.NewRequest("GET", "/get-account/1", nil)

	// Create a response recorder
	recorder := httptest.NewRecorder()

	// Call the handler function
	controller.GetAccountHandler(recorder, request)

	// Assertions
	assert.Equal(t, http.StatusOK, recorder.Code)

	// Parse the response body to check if it contains the expected data
	var account models.Account
	err := json.Unmarshal(recorder.Body.Bytes(), &account)
	assert.NoError(t, err)
	assert.Equal(t, int64(1), account.AccountID)
	assert.Equal(t, "123", account.DocumentNumber)
}

func TestLedgerController_CreateTransactionHandler(t *testing.T) {
	mockService := &mockLedgerService{}
	controller := NewLedgerController(mockService)

	// Prepare a request with a JSON payload
	transaction := models.Transaction{Amount: 100.0}
	payload, _ := json.Marshal(transaction)
	request, _ := http.NewRequest("POST", "/create-transaction", bytes.NewBuffer(payload))
	request.Header.Set("Content-Type", "application/json")

	// Create a response recorder
	recorder := httptest.NewRecorder()

	// Call the handler function
	controller.CreateTransactionHandler(recorder, request)

	// Assertions
	assert.Equal(t, http.StatusOK, recorder.Code)

	// Parse the response body to check if it contains the expected data
	var status string
	err := json.Unmarshal(recorder.Body.Bytes(), &status)
	assert.NoError(t, err)
	assert.Equal(t, "success", status)
}

func TestLedgerController_GetParams(t *testing.T) {
	mockService := &mockLedgerService{}
	controller := NewLedgerController(mockService)

	// Prepare a request with a JSON payload
	account := models.Account{AccountID: 1}
	payload, _ := json.Marshal(account)
	request, _ := http.NewRequest("POST", "/test", bytes.NewBuffer(payload))
	request.Header.Set("Content-Type", "application/json")

	// Call the GetParams method
	var target models.Account
	err := controller.GetParams(&target, request)

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, account.DocumentNumber, target.DocumentNumber)
}

func TestLedgerController_GetParams_InvalidContentType(t *testing.T) {
	mockService := &mockLedgerService{}
	controller := NewLedgerController(mockService)

	// Prepare a request with an unsupported content type
	request, _ := http.NewRequest("POST", "/test", nil)
	request.Header.Set("Content-Type", "text/plain")

	// Call the GetParams method
	var target models.Account
	err := controller.GetParams(&target, request)

	// Assertions
	assert.Error(t, err)
	assert.Equal(t, errors.New("unsupported media type"), err)
}
