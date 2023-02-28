package web

type UserUpdateRequest struct {
	Name  string `json:"name" binding:"required,min=3,max=30"`
	Email string `json:"email" binding:"required,email"`
}
