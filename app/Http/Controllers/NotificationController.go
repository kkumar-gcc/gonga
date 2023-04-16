package controllers

import (
	"net/http"

	"gorm.io/gorm"
)

type NotificationController struct {
	DB *gorm.DB
}

func (c NotificationController) Index(w http.ResponseWriter, r *http.Request) {
	// Handle GET /notificationcontroller request
}

func (c NotificationController) Show(w http.ResponseWriter, r *http.Request) {
	// Handle GET /notificationcontroller/{id} request
}

func (c NotificationController) Create(w http.ResponseWriter, r *http.Request) {
	// Handle POST /notificationcontroller request
	// You can get the request body by reading from r.Body
	// You can send a response by writing to w
}

func (c NotificationController) Update(w http.ResponseWriter, r *http.Request) {
	// Handle PUT /notificationcontroller/{id} request
	// You can get the request body by reading from r.Body
	// You can send a response by writing to w
}
func (c NotificationController) ReadAll(w http.ResponseWriter, r *http.Request) {
	// Handle PUT /notifications/{id} request
	// You can get the request body by reading from r.Body
	// You can send a response by writing to w
}

func (c NotificationController) Delete(w http.ResponseWriter, r *http.Request) {
	// Handle DELETE /notificationcontroller/{id} request
}
