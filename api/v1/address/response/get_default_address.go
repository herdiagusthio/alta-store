package response

import (
	"altaStore/business/address"
	"time"
)

//GetAddressResponse Get default address by UserID response payload
type GetAddressResponse struct {
	ID          uint    `json:"id"`
	UserID      uint    `json:"user_id"`
	Name        string  `json:"name"`
	PhoneNumber string  `json:"phone_number"`
	Street      string  `json:"street"`
	City        string  `json:"city"`
	Province    string  `json:"province"`
	District    string  `json:"district"`
	PostalCode  uint    `json:"postal_code"`
	AddressType *string `json:"address_type"`
	IsDefault   bool    `json:"is_default"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}

//NewGetUserResponse construct GetUserResponse
func NewGetAddressResponse(address address.Address) *GetAddressResponse {
	var getAddressResponse GetAddressResponse

	getAddressResponse.ID = address.ID
	getAddressResponse.UserID = address.UserID
	getAddressResponse.Name = address.Name
	getAddressResponse.PhoneNumber = address.PhoneNumber
	getAddressResponse.Street = address.Street
	getAddressResponse.City = address.City
	getAddressResponse.Province = address.Province
	getAddressResponse.District = address.District
	getAddressResponse.PostalCode = address.PostalCode
	getAddressResponse.AddressType = address.AddressType
	getAddressResponse.IsDefault = address.IsDefault
	getAddressResponse.CreatedAt = address.CreatedAt
	getAddressResponse.UpdatedAt = address.UpdatedAt
	getAddressResponse.DeletedAt = address.DeletedAt

	return &getAddressResponse
}
