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

// FetchAllTxnsByAccountId creates transaction in the system for the given transaction request
// Param - context object
// Param - accountID  is the account Id for which the transactions needs to be fetched
// Returns - slice of all transactions
// Returns - error if any
func (r *dBOps) FetchAllTxnsByAccountId(ctx context.Context, accountID int64) ([]models.Transaction, error) {
	rows, err := r.db.Query(fetchAllTxnByAccountIdQuery, accountID)
	if err != nil {
		log.Error("Error fetching txns for accountID", zap.Error(err))
		return nil, err
	}

	defer rows.Close()

	var txns []models.Transaction

	for rows.Next() {
		var txn models.Transaction

		err = rows.Scan(&txn.TransactionID, &txn.Balance)
		if err != nil {
			log.Error("Error scaning rows", zap.Error(err))
			return nil, err
		}

		txns = append(txns, txn)
	}

	if err = rows.Err(); err != nil {
		log.Error("Error iterating overs rows", zap.Error(err))
		return nil, err
	}

	return txns, nil
}

// UpdateBalanceByID creates transaction in the system for the given transaction request
// Param - context object
// Param - transaction_id  is transactionID against which update needs to be performed
// Param - balance is balance that needs to be upadated in the transaction ledger
// Returns - error if any
func (r *dBOps) UpdateBalanceByID(ctx context.Context, transaction_id string, balance float64) error {
	_, err := r.db.Exec(updateBalanceByIDQuery, balance, transaction_id)
	if err != nil {
		log.Error("Error Updating the balance ", zap.Error(err))
		return err
	}

	return nil
}
