package web

type UserRegisterRequest struct {
	Name                 string `json:"name" binding:"required,min=3,max=30"`
	Email                string `json:"email" binding:"required,email"`
	Password             string `json:"password" binding:"required,min=4,eqfield=PasswordConfirmation"`
	PasswordConfirmation string `json:"password_confirmation" binding:"required"`
}
