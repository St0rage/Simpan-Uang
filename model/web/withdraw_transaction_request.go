package web

type WithdrawTransactionRequest struct {
	Amount          float32 `json:"amount" binding:"required,numeric"`
	TransactionName string  `json:"transaction_name" binding:"required,min=3,max=15"`
}
