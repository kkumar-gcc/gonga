package auth

import (
	"gonga/utils"
	"log"
	"net/http"

	"gorm.io/gorm"
)

type LoginRequest struct {
	Username string `json:"username" validate:"required,min=3,max=5"`
	Password string `json:"password" validate:"required,min=8"`
}

type LoginResponse struct {
	Token   string `json:"token"`
	UserID  int    `json:"user_id"`
	Message string `json:"message"`
}

type LoginController struct {
	DB *gorm.DB
}

func (c LoginController) Index(w http.ResponseWriter, r *http.Request) {
	// Handle GET /logincontroller request
}

func (c LoginController) Show(w http.ResponseWriter, r *http.Request) {
	// Handle GET /logincontroller/{id} request
}
/*
@Summary [Short summary of API endpoint]
@Description [Detailed description of API endpoint]
@Tags [Comma-separated list of API tags]
@Accept [List of accepted MIME types]
@Produce [List of produced MIME types]
@Param [Name of input parameter] [Parameter location] [Data type] [Required?] [Description]
@Success [HTTP status code] [Response data type] [Response description]
@Failure [HTTP status code] [Response data type] [Response description]
@Router [API endpoint URL] [HTTP method]
*/
func (c LoginController) Create(w http.ResponseWriter, r *http.Request) {
	// Handle POST /logincontroller request
	// You can get the request body by reading from r.Body
	// You can send a response by writing to w
	// Parse request body
	var user LoginRequest

	if err := utils.DecodeRequestBody(r, &user); err != nil {
		utils.JSONResponse(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	// Validate user data
	log.Println(user)
	if err := utils.ValidateRequest(w, &user); err != nil {
		return
	}
	// Check credentials and get user ID from database
	userID, err := utils.Authenticate(user.Username, user.Password, c.DB)
	if err != nil {
		utils.JSONResponse(w, http.StatusUnauthorized, map[string]string{"error": "Invalid username or password"})
		return
	}

	// Generate JWT token
	token, err := utils.GenerateToken(userID)
	if err != nil {
		utils.JSONResponse(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	// Send response
	response := LoginResponse{
		Token:   token,
		UserID:  userID,
		Message: "Login successful",
	}
	utils.JSONResponse(w, http.StatusOK, response)
}

func (c LoginController) Update(w http.ResponseWriter, r *http.Request) {
	// Handle PUT /logincontroller/{id} request
	// You can get the request body by reading from r.Body
	// You can send a response by writing to w
}

func (c LoginController) Delete(w http.ResponseWriter, r *http.Request) {
	// Handle DELETE /logincontroller/{id} request
}
