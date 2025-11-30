package helpers

import (
	"ibnufth/backend-api/models"
	"ibnufth/backend-api/structs"
)

// ToUserResponse converts a single user model to user response
func ToUserResponse(user models.User) structs.UserResponse {
	return structs.UserResponse{
		Id:        user.ID,
		Name:      user.Name,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: user.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}

// ToUserResponses converts a slice of user models to user responses
func ToUserResponses(users []models.User) []structs.UserResponse {
	responses := make([]structs.UserResponse, len(users))
	for i, user := range users {
		responses[i] = ToUserResponse(user)
	}
	return responses
}
