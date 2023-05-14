package controllers

import (
	requests "gonga/app/Http/Requests"
	"gonga/app/Models"
	cloudinary "gonga/packages/Cloudinary"
	"gonga/utils"
	"net/http"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PostController struct {
	DB *gorm.DB
}

func (c PostController) Index(w http.ResponseWriter, r *http.Request) {

	var posts []Models.Post
	result, err := utils.Paginate(r, c.DB, &posts, "User")

	if err != nil {
		utils.JSONResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	// Send the list of users in JSON format
	utils.JSONResponse(w, http.StatusOK, result)
	// Handle GET /postcontroller request
}

func (c PostController) Show(w http.ResponseWriter, r *http.Request) {
	// Handle GET /postcontroller/{id} request
}

func (c PostController) Create(w http.ResponseWriter, r *http.Request) {
	// Parse update request from request body
	var createReq requests.CreatePostRequest
	if err := utils.DecodeRequestBody(r, &createReq); err != nil {
		utils.JSONResponse(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	// Validate the update request
	if err := utils.ValidateRequest(w, &createReq); err != nil {
		return
	}

	newPost := Models.Post{
		Title:           createReq.Title,
		Body:            createReq.Body,
		IsPromoted:      createReq.IsPromoted,
		IsFeatured:      createReq.IsFeatured,
		IsPublished:     createReq.IsPublished,
		PromotionExpiry: createReq.PromotionExpiry,
		FeaturedExpiry:  createReq.FeaturedExpiry,
	}
	result := c.DB.Create(&newPost)
	if result.Error != nil {
		utils.JSONResponse(w, http.StatusInternalServerError, map[string]string{"error": result.Error.Error()})
		return
	}

	// Send email or notification to subscribed users
	// if err := sendPasswordResetEmail(passwordReset.Email, token); err != nil {
	// 	utils.JSONResponse(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
	// 	return
	// }
	utils.JSONResponse(w, http.StatusOK, map[string]string{"message": "Post created successfully!"})
}

func (c PostController) Update(w http.ResponseWriter, r *http.Request) {
	// Handle PUT /postcontroller/{id} request
	// You can get the request body by reading from r.Body
	// You can send a response by writing to w
	file, _, err := r.FormFile("file")
	if err != nil {
		utils.JSONResponse(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	defer file.Close()
	// Generate a unique public ID for the file
	publicID := uuid.New().String()

	cloudinaryClient := cloudinary.NewCloudinaryClient()
	result, err := cloudinaryClient.UploadImage(file, publicID)
	if err != nil {
		utils.JSONResponse(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	utils.JSONResponse(w, http.StatusOK, result)
}

func (c PostController) Delete(w http.ResponseWriter, r *http.Request) {
	// Handle DELETE /postcontroller/{id} request
}
