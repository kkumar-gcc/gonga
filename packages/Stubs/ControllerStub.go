package Stubs

import (
	"fmt"
	"strings"
)

func GetControllerStub(controllerName string) string {
	return fmt.Sprintf(`package controllers

import (
	"net/http"

	"gorm.io/gorm"
)

type %s struct {
	DB *gorm.DB
}

func (c %s) Index(w http.ResponseWriter, r *http.Request) {
	// Handle GET /%s request
}

func (c %s) Show(w http.ResponseWriter, r *http.Request) {
	// Handle GET /%s/{id} request
}

func (c %s) Create(w http.ResponseWriter, r *http.Request) {
	// Handle POST /%s request
	// You can get the request body by reading from r.Body
	// You can send a response by writing to w
}

func (c %s) Update(w http.ResponseWriter, r *http.Request) {
	// Handle PUT /%s/{id} request
	// You can get the request body by reading from r.Body
	// You can send a response by writing to w
}

func (c %s) Delete(w http.ResponseWriter, r *http.Request) {
	// Handle DELETE /%s/{id} request
}
`, controllerName, controllerName, strings.ToLower(controllerName), controllerName, strings.ToLower(controllerName), controllerName, strings.ToLower(controllerName), controllerName, strings.ToLower(controllerName), controllerName, strings.ToLower(controllerName))

}
