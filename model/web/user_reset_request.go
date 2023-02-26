package web

type UserResetRequest struct {
	Email string `json:"email" binding:"required,email"`
}
