package repository

import (
	"context"
	"database/sql"
	"ledger-service-pismo/pkg/log"
	"ledger-service-pismo/pkg/models"
	"time"

	"go.uber.org/zap"
)

// CreateAccount creates account for the given request
// Param - context object
// Param - profile type account model is the profile request for which account is to be created
// Returns - error if any
func (r *dBOps) CreateAccount(ctx context.Context, profile *models.Account) error {

	profile.CreatedAt = time.Now()

	_, err := r.db.Exec(
		createAccountQuery,
		profile.DocumentNumber,
		profile.CreatedAt,
	)

	if err != nil && err != sql.ErrNoRows {
		log.Info("Create account error", zap.Error(err))
		return err
	}

	return nil
}

// GetAccount fetches account based on the ID
// Param - context object
// Param - id type int64 for which account details needs to be fetched
// Returns - account model
// Returns - error if any
func (r *dBOps) GetAccount(ctx context.Context, id int64) (*models.Account, error) {

	var account models.Account

	if err := r.db.Get(&account, getAccountQuery, id); err != nil {
		log.Info("fetch account error", zap.Error(err))
		return nil, err
	}
	return &account, nil
}

// GetAccountByDocument fetches account based on the document no
// Param - context object
// Param - docNo type string for which account details needs to be fetched
// Returns - account model
// Returns - error if any
func (r *dBOps) GetAccountByDocument(ctx context.Context, docNo string) (*models.Account, error) {

	var account models.Account

	if err := r.db.Get(&account, getAccountByDocumentQuery, docNo); err != nil && err != sql.ErrNoRows {
		log.Info("fetch account bu document error", zap.Error(err))
		return nil, err
	}
	return &account, nil
}
