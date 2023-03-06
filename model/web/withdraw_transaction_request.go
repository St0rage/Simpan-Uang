package web

type WithdrawTransactionRequest struct {
	Amount          any    `json:"amount" binding:"required,numeric,gt=500"`
	TransactionName string `json:"transaction_name" binding:"required,min=3,max=15"`
}
