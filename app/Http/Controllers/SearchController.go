package controllers

import (
	"net/http"

	"gorm.io/gorm"
)

type SearchController struct {
	DB *gorm.DB
}

func (c SearchController) Index(w http.ResponseWriter, r *http.Request) {
	// Handle GET /searchcontroller request
}

func (c SearchController) Show(w http.ResponseWriter, r *http.Request) {
	// Handle GET /searchcontroller/{id} request
}

func (c SearchController) Create(w http.ResponseWriter, r *http.Request) {
	// Handle POST /searchcontroller request
	// You can get the request body by reading from r.Body
	// You can send a response by writing to w
}

func (c SearchController) Update(w http.ResponseWriter, r *http.Request) {
	// Handle PUT /searchcontroller/{id} request
	// You can get the request body by reading from r.Body
	// You can send a response by writing to w
}

func (c SearchController) Delete(w http.ResponseWriter, r *http.Request) {
	// Handle DELETE /searchcontroller/{id} request
}
