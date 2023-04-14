package controllers

import (
	"net/http"

	"gorm.io/gorm"
)

type PostController struct {
	DB *gorm.DB
}

func (c PostController) Index(w http.ResponseWriter, r *http.Request) {
	// Handle GET /postcontroller request
}

func (c PostController) Show(w http.ResponseWriter, r *http.Request) {
	// Handle GET /postcontroller/{id} request
}

func (c PostController) Create(w http.ResponseWriter, r *http.Request) {
	// Handle POST /postcontroller request
	// You can get the request body by reading from r.Body
	// You can send a response by writing to w
}

func (c PostController) Update(w http.ResponseWriter, r *http.Request) {
	// Handle PUT /postcontroller/{id} request
	// You can get the request body by reading from r.Body
	// You can send a response by writing to w
}

func (c PostController) Delete(w http.ResponseWriter, r *http.Request) {
	// Handle DELETE /postcontroller/{id} request
}
