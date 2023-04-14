package auth

import (
	"net/http"

	"gorm.io/gorm"
)

type NewPasswordController struct {
	DB *gorm.DB
}

func (c NewPasswordController) Index(w http.ResponseWriter, r *http.Request) {
	// Handle GET /newpasswordcontroller request
}

func (c NewPasswordController) Show(w http.ResponseWriter, r *http.Request) {
	// Handle GET /newpasswordcontroller/{id} request
}

func (c NewPasswordController) Create(w http.ResponseWriter, r *http.Request) {
	// Handle POST /newpasswordcontroller request
	// You can get the request body by reading from r.Body
	// You can send a response by writing to w
}

func (c NewPasswordController) Update(w http.ResponseWriter, r *http.Request) {
	// Handle PUT /newpasswordcontroller/{id} request
	// You can get the request body by reading from r.Body
	// You can send a response by writing to w
}

func (c NewPasswordController) Delete(w http.ResponseWriter, r *http.Request) {
	// Handle DELETE /newpasswordcontroller/{id} request
}
