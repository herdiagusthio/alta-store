package address

import (
	"time"
)

type Address struct {
	ID          uint
	UserID      uint
	Name        string
	PhoneNumber string
	Street      string
	City        string
	Province    string
	District    string
	PostalCode  uint
	AddressType *string
	IsDefault   bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}

//NewAddress
func NewAddress(
	userID uint,
	name string,
	phoneNumber string,
	street string,
	city string,
	province string,
	district string,
	postalCode uint,
	addressType *string,
	isDefault bool,
	createdAt time.Time) Address {

	return Address{
		UserID:      userID,
		Name:        name,
		PhoneNumber: phoneNumber,
		Street:      street,
		City:        city,
		Province:    province,
		District:    district,
		PostalCode:  postalCode,
		AddressType: addressType,
		IsDefault:   isDefault,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		DeletedAt:   nil,
	}
}
