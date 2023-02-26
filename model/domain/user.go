package domain

type User struct {
	Id       string `json:"id"`
	Name     string `json:"name" binding:"required,min=3,max=30"`
	Email    string `json:"email" binding:"required,email,unique"`
	Password string `json:"password" binding:"required,min=4"`
}
