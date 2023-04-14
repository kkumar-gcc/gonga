package auth

import (
	"net/http"

	"gorm.io/gorm"
)

type PasswordResetLinkController struct {
	DB *gorm.DB
}

func (c PasswordResetLinkController) Index(w http.ResponseWriter, r *http.Request) {
	// Handle GET /passwordresetlinkcontroller request
}

func (c PasswordResetLinkController) Show(w http.ResponseWriter, r *http.Request) {
	// Handle GET /passwordresetlinkcontroller/{id} request
}

func (c PasswordResetLinkController) Create(w http.ResponseWriter, r *http.Request) {
	// Handle POST /passwordresetlinkcontroller request
	// You can get the request body by reading from r.Body
	// You can send a response by writing to w
}

func (c PasswordResetLinkController) Update(w http.ResponseWriter, r *http.Request) {
	// Handle PUT /passwordresetlinkcontroller/{id} request
	// You can get the request body by reading from r.Body
	// You can send a response by writing to w
}

func (c PasswordResetLinkController) Delete(w http.ResponseWriter, r *http.Request) {
	// Handle DELETE /passwordresetlinkcontroller/{id} request
}
