package repository

const (
	createAccountQuery        = `INSERT INTO accounts (document_number, created_at) VALUES ($1, $2)`
	getAccountQuery           = `SELECT * FROM accounts WHERE account_id = $1`
	getAccountByDocumentQuery = `SELECT * FROM accounts WHERE document_number = $1`
	getOperationTypeQuery     = `SELECT * FROM operation_types WHERE operation_type_id = $1`
	createTransactionQuery    = `INSERT INTO transactions (account_id, operation_type_id, amount, event_date) VALUES ($1, $2, $3, $4)`
)
