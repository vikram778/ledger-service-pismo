// dbops_test.go
package repository

import (
	"context"
	_ "database/sql"
	"errors"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"ledger-service-pismo/pkg/models"
)

func TestDBOps_CreateAccount(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock database connection: %v", err)
	}
	defer db.Close()

	dBRepo := NewDBOpsRepository(sqlx.NewDb(db, "sqlmock"))

	// Test case: Create account successfully
	mock.ExpectExec(createAccountQuery).
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = dBRepo.CreateAccount(context.Background(), &models.Account{DocumentNumber: "123"})
	assert.NoError(t, err)

	// Test case: Error during account creation
	mock.ExpectExec(createAccountQuery).
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg()).
		WillReturnError(errors.New("some error"))

	err = dBRepo.CreateAccount(context.Background(), &models.Account{DocumentNumber: "456"})
	assert.Error(t, err)
	assert.Equal(t, "some error", err.Error())
}

func TestDBOps_GetAccount(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock database connection: %v", err)
	}
	defer db.Close()

	dBRepo := NewDBOpsRepository(sqlx.NewDb(db, "sqlmock"))

	// Test case: Get account by ID successfully
	rows := sqlmock.NewRows([]string{"id", "document_number", "created_at"}).
		AddRow(1, "123", time.Now())

	mock.ExpectQuery(getAccountQuery).
		WithArgs(sqlmock.AnyArg()).
		WillReturnRows(rows)

	account, err := dBRepo.GetAccount(context.Background(), 1)
	assert.NoError(t, err)
	assert.Equal(t, &models.Account{AccountID: 1, DocumentNumber: "123", CreatedAt: time.Now()}, account)

	// Test case: Error during fetching account
	mock.ExpectQuery(getAccountQuery).
		WithArgs(sqlmock.AnyArg()).
		WillReturnError(errors.New("some error"))

	account, err = dBRepo.GetAccount(context.Background(), 2)
	assert.Error(t, err)
	assert.Equal(t, &models.Account{}, account)
	assert.Equal(t, "some error", err.Error())
}

func TestDBOps_GetAccountByDocument(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock database connection: %v", err)
	}
	defer db.Close()

	dBRepo := NewDBOpsRepository(sqlx.NewDb(db, "sqlmock"))

	// Test case: Get account by document number successfully
	rows := sqlmock.NewRows([]string{"id", "document_number", "created_at"}).
		AddRow(1, "123", time.Now())

	mock.ExpectQuery(getAccountByDocumentQuery).
		WithArgs(sqlmock.AnyArg()).
		WillReturnRows(rows)

	account, err := dBRepo.GetAccountByDocument(context.Background(), "123")
	assert.NoError(t, err)
	assert.Equal(t, &models.Account{AccountID: 1, DocumentNumber: "123", CreatedAt: time.Now()}, account)

	// Test case: Error during fetching account by document number
	mock.ExpectQuery(getAccountByDocumentQuery).
		WithArgs(sqlmock.AnyArg()).
		WillReturnError(errors.New("some error"))

	account, err = dBRepo.GetAccountByDocument(context.Background(), "456")
	assert.Error(t, err)
	assert.Equal(t, &models.Account{}, account)
	assert.Equal(t, "some error", err.Error())
}
