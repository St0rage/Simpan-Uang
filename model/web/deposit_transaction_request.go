package web

type DepositTransactionRequest struct {
	Amount any `json:"amount" binding:"required,numeric,gt=500"`
}
