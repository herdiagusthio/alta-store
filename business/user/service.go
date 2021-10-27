package user

import (
	"altaStore/business"
	util "altaStore/util/password"
	"altaStore/util/validator"
	"time"
)

//InsertUserSpec create user spec
type InsertUserSpec struct {
	Name        string `validate:"required"`
	Email       string `validate:"required,email"`
	PhoneNumber string `validate:"required,number"`
	Password    string `validate:"required"`
}

//=============== The implementation of those interface put below =======================
type service struct {
	repository Repository
}

//NewService Construct user service object
func NewService(repository Repository) Service {
	return &service{
		repository,
	}
}

//FindUserByID Get user by given ID, return nil if not exist
func (s *service) FindUserByID(id int) (*User, error) {
	return s.repository.FindUserByID(id)
}

//FindUserByEmailAndPassword Get user by given ID, return nil if not exist
func (s *service) FindUserByEmail(email string) (*User, error) {
	return s.repository.FindUserByEmail(email)
}

//FindAllUser Get all users , will be return empty array if no data or error occured
func (s *service) FindAllUser(skip int, rowPerPage int) ([]User, error) {

	user, err := s.repository.FindAllUser(skip, rowPerPage)
	if err != nil {
		return []User{}, err
	}

	return user, err
}

//InsertUser Create new user and store into database
func (s *service) InsertUser(insertUserSpec InsertUserSpec) error {
	err := validator.GetValidator().Struct(insertUserSpec)
	if err != nil {
		return business.ErrInvalidSpec
	}
	hashPassword, errorEncrypt := util.EncryptPassword(insertUserSpec.Password)
	if errorEncrypt != nil {
		return errorEncrypt
	}
	user := NewUser(
		insertUserSpec.Name,
		insertUserSpec.Email,
		insertUserSpec.PhoneNumber,
		string(hashPassword),
		time.Now(),
	)

	err = s.repository.InsertUser(user)
	if err != nil {
		return err
	}

	return nil
}

//UpdateUser will update found user, if not exists will be return error
// func (s *service) UpdateUser(id int, name string, modifiedBy string, currentVersion int) error {

// 	user, err := s.repository.FindUserByID(id)
// 	if err != nil {
// 		return err
// 	} else if user == nil {
// 		return business.ErrNotFound
// 	} else if user.Version != currentVersion {
// 		return business.ErrHasBeenModified
// 	}

// 	modifiedUser := user.ModifyUser(name, time.Now(), modifiedBy)

// 	return s.repository.UpdateUser(modifiedUser, currentVersion)
// }
