package services

import (
	"context"
	"ledger-service-pismo/pkg/errs"
	"ledger-service-pismo/pkg/models"
	"ledger-service-pismo/pkg/repository"
)

type LedgerService struct {
	DBOps repository.DBOps
}

func NewLedgerService(DBOps repository.DBOps) *LedgerService {
	return &LedgerService{
		DBOps: DBOps,
	}
}

func (s *LedgerService) CreateAccount(ctx context.Context, account models.Account) (string, error) {
	acct, err := s.DBOps.GetAccountByDocument(ctx, account.DocumentNumber)
	if err != nil {
		return "", err
	}

	if acct.AccountID != 0 {
		return "", errs.ErrorAccountExist
	}

	err = s.DBOps.CreateAccount(ctx, &account)
	if err != nil {
		return "", err
	}

	return "account created successfully!!!", nil
}

func (s *LedgerService) GetAccountByID(ctx context.Context, id int64) (*models.Account, error) {
	return s.DBOps.GetAccount(ctx, id)
}

func (s *LedgerService) CreateTransaction(ctx context.Context, transaction models.Transaction) (string, error) {
	_, err := s.DBOps.GetAccount(ctx, transaction.AccountID)
	if err != nil {
		return "", err
	}

	_, err = s.DBOps.GetOperationType(ctx, transaction.OperationTypeID)
	if err != nil {
		return "", errs.ErrorIncorrectOperationType
	}

	if transaction.OperationTypeID != 4 {
		transaction.Amount *= -1
		transaction.Balance *= -1
	} else {
		transaction.Balance, err = s.discharge(ctx, transaction.Amount, transaction.AccountID)
		if err != nil {
			return "", err
		}
	}

	err = s.DBOps.CreateTransaction(ctx, &transaction)
	if err != nil {
		return "", err
	}

	return "Success", err
}

func (s *LedgerService) discharge(ctx context.Context, amount float64, account_id int64) (float64, error) {
	// Fetching all the negative balance txns for account for discharging

	txns, err := s.DBOps.FetchAllTxnsByAccountId(ctx, account_id)
	if err != nil {
		return 0, err
	}

	for i := range txns {
		if amount == 0 {
			break
		}
		if amount+txns[i].Balance >= 0 {
			txns[i].Balance = 0
			amount += txns[i].Balance
		} else {
			txns[i].Balance += amount
			amount = 0
		}
		err = s.DBOps.UpdateBalanceByID(ctx, txns[i].TransactionID, txns[i].Balance)
		if err != nil {
			return 0, err
		}
	}

	return amount, nil
}
