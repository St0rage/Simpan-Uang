package domain

type Wishlist struct {
	Id             string  `json:"id"`
	UserId         string  `json:"user_id" db:"user_id"`
	WishlistName   string  `json:"wishlist_name" db:"wishlist_name"`
	WishlistTarget float32 `json:"wishlist_target" db:"wishlist_target"`
}
