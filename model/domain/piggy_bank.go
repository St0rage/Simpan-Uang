package domain

type PiggyBank struct {
	Id            string `json:"id" db:"id"`
	UserId        string `json:"user_id" db:"user_id"`
	PiggyBankName string `json:"piggy_bank_name" db:"piggy_bank_name"`
	Type          bool   `json:"type" db:"type"`
}
