package controllers

import (
	"net/http"

	"gorm.io/gorm"
)

type FollowController struct {
	DB *gorm.DB
}

func (c FollowController) Index(w http.ResponseWriter, r *http.Request) {
	// Handle GET /followcontroller request
}

func (c FollowController) Show(w http.ResponseWriter, r *http.Request) {
	// Handle GET /followcontroller/{id} request
}

func (c FollowController) Create(w http.ResponseWriter, r *http.Request) {
	// Handle POST /followcontroller request
	// You can get the request body by reading from r.Body
	// You can send a response by writing to w
}

func (c FollowController) Update(w http.ResponseWriter, r *http.Request) {
	// Handle PUT /followcontroller/{id} request
	// You can get the request body by reading from r.Body
	// You can send a response by writing to w
}

func (c FollowController) Delete(w http.ResponseWriter, r *http.Request) {
	// Handle DELETE /followcontroller/{id} request
}
