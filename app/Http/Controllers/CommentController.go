package controllers

import (
	"net/http"

	"gorm.io/gorm"
)

type CommentController struct {
	DB *gorm.DB
}

func (c CommentController) Index(w http.ResponseWriter, r *http.Request) {
	// Handle GET /commentcontroller request
}

func (c CommentController) Show(w http.ResponseWriter, r *http.Request) {
	// Handle GET /commentcontroller/{id} request
}

func (c CommentController) Create(w http.ResponseWriter, r *http.Request) {
	// Handle POST /commentcontroller request
	// You can get the request body by reading from r.Body
	// You can send a response by writing to w
}

func (c CommentController) Update(w http.ResponseWriter, r *http.Request) {
	// Handle PUT /commentcontroller/{id} request
	// You can get the request body by reading from r.Body
	// You can send a response by writing to w
}

func (c CommentController) Delete(w http.ResponseWriter, r *http.Request) {
	// Handle DELETE /commentcontroller/{id} request
}
