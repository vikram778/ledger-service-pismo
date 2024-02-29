// ledger_service_test.go
package services

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"ledger-service-pismo/pkg/errs"
	"ledger-service-pismo/pkg/models"
)

type mockDBOps struct {
	CreateAccountFunc        func(ctx context.Context, account *models.Account) error
	GetAccountByDocumentFunc func(ctx context.Context, documentNumber string) (*models.Account, error)
	GetAccountFunc           func(ctx context.Context, id int64) (*models.Account, error)
	CreateTransactionFunc    func(ctx context.Context, transaction *models.Transaction) error
	GetOperationTypeFunc     func(ctx context.Context, operationTypeID int64) (*models.Operations, error)
}

func (m *mockDBOps) CreateAccount(ctx context.Context, account *models.Account) error {
	return m.CreateAccountFunc(ctx, account)
}

func (m *mockDBOps) GetAccountByDocument(ctx context.Context, documentNumber string) (*models.Account, error) {
	return m.GetAccountByDocumentFunc(ctx, documentNumber)
}

func (m *mockDBOps) GetAccount(ctx context.Context, id int64) (*models.Account, error) {
	return m.GetAccountFunc(ctx, id)
}

func (m *mockDBOps) CreateTransaction(ctx context.Context, transaction *models.Transaction) error {
	return m.CreateTransactionFunc(ctx, transaction)
}

func (m *mockDBOps) GetOperationType(ctx context.Context, operationTypeID int64) (*models.Operations, error) {
	return m.GetOperationTypeFunc(ctx, operationTypeID)
}

func TestLedgerService_CreateAccount(t *testing.T) {
	mockRepo := &mockDBOps{}

	// Initialize the service with the mock repository
	service := NewLedgerService(mockRepo)

	// Test case: Create a new account successfully
	mockRepo.CreateAccountFunc = func(ctx context.Context, account *models.Account) error {
		return nil
	}

	mockRepo.GetAccountByDocumentFunc = func(ctx context.Context, documentNumber string) (*models.Account, error) {
		return &models.Account{}, nil
	}

	result, err := service.CreateAccount(context.Background(), models.Account{})
	assert.NoError(t, err)
	assert.Equal(t, "account created successfully!!!", result)

	// Test case: Account already exists
	mockRepo.CreateAccountFunc = func(ctx context.Context, account *models.Account) error {
		return errs.ErrorAccountExist
	}

	mockRepo.GetAccountByDocumentFunc = func(ctx context.Context, documentNumber string) (*models.Account, error) {
		return &models.Account{AccountID: 1}, nil
	}

	result, err = service.CreateAccount(context.Background(), models.Account{})
	assert.Error(t, err)
	assert.Equal(t, "", result)
	assert.Equal(t, errs.ErrorAccountExist, errors.Unwrap(err))
}

func TestLedgerService_GetAccountByID(t *testing.T) {
	mockRepo := &mockDBOps{}

	// Initialize the service with the mock repository
	service := NewLedgerService(mockRepo)

	// Test case: Get account by ID successfully
	mockRepo.GetAccountFunc = func(ctx context.Context, id int64) (*models.Account, error) {
		return &models.Account{AccountID: 1, DocumentNumber: "123"}, nil
	}

	result, err := service.GetAccountByID(context.Background(), 1)
	assert.NoError(t, err)
	assert.Equal(t, &models.Account{AccountID: 1, DocumentNumber: "123"}, result)

	// Test case: Account not found
	mockRepo.GetAccountFunc = func(ctx context.Context, id int64) (*models.Account, error) {
		return nil, errs.ErrorAccountNotExist
	}

	result, err = service.GetAccountByID(context.Background(), 2)
	assert.Error(t, err)
	assert.Equal(t, &models.Account{}, result)
	assert.Equal(t, errs.ErrorAccountNotExist, errors.Unwrap(err))
}

func TestLedgerService_CreateTransaction(t *testing.T) {
	mockRepo := &mockDBOps{}

	// Initialize the service with the mock repository
	service := NewLedgerService(mockRepo)

	// Test case: Create a new transaction successfully
	mockRepo.GetAccountFunc = func(ctx context.Context, id int64) (*models.Account, error) {
		return &models.Account{AccountID: 1}, nil
	}

	mockRepo.GetOperationTypeFunc = func(ctx context.Context, operationTypeID int64) (*models.Operations, error) {
		return &models.Operations{}, nil
	}

	mockRepo.CreateTransactionFunc = func(ctx context.Context, transaction *models.Transaction) error {
		return nil
	}

	result, err := service.CreateTransaction(context.Background(), models.Transaction{})
	assert.NoError(t, err)
	assert.Equal(t, "Success", result)

	// Test case: Incorrect operation type
	mockRepo.GetOperationTypeFunc = func(ctx context.Context, operationTypeID int64) (*models.Operations, error) {
		return nil, errs.ErrorIncorrectOperationType
	}

	result, err = service.CreateTransaction(context.Background(), models.Transaction{})
	assert.Error(t, err)
	assert.Equal(t, "", result)
	assert.Equal(t, errs.ErrorIncorrectOperationType, errors.Unwrap(err))
}
