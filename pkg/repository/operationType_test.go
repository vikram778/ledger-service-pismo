// repository_test.go
package repository

import (
	"context"
	"errors"
	"testing"

	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"ledger-service-pismo/pkg/models"
)

func TestDBOps_GetOperationType(t *testing.T) {
	// Create a new mock DB and DBOps instance
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	dbOps := NewDBOpsRepository(sqlx.NewDb(db, "sqlmock")) // &dBOps{db: sqlx.NewDb(db, "sqlmock")}

	// Test case: Get operation type successfully
	mock.ExpectQuery("SELECT .* FROM operation_types WHERE id = ?").
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "Deposit"))

	result, err := dbOps.GetOperationType(context.Background(), 1)
	assert.NoError(t, err)
	assert.Equal(t, &models.Operations{OperationTypeID: 1, Description: "Deposit"}, result)

	// Test case: Operation type not found
	mock.ExpectQuery("SELECT .* FROM operation_types WHERE id = ?").
		WithArgs(2).
		WillReturnError(sql.ErrNoRows)

	result, err = dbOps.GetOperationType(context.Background(), 2)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.True(t, errors.Is(err, sql.ErrNoRows))
}
