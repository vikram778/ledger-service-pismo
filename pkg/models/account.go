package models

import "time"

type Account struct {
	AccountID      int64     `json:"account_id" db:"account_id"`           // AccountId is the account ID of the user's account
	DocumentNumber string    `json:"document_number" db:"document_number"` // DocumentNumber is the unique number associated with each account
	CreatedAt      time.Time `json:"created_at" db:"created_at"`           // CreatedAt is the timestamp at which account is created
}
