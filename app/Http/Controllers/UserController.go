package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/pterm/pterm"
	"gorm.io/gorm"
)

type UserController struct {
	DB *gorm.DB
}

func (uc UserController) Index(w http.ResponseWriter, r *http.Request) {
	// Handle GET /users request
	// Sample json response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	resp := make(map[string]string)
	resp["message"] = "Status OK"
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		pterm.Error.Println("In UserController index:", err)
	}
	w.Write(jsonResp)
}
// swagger:response UserResponse
type UserResponse struct {
    // in:body
    Body struct {
        ID       int    `json:"id"`
        Name     string `json:"name"`
        Email    string `json:"email"`
        Password string `json:"password"`
    }
}
// swagger:response ErrorResponse
type ErrorResponse struct {
	Error string `json:"error"`
}

// @Summary Get user by ID
// @Description Get user information by user ID
// @ID get-user-by-id
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} UserResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /users/{id} [get]
func (uc UserController) Show(w http.ResponseWriter, r *http.Request) {
	// Handle GET /users/{id} request
}

func (uc UserController) Create(w http.ResponseWriter, r *http.Request) {
	// Handle POST /users request
	// You can get the request body by reading from r.Body
	// You can send a response by writing to w
}

func (uc UserController) Update(w http.ResponseWriter, r *http.Request) {
	// Handle PUT /users/{id} request
	// You can get the request body by reading from r.Body
	// You can send a response by writing to w
}

func (uc UserController) Delete(w http.ResponseWriter, r *http.Request) {
	// Handle DELETE /users/{id} request
}
