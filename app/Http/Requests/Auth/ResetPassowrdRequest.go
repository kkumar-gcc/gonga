package requests


type ResetPassowrdRequest struct {
	Email    string `json:"email" validate:"required,email"`
}