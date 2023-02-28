package web

type UserResponse struct {
	Id      string  `json:"id"`
	Name    string  `json:"name"`
	Email   string  `json:"email"`
	IsAdmin bool    `json:"is_admin" db:"is_admin"`
	Balance float32 `json:"balance"`
}
