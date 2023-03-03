package domain

type WishlistTransaction struct {
	Id              string  `json:"id" db:"id"`
	WishlistId     string  `json:"wishlist_id" db:"wishlist_id"`
	TransactionName string  `json:"transcation_name" db:"transaction_name"`
	Amount          float32 `json:"amount" db:"amount"`
	Status          bool    `json:"status" db:"status"`
	Date            int64   `json:"date" db:"date"`
}
