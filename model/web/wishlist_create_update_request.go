package web

type WishlistCreateUpdateRequest struct {
	WishlistName   string `json:"wishlist_name" binding:"required,min=3,max=15"`
	WishlistTarget any    `json:"wishlist_target" binding:"required,numeric,gt=10000"`
}
