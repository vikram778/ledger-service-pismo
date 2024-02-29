package repository

import (
	"context"
	"ledger-service-pismo/pkg/models"

	"github.com/jmoiron/sqlx"
)

type dBOps struct {
	db *sqlx.DB
}

// NewDBOpsRepository initialises db operation repo
func NewDBOpsRepository(db *sqlx.DB) *dBOps {
	return &dBOps{db: db}
}

// DBOps interface defines all the operations performed on the DB
type DBOps interface {
	CreateAccount(ctx context.Context, profile *models.Account) error
	GetAccount(ctx context.Context, id int64) (*models.Account, error)
	GetAccountByDocument(ctx context.Context, docNo string) (*models.Account, error)
	GetOperationType(ctx context.Context, id int64) (*models.Operations, error)
	CreateTransaction(ctx context.Context, txn *models.Transaction) error
}
