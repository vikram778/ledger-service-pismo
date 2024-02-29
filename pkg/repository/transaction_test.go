// repository_test.go
package repository

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"ledger-service-pismo/pkg/models"
)

func TestDBOps_CreateTransaction(t *testing.T) {
	// Create a new mock DB and DBOps instance
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	dbOps := NewDBOpsRepository(sqlx.NewDb(db, "sqlmock"))

	// Test case: Create transaction successfully
	mock.ExpectExec("INSERT INTO transactions (.+)").
		WithArgs(1, 1, 100.0, sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = dbOps.CreateTransaction(context.Background(), &models.Transaction{
		AccountID:       1,
		OperationTypeID: 1,
		Amount:          100.0,
		EventDate:       time.Now(),
	})
	assert.NoError(t, err)

	// Test case: Error while creating transaction
	mock.ExpectExec("INSERT INTO transactions (.+)").
		WillReturnError(errors.New("database error"))

	err = dbOps.CreateTransaction(context.Background(), &models.Transaction{
		AccountID:       1,
		OperationTypeID: 1,
		Amount:          100.0,
		EventDate:       time.Now(),
	})
	assert.Error(t, err)
	assert.Equal(t, "database error", err.Error())
}
