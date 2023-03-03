package web

type WishlistUpdateRequest struct {
	WishlistName   string  `json:"wishlist_name" db:"wishlist_name"`
	WishlistTarget float32 `json:"wishlist_target" db:"wishlist_target"`
}
