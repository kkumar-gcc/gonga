package responses

type RegisterResponse struct {
	Token   string `json:"token"`
	UserID  int    `json:"user_id"`
	Message string `json:"message"`
}