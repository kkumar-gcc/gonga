package requests

type RegisterRequest struct {
	Username string `json:"username" validate:"required,min=3,max=5"`
	Password string `json:"password" validate:"required,min=8"`
	Email    string `json:"email" validate:"required,email"`
}