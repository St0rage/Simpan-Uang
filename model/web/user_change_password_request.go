package web

type UserChangePasswordRequest struct {
	Password             string `json:"password" binding:"required,min=4,eqfield=PasswordConfirmation"`
	PasswordConfirmation string `json:"password_confirmation" binding:"required"`
}
