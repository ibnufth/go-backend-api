package controller

import (
	"ibnufth/backend-api/database"
	"ibnufth/backend-api/helpers"
	"ibnufth/backend-api/models"
	"ibnufth/backend-api/structs"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	// initialize a variable to hold the incoming JSON data
	var req = structs.UserCreateRequest{}

	// validate the incoming JSON data against the struct by binding from Gin
	if err := c.ShouldBindJSON(&req); err != nil {
		// if there is an error, return a JSON response with the error details
		c.JSON(http.StatusUnprocessableEntity, structs.ErrorResponse{
			Success: false,
			Message: "Validation Errors",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	// create new user with hashed password
	user := models.User{
		Name:     req.Name,
		Email:    req.Email,
		Username: req.Username,
		Password: helpers.HashPassword(req.Password),
	}

	if err := database.DB.Create(&user).Error; err != nil {
		if helpers.IsDuplicateEntryError(err) {
			c.JSON(http.StatusConflict, structs.ErrorResponse{
				Success: false,
				Message: "Duplicate entry error",
			})

		} else {
			c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
				Success: false,
				Message: "Failed to create user",
				Errors:  helpers.TranslateErrorMessage(err),
			})
		}
		return
	}

	c.JSON(http.StatusCreated, structs.SuccessResponse{
		Success: true,
		Message: "User created successfully",
		Data: structs.UserResponse{
			Id:        user.ID,
			Name:      user.Name,
			Email:     user.Email,
			Username:  user.Username,
			CreatedAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: user.UpdatedAt.Format("2006-01-02 15:04:05"),
		},
	})
}
