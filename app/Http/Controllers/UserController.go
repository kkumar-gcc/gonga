package controllers

import (
	"net/http"

)

type UserController struct{}

func (uc UserController) Index(w http.ResponseWriter, r *http.Request) {
	// Handle GET /users request
	
}

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
