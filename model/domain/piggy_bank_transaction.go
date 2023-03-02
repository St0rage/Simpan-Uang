package domain

type PiggyBankTransaction struct {
	Id              string  `json:"id" db:"id"`
	PiggyBankId     string  `json:"piggy_bank_id" db:"piggy_bank_id"`
	TransactionName string  `json:"transcation_name" db:"transaction_name"`
	Amount          float32 `json:"amount" db:"amount"`
	Status          bool    `json:"status" db:"status"`
	Date            int64   `json:"date" db:"date"`
}
