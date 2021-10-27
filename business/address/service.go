package address

import (
	"altaStore/business"
	"altaStore/util/validator"
	"time"
)

//InsertAddressSpec create address spec
type InsertAddressSpec struct {
	UserID      uint   `validate:"required"`
	Name        string `validate:"required"`
	PhoneNumber string `validate:"required,number"`
	Street      string `validate:"required"`
	City        string `validate:"required"`
	Province    string `validate:"required"`
	District    string `validate:"required`
	PostalCode  uint   `validate:"required`
	AddressType *string
	IsDefault   bool
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

//InsertAddress Create new address and store into database
func (s *service) InsertAddress(insertAddressSpec InsertAddressSpec) error {
	err := validator.GetValidator().Struct(insertAddressSpec)
	if err != nil {
		return business.ErrInvalidSpec
	}

	defaultAddress, err := s.repository.GetDefaultAddress(insertAddressSpec.UserID)
	if err != nil {
		if insertAddressSpec.IsDefault != true {
			insertAddressSpec.IsDefault = true
		}
	} else {
		if insertAddressSpec.IsDefault == true {
			err := s.repository.UpdateDefaultAddress(*defaultAddress)
			if err != nil {
				return business.ErrInternalServerError
			}
		}
	}

	if *insertAddressSpec.AddressType == "" {
		insertAddressSpec.AddressType = nil
	}
	address := NewAddress(
		insertAddressSpec.UserID,
		insertAddressSpec.Name,
		insertAddressSpec.PhoneNumber,
		insertAddressSpec.Street,
		insertAddressSpec.City,
		insertAddressSpec.Province,
		insertAddressSpec.District,
		insertAddressSpec.PostalCode,
		insertAddressSpec.AddressType,
		insertAddressSpec.IsDefault,
		time.Now(),
	)

	err = s.repository.InsertAddress(address)
	if err != nil {
		return err
	}

	return nil
}

//GetDefaultAddress Get default address from user
func (s *service) GetDefaultAddress(UserID uint) (*Address, error) {
	return s.repository.GetDefaultAddress(UserID)
}

//GetAllAddress Get all addresses , will be return empty array if no data or error occured
func (s *service) GetAllAddress(UserID uint) ([]Address, error) {

	address, err := s.repository.GetAllAddress(UserID)
	if err != nil {
		return []Address{}, err
	}

	return address, err
}
