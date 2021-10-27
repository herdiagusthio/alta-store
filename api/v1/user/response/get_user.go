package response

import (
	"altaStore/business/user"
	"time"
)

//GetUserResponse Get user by ID response payload
type GetUserResponse struct {
	ID          uint       `json:"id"`
	Name        string     `json:"name"`
	Email       string     `json: email`
	PhoneNumber string     `json: phone_number`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at`
	DeletedAt   *time.Time `json:"deleted_at`
}

//NewGetUserResponse construct GetUserResponse
func NewGetUserResponse(user user.User) *GetUserResponse {
	var getUserResponse GetUserResponse

	getUserResponse.ID = user.ID
	getUserResponse.Name = user.Name
	getUserResponse.Email = user.Email
	getUserResponse.PhoneNumber = user.PhoneNumber
	getUserResponse.CreatedAt = user.CreatedAt
	getUserResponse.UpdatedAt = user.UpdatedAt
	getUserResponse.DeletedAt = user.DeletedAt

	return &getUserResponse
}
