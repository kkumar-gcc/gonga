package controllers

import (
	"net/http"

	"gorm.io/gorm"
)

type LikeController struct {
	DB *gorm.DB
}

func (c LikeController) Index(w http.ResponseWriter, r *http.Request) {
	// Handle GET /likecontroller request
}

func (c LikeController) Show(w http.ResponseWriter, r *http.Request) {
	// Handle GET /likecontroller/{id} request
}

func (c LikeController) Create(w http.ResponseWriter, r *http.Request) {
	// Handle POST /likecontroller request
	// You can get the request body by reading from r.Body
	// You can send a response by writing to w
}

func (c LikeController) Update(w http.ResponseWriter, r *http.Request) {
	// Handle PUT /likecontroller/{id} request
	// You can get the request body by reading from r.Body
	// You can send a response by writing to w
}

func (c LikeController) Delete(w http.ResponseWriter, r *http.Request) {
	// Handle DELETE /likecontroller/{id} request
}
