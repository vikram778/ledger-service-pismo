package models

import "time"

type Transaction struct {
	TransactionID   string    `json:"transaction_id" db:"transaction_id"`       // TransactionID is the unique transaction ID generated for each transaction
	AccountID       int64     `json:"account_id" db:"account_id"`               // AccountID is the account ID on which the transaction is performed for
	OperationTypeID int64     `json:"operation_type_id" db:"operation_type_id"` // OperationTypeID identifies the type of transaction performed e.g. credit/debit
	Amount          float64   `json:"amount" db:"amount"`                       // Amount is the amount of transaction is performed
	EventDate       time.Time `json:"event_date" db:"event_date"`               // EventDate date of transaction performed
}
