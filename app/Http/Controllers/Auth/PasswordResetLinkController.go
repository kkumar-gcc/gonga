package auth

import (
	"gonga/app/Models"
	"gonga/utils"
	"net/http"
	"time"

	"gorm.io/gorm"
)

type PasswordResetLinkController struct {
	DB *gorm.DB
}

type ResetPassowrdRequest struct {
	Email    string `json:"email" validate:"required,email"`
}


func (c PasswordResetLinkController) Index(w http.ResponseWriter, r *http.Request) {
	// Handle GET /passwordresetlinkcontroller request
}

func (c PasswordResetLinkController) Show(w http.ResponseWriter, r *http.Request) {
	// Handle GET /passwordresetlinkcontroller/{id} request
}

func (c PasswordResetLinkController) Create(w http.ResponseWriter, r *http.Request) {
    var resetPassword ResetPassowrdRequest
    if err := utils.DecodeRequestBody(r, &resetPassword); err != nil {
		utils.JSONResponse(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	// Validate user data
	if err := utils.ValidateRequest(w, &resetPassword); err != nil {
		return
	}
	// Check if a user with the given email exists
    if !utils.UserExistsWithEmail(c.DB, resetPassword.Email) {
        // Return a 404 response if the user doesn't exist
		utils.JSONResponse(w, http.StatusNotFound, map[string]string{"error": "User not found"})
        return
    }
    var passwordReset Models.PasswordReset
    passwordReset.Email = resetPassword.Email

    // Generate a unique token
    token, err := utils.GenerateRandomString(32)
    if err != nil {
		utils.JSONResponse(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
        return
    }
    passwordReset.Token = token

    // Set expiry time
    passwordReset.Expiry = time.Now().Add(time.Hour * 24).Unix()

    // Save the password reset token in the database
    if err := c.DB.Create(&passwordReset).Error; err != nil {
		utils.JSONResponse(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
        return
    }
    
    // Send email with the password reset link containing the token
	// if err := utils.SendPasswordResetEmail(passwordReset.Email, token); err != nil {
	// 	utils.JSONResponse(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
	// 	return
	// }
	utils.JSONResponse(w, http.StatusOK, map[string]string{"message": "Password reset link sent"})
}


func (c PasswordResetLinkController) Update(w http.ResponseWriter, r *http.Request) {
	// Handle PUT /passwordresetlinkcontroller/{id} request
	// You can get the request body by reading from r.Body
	// You can send a response by writing to w
}

func (c PasswordResetLinkController) Delete(w http.ResponseWriter, r *http.Request) {
	// Handle DELETE /passwordresetlinkcontroller/{id} request
}
