package web

type PiggyBankCreateUpdateRequest struct {
	PiggyBankName string `json:"piggy_bank_name" binding:"required,min=3,max=15"`
}
