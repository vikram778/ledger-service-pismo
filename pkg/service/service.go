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
	}

	err = s.DBOps.CreateTransaction(ctx, &transaction)
	if err != nil {
		return "", err
	}

	return "Success", err
}
