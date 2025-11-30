package controller

import (
	"ibnufth/backend-api/database"
	"ibnufth/backend-api/helpers"
	"ibnufth/backend-api/models"
	"ibnufth/backend-api/structs"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FindUser(c *gin.Context) {
	// initiate slice for users
	var users []models.User

	// find users
	database.DB.Find(&users)

	// convert to user responses
	userResponses := helpers.ToUserResponses(users)

	// return users
	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "List Data Users",
		Data:    userResponses,
	})
}

func CreateUser(c *gin.Context) {
	// struct user request
	var req = structs.UserCreateRequest{}

	// bind JSON request to struct UserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, structs.ErrorResponse{
			Success: false,
			Message: "Validation Error",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	user := models.User{
		Name:     req.Name,
		Email:    req.Email,
		Username: req.Username,
		Password: helpers.HashPassword(req.Password),
	}

	if err := database.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
			Success: false,
			Message: "Failed to create user",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	userResponse := helpers.ToUserResponse(user)

	c.JSON(http.StatusCreated, structs.SuccessResponse{
		Success: true,
		Message: "User created successfully",
		Data:    userResponse,
	})
}

func FindUserById(c *gin.Context) {
	// get user id from URL parameter
	id := c.Param("id")

	// initiate user model
	var user models.User

	// find user by id
	if err := database.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, structs.ErrorResponse{
			Success: false,
			Message: "User not found",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	// convert to user response
	userResponse := helpers.ToUserResponse(user)

	// return user
	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "User found",
		Data:    userResponse,
	})
}

func UpdateUser(c *gin.Context) {
	id := c.Param("id")

	var user models.User

	if err := database.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, structs.ErrorResponse{
			Success: false,
			Message: "User not found",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	var req = structs.UserUpdateRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, structs.ErrorResponse{
			Success: false,
			Message: "Validation Error",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	user.Name = req.Name
	user.Username = req.Username
	user.Email = req.Email
	if req.Password != "" {
		user.Password = helpers.HashPassword(req.Password)
	}

	if err := database.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
			Success: false,
			Message: "Failed to update user",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	userResponse := helpers.ToUserResponse(user)

	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "User updated successfully",
		Data:    userResponse,
	})
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")

	var user models.User

	if err := database.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, structs.ErrorResponse{
			Success: false,
			Message: "User not found",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	if err := database.DB.Delete(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
			Success: false,
			Message: "Failed to delete user",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "User deleted successfully",
	})
}