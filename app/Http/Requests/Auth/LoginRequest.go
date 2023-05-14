package requests

type LoginRequest struct {
	Username string `json:"username" validate:"required,min=3,max=5"`
	Password string `json:"password" validate:"required,min=8"`
}
