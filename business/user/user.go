package user

import "time"

type User struct {
	ID          int
	Name        string
	Email       string
	PhoneNumber string
	Username    string
	Password    string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}

//NewUser create new User
func NewUser(
	id int,
	name string,
	email string,
	phoneNumber string,
	username string,
	password string,
	createdAt time.Time) User {

	return User{
		ID:          id,
		Name:        name,
		Email:       email,
		PhoneNumber: phoneNumber,
		Username:    username,
		Password:    password,
		CreatedAt:   createdAt,
		UpdatedAt:   time.Now(),
		DeletedAt:   nil,
	}
}

//ModifyUser update existing User data
func (oldData *User) ModifyUser(newName string, newPhoneNumber string, newPassword string, modifiedAt time.Time) User {
	name := oldData.Name
	phoneNumber := oldData.PhoneNumber
	password := oldData.Password

	if newName != "" {
		name = newName
	}

	if newPhoneNumber != "" {
		phoneNumber = newPhoneNumber
	}

	if newPassword != "" {
		password = newPassword
	}

	return User{
		ID:          oldData.ID,
		Name:        name,
		Email:       oldData.Email,
		PhoneNumber: phoneNumber,
		Username:    oldData.Username,
		Password:    password,
		CreatedAt:   oldData.CreatedAt,
		UpdatedAt:   modifiedAt,
		DeletedAt:   oldData.DeletedAt,
	}
}
