package web

type DepositTransactionRequest struct {
	Amount float32 `json:"amount" binding:"required,numeric"`
}
