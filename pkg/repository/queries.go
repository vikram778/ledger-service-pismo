package repository

const (
	createAccountQuery          = `INSERT INTO accounts (document_number, created_at) VALUES ($1, $2)`
	getAccountQuery             = `SELECT * FROM accounts WHERE account_id = $1`
	getAccountByDocumentQuery   = `SELECT * FROM accounts WHERE document_number = $1`
	getOperationTypeQuery       = `SELECT * FROM operation_types WHERE operation_type_id = $1`
	createTransactionQuery      = `INSERT INTO transactions (account_id, operation_type_id, amount, event_date) VALUES ($1, $2, $3, $4)`
	fetchAllTxnByAccountIdQuery = `SELECT transaction_id,balance FROM transactions WHERE account_id = $1 AND balance < 0 ORDER By event_date ASC`
	updateBalanceByIDQuery      = `UPDATE transactions SET balance = $1 WHERE transaction_id = $2`
)
