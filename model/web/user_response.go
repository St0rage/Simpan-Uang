package web

type UserResponse struct {
	Id      string  `json:"id"`
	Name    string  `json:"name"`
	Email   string  `json:"email"`
	Balance float32 `json:"balance"`
}
