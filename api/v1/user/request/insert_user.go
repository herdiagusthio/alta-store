package request

import "altaStore/business/user"

type InsertUserRequest struct {
	Name        string `validate:"required"`
	Email       string `validate:"required,email"`
	PhoneNumber string `validate:"required,number"`
	Password    string `validate:"required"`
}

func (req *InsertUserRequest) ToUpsertUserSpec() *user.InsertUserSpec {

	var insertUserSpec user.InsertUserSpec

	insertUserSpec.Name = req.Name
	insertUserSpec.Email = req.Email
	insertUserSpec.PhoneNumber = req.PhoneNumber
	insertUserSpec.Password = req.Password

	return &insertUserSpec
}
