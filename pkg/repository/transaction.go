package repository

import (
	"context"
	"go.uber.org/zap"
	"ledger-service-pismo/pkg/log"
	"ledger-service-pismo/pkg/models"
	"time"
)

// CreateTransaction creates transaction in the system for the given transaction request
// Param - context object
// Param - transaction  model is the transaction request for which transaction is to be created
// Returns - error if any
func (r *dBOps) CreateTransaction(ctx context.Context, txn *models.Transaction) error {

	txn.EventDate = time.Now()
	_, err := r.db.Exec(
		createTransactionQuery,
		txn.AccountID,
		txn.OperationTypeID,
		txn.Amount,
		txn.EventDate,
	)

	if err != nil {
		log.Error("Create transaction error", zap.Error(err))
		return err
	}

	return nil
}
