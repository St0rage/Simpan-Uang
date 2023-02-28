package web

type WishlistIdRequest struct {
	Id       		string `json:"id"`
	
	WishlistName 	string `json:"wishlist_name" db:"wishlist_name"`
	WishlistTarget 	float32 `json:"wishlist_target" db:"wishlist_target"`
	Progress       	int 	`json:"progress"`
	Total			float32 `json:"total"`
}

	// WishlistId	 	string  `json:"wishlist_id" db:"wishlist_id"`
	// TransactionName 	string  `json:"transaction_name" db:"transaction_name"`
	// Amount 			float32 `json:"amount"`
	// Status       	bool 	`json:"status"`
	// Date				int	    `json:"total"`
